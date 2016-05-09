package server

import (
	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/golang/glog"
)

func (sei *session) handleInsert(stmt *sqlparser.Insert) error {
	if sei.db == "" {
		return sei.writeError(mysql.NewDefaultError(mysql.ER_NO_DB_ERROR))
	}

	isNoCol := false
	isNeedInitCol := true
	colLen := len(stmt.Columns)

	// get table
	table, err := sei.getMeta().GetTable(sei.user, sei.db, string(stmt.Table.Name))
	if err != nil {
		return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
	}
	if colLen == 0 {
		isNoCol = true
	}
	// get sqls
	values, ok := stmt.Rows.(sqlparser.Values)
	if !ok {
		return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport"))
	}

	for row, value := range values {
		tuple, ok := value.(sqlparser.ValTuple)
		if !ok {
			return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport"))
		}
		if (isNoCol && len(tuple) != table.Table.ColsLen) ||
			!isNoCol && len(tuple) != colLen {
			return sei.writeError(mysql.NewDefaultError(mysql.ER_WRONG_VALUE_COUNT_ON_ROW, row+1))
		}
		// rewrite auto key
		if isNoCol {
			// rewrite by index
			for _, autokey := range table.Table.AutoKeys {
				id, err := table.GetKey(autokey.Name, config.Config.Proxy.AutoKeyInterval, sei.server.info)
				if err != nil {
					glog.Warningf("Get autokey(%v) get error(%v)", autokey.Name, err)
					return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
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
					return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
				}
				for n, column := range stmt.Columns {
					col, ok := column.(*sqlparser.NonStarExpr)
					if !ok {
						return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport"))
					}
					colName, ok := col.Expr.(*sqlparser.ColName)
					if !ok {
						return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unsupport"))
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
				return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type"))
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
						return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type"))
					}
				}
			}
		}
		if value == "" {
			glog.Warning("Get partition key error value is null")
			return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "unknow partition key type"))
		}
		index := FindForKey(value, len(table.Backends))
		s, ok := sqls[index]
		if !ok {
			sqls[index] = sqlparser.Values{values[row]}
		} else {
			sqls[index] = append(s.(sqlparser.Values), values[row])
		}
	}
	// print sqls
	buf := sqlparser.NewTrackedBuffer(nil)
	for n, v := range sqls {
		stmt.Rows = v
		stmt.Format(buf)
		glog.Infof("partition:%d sql is:%s,backend:%v", n, buf.String())
		buf.Reset()
	}
	sei.writeOK(nil)
	return nil
}
