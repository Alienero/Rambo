package main

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func main() {
	sql := `select * from user where id = 1 order by dec limit 0,10`
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		panic(err)
	}
	t := reflect.TypeOf(stmt)

	log.Println(t, "name:", t.Name())

	s := stmt.(*sqlparser.Select)

	// buf
	buf := sqlparser.NewTrackedBuffer(nil)
	s.Format(buf)
	log.Println(buf)

	ntree := sqlparser.NewTree()
	ntree.SetTree(s)
	log.Println(ntree)
	data, err := json.MarshalIndent(ntree, "", "\t")
	if err != nil {
		panic(err)
	}
	log.Println(string(data))
}

type d string

func f(i interface{}) interface{} {
	return i
}
