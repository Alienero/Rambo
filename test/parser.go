package main

import (
	"log"
	"reflect"

	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func main() {
	sql := `
	explain`
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}
	t := reflect.TypeOf(stmt)
	log.Println(t)
	s := stmt.(*sqlparser.Explain)

	// s := stmt.(*sqlparser.Select)
	// v := s.From[0].(*sqlparser.AliasedTableExpr)
	// // from
	// log.Println(sqlparser.String(v.Expr))
	// // limit
	// log.Println(s.Limit.Offset, s.Limit.Rowcount)
	// // order
	// log.Println(s.OrderBy[0].Expr.(*sqlparser.ColName).Name)
	// // where
	// where := s.Where.Expr.(*sqlparser.ComparisonExpr)
	// log.Println(where.Left, where.Operator, where.Right)
	// // colums
	// log.Println(s.SelectExprs[0].(*sqlparser.NonStarExpr).Expr)

	// // buf
	buf := sqlparser.NewTrackedBuffer(nil)
	s.Format(buf)
	log.Println(buf)
}
