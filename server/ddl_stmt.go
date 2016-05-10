package server

import (
	"encoding/json"
	"fmt"

	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql"

	"github.com/golang/glog"
	"github.com/pingcap/tidb/ast"
	ftype "github.com/pingcap/tidb/mysql"
	"github.com/pingcap/tidb/parser"
)

func (sei *session) handleDDL(sql string) error {
	stmt, err := parser.ParseOneStmt(sql, "", "")
	if err != nil {
		glog.Infof("parse ddl sql(%s) error:%v", sql, err)
		return sei.writeError(mysql.NewDefaultError(mysql.ER_SYNTAX_ERROR))
	}
	switch v := stmt.(type) {
	case *ast.CreateDatabaseStmt:
		dbname := v.Name
		id, rows, err := sei.ddlManage().CreateDatabase(sei.user, dbname, sei.dbnum)
		glog.Infof("DDL plan id(%v)", id)
		if err != nil {
			glog.Infof("CREATE DATABASE has an error(%v)", err)
			err = sei.writeError(err)
		} else {
			// one time only creata a db
			r := &mysql.Result{
				AffectedRows: rows,
			}
			err = sei.writeOK(r)
		}
		return err

	case *ast.CreateTableStmt:
		if sei.db == "" {
			return sei.writeError(mysql.NewDefaultError(mysql.ER_NO_DB_ERROR))
		}
		table := &meta.Table{
			Scheme: "hash",
			Name:   v.Table.Name.String(),
			PartitionKey: &meta.Key{
				Name: sei.partitionKey,
			},
			ColsLen: len(v.Cols),
		}

		existMap := make(map[string]bool)
		// get constraints
		for _, constraint := range v.Constraints {
			if constraint.Tp == ast.ConstraintPrimaryKey ||
				constraint.Tp == ast.ConstraintUniq {
				if len(constraint.Keys) > 1 {
					err := mysql.NewError(mysql.ER_SYNTAX_ERROR,
						"not support constraint keys' length > 1")
					return sei.writeError(err)
				}
				// get type
				name := constraint.Keys[0].Column.Name.String()
				index, typ := sei.getFieldType(v.Cols, name)
				if typ == meta.TypeKeyUnknow {
					err := mysql.NewError(mysql.ER_SYNTAX_ERROR,
						"unsupport key's type ")
					return sei.writeError(err)
				}

				if constraint.Tp == ast.ConstraintPrimaryKey && sei.partitionKey == "" {
					// set primary key for partition key
					table.PartitionKey.Name = name
					table.PartitionKey.Type = typ
					table.PartitionKey.Index = index
				}

				table.AutoKeys = append(table.AutoKeys, &meta.Key{
					Name:  name,
					Type:  typ,
					Index: index,
				})
				existMap[name] = true
			}
		}
		// check auto increment
		for _, col := range v.Cols {
			for n, option := range col.Options {
				t := sei.getOneFeildType(col)
				if t == meta.TypeKeyUnknow {
					err := mysql.NewError(mysql.ER_SYNTAX_ERROR,
						"unsupport key's type ")
					return sei.writeError(err)
				}
				switch option.Tp {
				case ast.ColumnOptionAutoIncrement, ast.ColumnOptionPrimaryKey, ast.ColumnOptionUniq:
					if ast.ColumnOptionPrimaryKey == option.Tp && table.PartitionKey.Name == "" {
						table.PartitionKey.Name = col.Name.Name.String()
						table.PartitionKey.Type = t
						table.PartitionKey.Index = n
					}
					// check if exist not append
					if existMap[col.Name.Name.String()] {
						continue
					}
					table.AutoKeys = append(table.AutoKeys, &meta.Key{
						Name:  col.Name.Name.String(),
						Type:  t,
						Index: n,
					})
					existMap[col.Name.Name.String()] = true
				}
				if option.Tp == ast.ColumnOptionAutoIncrement {
					glog.Infof("record auto increment option index(%v)", n)
					// record
					table.AutoIns = append(table.AutoIns, &meta.Key{
						Name:  col.Name.Name.String(),
						Type:  t,
						Index: n,
					})
				}
			}
		}
		// check partition key
		if table.PartitionKey.Name == "" {
			err := mysql.NewError(mysql.ER_SYNTAX_ERROR,
				"partitionKey is null")
			return sei.writeError(err)
		}
		data, _ := json.MarshalIndent(table, "", "\t")
		glog.Info("table is\n", string(data))
		id, rows, err := sei.ddlManage().CreateTable(sei.user, sei.db, sql, table)
		glog.Infof("DDL plan id(%v)", id)
		if err != nil {
			glog.Infof("CREATE TABLE has an error(%v)", err)
			err = sei.writeError(err)
		} else {
			// one time only creata a db
			r := &mysql.Result{
				AffectedRows: rows,
			}
			err = sei.writeOK(r)
		}
		return err

	default:
		return fmt.Errorf("create statement %T not support now", stmt)
	}
}

func (sei *session) getFieldType(cols []*ast.ColumnDef, name string) (index, typ int) {
	for n, col := range cols {
		if col.Name.Name.String() == name {
			return n, sei.getOneFeildType(col)
		}
	}
	return 0, meta.TypeKeyUnknow
}

func (sei *session) getOneFeildType(col *ast.ColumnDef) int {
	switch col.Tp.Tp {
	case ftype.TypeString, ftype.TypeVarchar:
		return meta.TypeKeyString
	case ftype.TypeTimestamp, ftype.TypeShort, ftype.TypeLong,
		ftype.TypeFloat, ftype.TypeDouble, ftype.TypeDecimal:
		return meta.TypeKeyInt
	default:
		return meta.TypeKeyUnknow
	}
}
