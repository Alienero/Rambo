package server

import (
	"errors"
	"strconv"

	"github.com/Alienero/Rambo/mysql"
	"github.com/Alienero/Rambo/mysql/sqlparser"
	"github.com/golang/glog"
)

var funcNameMap = map[string]string{
	"sum":   "sum",
	"count": "count",
	"max":   "max",
	"min":   "min",
}

func (sei *session) handleSelect(stmt *sqlparser.Select) error {
	if stmt.From == nil {
		return sei.handleSimpleSelect(stmt)
	}
	if len(stmt.From) > 1 {
		return sei.writeError(mysql.NewDefaultError(mysql.ER_SYNTAX_ERROR))
	}
	plan, err := sei.buildPlan(stmt)
	if err != nil {
		return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
	}
	rs, err := sei.executePlan(plan)
	if err != nil {
		return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, err.Error()))
	}
	// merge select rs
	return sei.mergeSelectResult(rs, stmt)
}

func (sei *session) mergeSelectResult(rs []*mysql.Result, stmt *sqlparser.Select) error {
	var r *mysql.Result
	var err error

	if len(stmt.GroupBy) == 0 {
		// only select
		r, err = sei.buildSelectResult(rs, stmt)
		if err != nil {
			return err
		}
	} else {
		// TODO: support group by
		return errors.New("not support")
	}

	err = sei.sortResultSet(r.Resultset, stmt)
	if err != nil {
		return err
	}

	// TODO: limit result

	sei.status = r.Status

	return sei.writeResultset(r.Status, r.Resultset)
}

func (sei *session) sortResultSet(set *mysql.Resultset, stmt *sqlparser.Select) error {
	if len(stmt.OrderBy) > 0 {
		glog.Info("Sort select results")
		ks := make([]mysql.SortKey, 0, len(stmt.OrderBy))
		for _, o := range stmt.OrderBy {
			ks = append(ks, mysql.SortKey{
				Direction: o.Direction,
				Name:      sqlparser.String(o.Expr),
			})
		}
		return set.Sort(ks)
	}
	return nil
}

// build select result without group by
func (sei *session) buildSelectResult(rs []*mysql.Result,
	stmt *sqlparser.Select) (*mysql.Result, error) {
	var err error
	r := rs[0].Resultset
	status := sei.status | rs[0].Status

	funcExprs := sei.getFuncExprs(stmt)
	if len(funcExprs) == 0 {
		for i := 1; i < len(rs); i++ {
			status |= rs[i].Status
			for j := range rs[i].Values {
				r.Values = append(r.Values, rs[i].Values[j])
				r.RowDatas = append(r.RowDatas, rs[i].RowDatas[j])
			}
		}
	} else {
		// result only one row, status doesn't need set
		r, err = sei.buildFuncExprResultSet(stmt, rs, funcExprs)
		if err != nil {
			return nil, err
		}
	}
	return &mysql.Result{
		Status:    status,
		Resultset: r,
	}, nil
}

func (sei *session) buildFuncExprResultSet(stmt *sqlparser.Select, rs []*mysql.Result, funcExprs map[int]string) (*mysql.Resultset, error) {
	var (
		names  []string
		err    error
		r      = rs[0].Resultset
		value  interface{}
		values []interface{}
	)

	for index, fn := range funcExprs {
		switch fn {
		case "sum", "count":
			value, err = sei.getSumFuncExprValue(rs, index)
		case "max":
			value, err = sei.getMaxFuncExprValue(rs, index)
		case "min":
			value, err = sei.getMinFuncExprValue(rs, index)
		default:
			err = errors.New("unknow function")
		}
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	// build result
	if len(values) == 0 {
		// empty
		r, err = sei.buildEmptySet([]string{"v"}, []interface{}{""})
	} else {
		// get names
		names = make([]string, 0, len(r.Fields))
		for _, name := range r.Fields {
			names = append(names, string(name.Name))
		}
		r, err = sei.buildResultset(r.Fields, names, [][]interface{}{values})
	}
	return r, err
}

// get the index of funcExpr, the value is function name
func (sei *session) getFuncExprs(stmt *sqlparser.Select) map[int]string {
	var f *sqlparser.FuncExpr
	funcExprs := make(map[int]string)

	for i, expr := range stmt.SelectExprs {
		nonStarExpr, ok := expr.(*sqlparser.NonStarExpr)
		if !ok {
			continue
		}

		f, ok = nonStarExpr.Expr.(*sqlparser.FuncExpr)
		if !ok {
			continue
		} else {
			f = nonStarExpr.Expr.(*sqlparser.FuncExpr)
			name, ok := funcNameMap[string(f.Name)]
			if ok {
				funcExprs[i] = name
			}
		}
	}
	return funcExprs
}

func (sei *session) handleSimpleSelect(stmt *sqlparser.Select) error {
	switch expr := stmt.SelectExprs[0].(type) {
	case *sqlparser.NonStarExpr:
		switch f := expr.Expr.(type) {
		case *sqlparser.FuncExpr:
			return sei.handleSimpleSelectFunc(stmt, f)
		}
	case *sqlparser.StarExpr:

	}
	return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "not support stmt"))
}

