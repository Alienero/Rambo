package server

import (
	"fmt"
	"testing"

	"github.com/pingcap/tidb/ast"
	"github.com/pingcap/tidb/parser"

	ftype "github.com/pingcap/tidb/mysql"
)

func TestDDL(t *testing.T) {
	sql := `
	Create table test5 (id int PRIMARY KEY AUTO_INCREMENT,b int not null,UNIQUE(b),info varchar(100),c int,  cc bigint(64) unsigned )
	`
	stmt, err := parser.ParseOneStmt(sql, "", "")
	if err != nil {
		t.Log("parser error", err.Error())
		return
	}
	tc := stmt.(*ast.CreateTableStmt)
	fmt.Println(len(tc.Constraints), len(tc.Constraints[0].Keys), tc.Constraints[0].Keys[0].Column.Name.String())
	fmt.Println(len(tc.Cols[0].Options), tc.Cols[0].Tp, tc.Cols[0].Options[0].Tp)
	fmt.Println(tc.Cols[0].Tp.Tp, ftype.TypeBlob)

}
