package server

import (
	"errors"
	"fmt"
	"sort"
	"strconv"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/golang/glog"
)

type Plan struct {
	SQLs  []*SQL
	Table *Table

	offset int64
	count  int64
}

type SQL struct {
	Backend *meta.Backend
	SQL     []string
}

func (p *Plan) generateSQL(statement ...sqlparser.Statement) error {
	switch stmt := statement[0].(type) {
	case *sqlparser.Select:
		p.generateSQLALL(stmt)

	case *sqlparser.Update:
		p.generateSQLALL(stmt)

	case *sqlparser.Delete:
		p.generateSQLALL(stmt)

	default:
		return errors.New("not support plan")
	}
	return nil
}

func (p *Plan) generateSQLALL(stmt sqlparser.Statement) {
	buf := sqlparser.NewTrackedBuffer(nil)
	stmt.Format(buf)
	sql := buf.String()
	glog.Infof("generate select SQL:%s", sql)
	for _, backend := range p.Table.Backends {
		p.SQLs = append(p.SQLs, &SQL{
			Backend: backend,
			SQL:     []string{sql},
		})
	}
}

func (p *Plan) generateOneSQL(stmt sqlparser.Statement, index int) {
	buf := sqlparser.NewTrackedBuffer(nil)
	stmt.Format(buf)
	sql := buf.String()
	glog.Infof("generate select SQL:%s", sql)
	p.SQLs = append(p.SQLs, &SQL{
		Backend: p.Table.Backends[index],
		SQL:     []string{sql},
	})
}

// limit
func (p *Plan) rewriteLimit(stmt *sqlparser.Select) (err error) {
	origin := stmt.Limit
	// get offset and count
	if origin == nil {
		return
	}

	var offset, count int64

	if origin.Offset == nil {
		offset = 0
	} else {
		o, ok := origin.Offset.(sqlparser.NumVal)
		if !ok {
			err = fmt.Errorf("invalid select limit %s", sqlparser.String(stmt.Limit))
			return
		}
		if offset, err = strconv.ParseInt(string([]byte(o)), 10, 64); err != nil {
			return
		}
	}

	o, ok := origin.Rowcount.(sqlparser.NumVal)
	if !ok {
		err = fmt.Errorf("invalid limit %s", sqlparser.String(stmt.Limit))
		return
	}
	if count, err = strconv.ParseInt(string([]byte(o)), 10, 64); err != nil {
		return
	}
	if count < 0 {
		err = fmt.Errorf("invalid limit %s", sqlparser.String(stmt.Limit))
		return
	}

	// rewrite limit stmt
	stmt.Limit.Offset = sqlparser.NumVal([]byte("0"))
	stmt.Limit.Rowcount = sqlparser.NumVal([]byte(strconv.FormatInt(count+offset, 10)))

	p.offset = offset
	p.count = count
	return
}

//build a router plan
func (sei *session) buildPlan(statement sqlparser.Statement) (*Plan, error) {
	switch stmt := statement.(type) {
	case *sqlparser.Insert:
		return sei.buildInsertPlan(stmt)
	case *sqlparser.Select:
		return sei.buildSelectPlan(stmt)
	case *sqlparser.Update:
		return sei.buildUpdatePlan(stmt)
	case *sqlparser.Delete:
		return sei.buildDelPlan(stmt)
	}
	return nil, errors.New("not support plan")
}