func (sei *session) handleSimpleSelectFunc(stmt *sqlparser.Select, f *sqlparser.FuncExpr) error {
	expr := stmt.SelectExprs[0].(*sqlparser.NonStarExpr)
	switch f.Name {
	case "database":
		if f.Exprs != nil {
			return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "not support stmt"))
		}
		// get now database
		var name []string
		if expr.As == "" {
			name = []string{"DATABASE()"}
		} else {
			name = []string{string(expr.As)}
		}
		var values [][]interface{}
		if sei.db == "" {
			values = [][]interface{}{
				[]interface{}{"NULL"},
			}
		} else {
			values = [][]interface{}{
				[]interface{}{sei.db},
			}
		}
		r, err := sei.buildResultset(nil, name, values)
		if err != nil {
			return err
		}
		return sei.writeResultset(sei.status, r)
	}
	return sei.writeError(mysql.NewError(mysql.ER_UNKNOWN_ERROR, "not support stmt"))
}

func (sei *session) getSumFuncExprValue(rs []*mysql.Result,
	index int) (interface{}, error) {
	var sumf float64
	var sumi int64
	var IsInt bool
	var err error
	var result interface{}

	for _, r := range rs {
		for k := range r.Values {
			result, err = r.GetValue(k, index)
			if err != nil {
				return nil, err
			}
			if result == nil {
				continue
			}

			switch v := result.(type) {
			case int:
				sumi = sumi + int64(v)
				IsInt = true
			case int32:
				sumi = sumi + int64(v)
				IsInt = true
			case int64:
				sumi = sumi + v
				IsInt = true
			case float32:
				sumf = sumf + float64(v)
			case float64:
				sumf = sumf + v
			case []byte:
				tmp, err := strconv.ParseFloat(string(v), 64)
				if err != nil {
					return nil, err
				}

				sumf = sumf + tmp
			default:
				return nil, errors.New("error col type")
			}
		}
	}
	if IsInt {
		return sumi, nil
	}
	return sumf, nil
}

func (sei *session) getMaxFuncExprValue(rs []*mysql.Result, index int) (interface{}, error) {
	return sei.funcCompare(rs, index, func(vaule, result interface{}) bool {
		switch result.(type) {
		case int64:
			return vaule.(int64) < result.(int64)
		case uint64:
			return vaule.(uint64) < result.(uint64)
		case float64:
			return vaule.(float64) < result.(float64)
		case string:
			return vaule.(string) < result.(string)
		default:
			return false
		}
	})
}

func (sei *session) getMinFuncExprValue(rs []*mysql.Result, index int) (interface{}, error) {
	return sei.funcCompare(rs, index, func(vaule, result interface{}) bool {
		switch result.(type) {
		case int64:
			return vaule.(int64) > result.(int64)
		case uint64:
			return vaule.(uint64) > result.(uint64)
		case float64:
			return vaule.(float64) > result.(float64)
		case string:
			return vaule.(string) > result.(string)
		default:
			return false
		}
	})
}

func (sei *session) funcCompare(rs []*mysql.Result, index int, f func(v interface{}, n interface{}) bool) (interface{}, error) {
	var value interface{}
	var findNotNull bool
	if len(rs) == 0 {
		return nil, nil
	}
	// get init value
	for _, r := range rs {
		for k := range r.Values {
			result, err := r.GetValue(k, index)
			if err != nil {
				return nil, err
			}
			if result != nil {
				value = result
				findNotNull = true
				break
			}
		}
		if findNotNull {
			break
		}
	}
	for _, r := range rs {
		for k := range r.Values {
			result, err := r.GetValue(k, index)
			if err != nil {
				return nil, err
			}
			if result == nil {
				continue
			}
			if f(value, result) {
				value = result
			}
		}
	}
	return value, nil
}
