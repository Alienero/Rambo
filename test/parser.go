package main

import (
	"log"
	"reflect"

	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func main() {
	sql := `select t.id from t`
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}
	t := reflect.TypeOf(stmt)

	log.Println(t)
	slc := stmt.(*sqlparser.Select)
	s := slc.SelectExprs[0].(*sqlparser.NonStarExpr).Expr
	log.Println(s.(*sqlparser.ColName).Qualifier)
}