func (sei *session) buildInsertPlan(stmt *sqlparser.Insert) (*Plan, error) {
	plan := &Plan{}

	isNoCol := false
	isNeedInitCol := true
	colLen := len(stmt.Columns)

	// get table
	table, err := sei.getMeta().GetTable(sei.user, sei.db, string(stmt.Table.Name))
	if err != nil {
		return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error())
	}
	plan.Table = table
	if colLen == 0 {
		isNoCol = true
	}
	// get sqls
	values, ok := stmt.Rows.(sqlparser.Values)
	if !ok {
		return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport")
	}

	for row, value := range values {
		tuple, ok := value.(sqlparser.ValTuple)
		if !ok {
			return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport")
		}
		if (isNoCol && len(tuple) != table.Table.ColsLen) ||
			!isNoCol && len(tuple) != colLen {
			return nil, mysql.NewDefaultError(mysql.ER_WRONG_VALUE_COUNT_ON_ROW, row+1)
		}
		// rewrite auto key
		if isNoCol {
			// rewrite by index
			for _, autokey := range table.Table.AutoKeys {
				id, err := table.GetKey(autokey.Name, config.Config.Proxy.AutoKeyInterval, sei.server.info)
				if err != nil {
					glog.Warningf("Get autokey(%v) get error(%v)", autokey.Name, err)
					return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error())
				}
				if autokey.Type == meta.TypeKeyInt {
					tuple[autokey.Index] = sqlparser.NumVal(id)
				} else {
					tuple[autokey.Index] = sqlparser.StrVal(id)
				}
			}
		} else {
			for _, autokey := range table.Table.AutoKeys {
				isExist := false
				id, err := table.GetKey(autokey.Name, config.Config.Proxy.AutoKeyInterval, sei.server.info)
				if err != nil {
					glog.Warningf("Get autokey(%v) get error(%v)", autokey.Name, err)
					return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error())
				}
				for n, column := range stmt.Columns {
					col, ok := column.(*sqlparser.NonStarExpr)
					if !ok {
						return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport")
					}
					colName, ok := col.Expr.(*sqlparser.ColName)
					if !ok {
						return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport")
					}
					if autokey.Name == string(colName.Name) && n < len(tuple) {
						isExist = true
						if autokey.Type == meta.TypeKeyInt {
							tuple[n] = sqlparser.NumVal(id)
						} else {
							tuple[n] = sqlparser.StrVal(id)
						}
					}
				}
				if !isExist {
					if isNeedInitCol {
						stmt.Columns = append(stmt.Columns, &sqlparser.NonStarExpr{
							Expr: &sqlparser.ColName{
								Name: sqlparser.SQLName(autokey.Name),
							},
						})
					}
					if autokey.Type == meta.TypeKeyInt {
						tuple = append(tuple, sqlparser.NumVal(id))
					} else {
						tuple = append(tuple, sqlparser.StrVal(id))
					}
				}
			}
		}
		values[row] = tuple
		isNeedInitCol = false
	}
	stmt.Rows = values

	sqls := make(map[int]sqlparser.InsertRows)
	// split stmt
	for row, value := range values {
		tuple := value.(sqlparser.ValTuple)
		var value string
		if isNoCol {
			// get hash key
			switch v := tuple[table.Table.PartitionKey.Index].(type) {
			case sqlparser.NumVal:
				value = string(v)
			case sqlparser.StrVal:
				value = string(v)
			default:
				return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type")
			}
		} else {
			for col, column := range stmt.Columns {
				name := string(column.(*sqlparser.NonStarExpr).Expr.(*sqlparser.ColName).Name)
				if name == table.Table.PartitionKey.Name {
					switch v := tuple[col].(type) {
					case sqlparser.NumVal:
						value = string(v)
					case sqlparser.StrVal:
						value = string(v)
					default:
						return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type")
					}
				}
			}
		}
		if value == "" {
			glog.Warning("Get partition key error value is null")
			return nil, mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type")
		}
		index := FindForKey(value, len(table.Backends))
		s, ok := sqls[index]
		if !ok {
			sqls[index] = sqlparser.Values{values[row]}
		} else {
			sqls[index] = append(s.(sqlparser.Values), values[row])
		}
	}

	var lastIDs []uint64
	if len(table.Table.AutoIns) > 0 {
		index := 0
		if isNoCol {
			// range column
			index = table.Table.AutoIns[0].Index
		} else {
			// get index
			index = sort.Search(len(stmt.Columns), func(i int) bool {
				return string(stmt.Columns[i].(*sqlparser.NonStarExpr).Expr.(*sqlparser.ColName).Name) ==
					table.Table.AutoIns[0].Name
			})
		}
		lastIDs = make([]uint64, len(sqls))
		n := 0
		for _, sql := range sqls {
			switch v := sql.(sqlparser.Values)[0].(sqlparser.ValTuple)[index].(type) {
			case sqlparser.NumVal:
				glog.Info(n, len(lastIDs))
				lastIDs[n], _ = strconv.ParseUint(string(v), 10, 64)
			case sqlparser.StrVal:
				lastIDs[n], _ = strconv.ParseUint(string(v), 10, 64)
			}
			n++
		}
	}

	for k, v := range sqls {
		stmt.Rows = v
		plan.generateOneSQL(stmt, k)
	}
	return plan, nil
}

func (sei *session) buildSelectPlan(stmt *sqlparser.Select) (*Plan, error) {
	plan := &Plan{}
	var err error
	var tableName string

	switch v := (stmt.From[0]).(type) {
	case *sqlparser.AliasedTableExpr:
		tableName = sqlparser.String(v.Expr)
	case *sqlparser.JoinTableExpr:
		if ate, ok := (v.LeftExpr).(*sqlparser.AliasedTableExpr); ok {
			tableName = sqlparser.String(ate.Expr)
		} else {
			tableName = sqlparser.String(v)
		}
	default:
		tableName = sqlparser.String(v)
	}

	// get tableName
	t, err := sei.getMeta().GetTable(sei.user, sei.db, tableName)
	if err != nil {
		return nil, err
	}
	plan.Table = t

	// rewrite limit
	err = plan.rewriteLimit(stmt)
	if err != nil {
		return nil, err
	}

	// generate sql
	err = plan.generateSQL(stmt)
	return plan, err
}

func (sei *session) buildDelPlan(stmt *sqlparser.Delete) (*Plan, error) {
	plan := &Plan{}
	var err error
	var tableName = sqlparser.String(stmt.Table)
	// get tableName
	t, err := sei.getMeta().GetTable(sei.user, sei.db, tableName)
	if err != nil {
		return nil, err
	}
	plan.Table = t

	// generate sql
	err = plan.generateSQL(stmt)
	return plan, err
}

func (sei *session) buildUpdatePlan(stmt *sqlparser.Update) (*Plan, error) {
	plan := &Plan{}
	var err error
	var tableName = sqlparser.String(stmt.Table)
	// get tableName
	t, err := sei.getMeta().GetTable(sei.user, sei.db, tableName)
	if err != nil {
		return nil, err
	}
	plan.Table = t

	// generate sql
	err = plan.generateSQL(stmt)
	return plan, err
}

func (p *Plan) getOffetCount() (int64, int64) {
	return p.offset, p.count
}

func (sei *session) executePlan(plan *Plan) ([]*mysql.Result, error) {
	var rs []*mysql.Result
	for _, sql := range plan.SQLs {
		tr, err := sei.executeBackend(sql.Backend, sql.SQL)
		if err != nil {
			return nil, err
		}
		rs = append(rs, tr...)
	}
	return rs, nil
}

func (sei *session) executeBackend(backend *meta.Backend, sqls []string) ([]*mysql.Result, error) {
	var (
		r  *mysql.Result
		rs []*mysql.Result
	)
	conn, err := sei.getBpool().GetConn(backend)
	if err != nil {
		return nil, err
	}
	for _, sql := range sqls {
		r, err = conn.Execute(sql)
		if err != nil {
			break
		}
		rs = append(rs, r)
	}
	sei.getBpool().PushConn(backend, conn, err)
	return rs, err
}
