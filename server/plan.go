package server

import (
	"errors"
	"fmt"
	"strconv"

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

func (p *Plan) generateSQL(statement sqlparser.Statement) error {
	switch stmt := statement.(type) {
	case *sqlparser.Insert:
		// return sei.buildInsertPlan(stmt)
	case *sqlparser.Select:
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
	case *sqlparser.Update:
		// return sei.buildUpdatePlan(stmt)
	case *sqlparser.Delete:
		// return sei.buildDeletePlan(stmt)
	}
	return nil
}

// limit
func (plan *Plan) rewriteLimit(stmt *sqlparser.Select) (err error) {

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

	plan.offset = offset
	plan.count = count
	return
}

//build a router plan
func (sei *session) buildPlan(statement sqlparser.Statement) (*Plan, error) {
	switch stmt := statement.(type) {
	case *sqlparser.Insert:
		// return sei.buildInsertPlan(stmt)
	case *sqlparser.Select:
		return sei.buildSelectPlan(stmt)
	case *sqlparser.Update:
		// return sei.buildUpdatePlan(stmt)
	case *sqlparser.Delete:
		// return sei.buildDeletePlan(stmt)
	}
	return nil, errors.New("not support this plan")
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
	if err != nil {
		return nil, err
	}
	return plan, nil
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
