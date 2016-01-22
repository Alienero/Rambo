package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func AST_HTTP() {
	http.HandleFunc("/ast", func(w http.ResponseWriter, r *http.Request) {
		// get sql
		sql, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println("sql is:", string(sql))
		stmt, err := sqlparser.Parse(string(sql))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ntree := sqlparser.NewTree()
		ntree.SetTree(stmt)
		data, err := json.Marshal(ntree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
	})
	gopath := os.Getenv("GOPATH")
	fs := http.FileServer(http.Dir(path.Join(gopath, "src", "github.com", "Alienero", "Rambo", "front")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

// select * from
// (
// select tb1.t c1 ,tb2.t c2 ,tb3.t c3 ,tb4.t c4 ,tb5.t c5 ,tb6.t c6 ,tb7.t c7 ,tb8.t c8 ,tb9.t c9 ,tb10.t c10,tb11.t c11,tb12.t c12,tb13.t c13,tb14.t c14,tb15.t c15,tb16.t c16,tb17.t c17,tb18.t c18,tb19.t c19,tb20.t c20,tb21.t c21,tb22.t c22,tb23.t c23,tb24.t c24,tb25.t c25,tb26.t c26,tb27.t c27,tb28.t c28,tb29.t c29,tb30.t c30,tb31.t c31,tb32.t c32,tb33.t c33,tb34.t c34,tb35.t c35,tb36.t c36,tb37.t c37,tb38.t c38,tb39.t c39,tb40.t c40,tb41.t c41,tb42.t c42,tb43.t c43,tb44.t c44,tb45.t c45,tb46.t c46,tb47.t c47,tb48.t c48,tb49.t c49,tb50.t c50,tb51.t c51,tb52.t c52,tb53.t c53,tb54.t c54,tb55.t c55,tb56.t c56,tb57.t c57,tb58.t c58,tb59.t c59,tb60.t c60,tb61.t c61,tb62.t c62,tb63.t c63,tb64.t c64
// from
// tb tb1  cross join tb tb2 cross join tb tb3 cross join tb tb4 cross join tb tb5 cross join tb tb6 cross join tb tb7 cross join tb tb8 cross join tb tb9 cross join tb tb10 cross join tb tb11 cross join tb tb12 cross join tb tb13 cross join tb tb14 cross join tb tb15 cross join tb tb16 cross join tb tb17 cross join tb tb18 cross join tb tb19 cross join tb tb20 cross join tb tb21 cross join tb tb22 cross join tb tb23 cross join tb tb24 cross join tb tb25 cross join tb tb26 cross join tb tb27 cross join tb tb28 cross join tb tb29 cross join tb tb30 cross join tb tb31 cross join tb tb32 cross join tb tb33 cross join tb tb34 cross join tb tb35 cross join tb tb36 cross join tb tb37 cross join tb tb38 cross join tb tb39 cross join tb tb40 cross join tb tb41 cross join tb tb42 cross join tb tb43 cross join tb tb44 cross join tb tb45 cross join tb tb46 cross join tb tb47 cross join tb tb48 cross join tb tb49 cross join tb tb50 cross join tb tb51 cross join tb tb52 cross join tb tb53 cross join tb tb54 cross join tb tb55 cross join tb tb56 cross join tb tb57 cross join tb tb58 cross join tb tb59 cross join tb tb60 cross join tb tb61 cross join tb tb62 cross join tb tb63 cross join tb tb64
// ) a
// where
// c1 +c2 +c3 +c4 +c5 +c6 +c7 +c8  in(0,1) and c9 +c10+c11+c12+c13+c14+c15+c16 in(0,1) and c17+c18+c19+c20+c21+c22+c23+c24 in(0,1) and c25+c26+c27+c28+c29+c30+c31+c32 in(0,1) and c33+c34+c35+c36+c37+c38+c39+c40 in(0,1) and c41+c42+c43+c44+c45+c46+c47+c48 in(0,1) and c49+c50+c51+c52+c53+c54+c55+c56 in(0,1) and c57+c58+c59+c60+c61+c62+c63+c64 in(0,1)
// and
// c1 +c9 +c17+c25+c33+c41+c49+c57 in(0,1) and c2 +c10+c18+c26+c34+c42+c50+c58 in(0,1) and c3 +c11+c19+c27+c35+c43+c51+c59 in(0,1) and c4 +c12+c20+c28+c36+c44+c52+c60 in(0,1) and c5 +c13+c21+c29+c37+c45+c53+c61 in(0,1) and c6 +c14+c22+c30+c38+c46+c54+c62 in(0,1) and c7 +c15+c23+c31+c39+c47+c55+c63 in(0,1) and c8 +c16+c24+c32+c40+c48+c56+c64 in(0,1)
// and
// c1  in(0,1) and c16+c23+c30+c37+c44+c51+c58 in(0,1) and c2 +c9  in(0,1) and c24+c31+c38+c45+c52+c59 in(0,1) and c3 +c10+c17 in(0,1) and c32+c39+c46+c53+c60 in(0,1) and c4 +c11+c18+c25 in(0,1) and c40+c47+c54+c61 in(0,1) and c5 +c12+c19+c26+c33 in(0,1) and c48+c55+c62 in(0,1) and c6 +c13+c20+c27+c34+c41 in(0,1) and c56+c63 in(0,1) and c7 +c14+c21+c28+c35+c42+c49 in(0,1) and c64 in(0,1) and c8 +c15+c22+c29+c36+c43+c50+c57 in(0,1)
// and
// c1 +c10+c19+c28+c37+c46+c55+c64 in(0,1) and c9 +c18+c27+c36+c45+c54+c63 in(0,1) and c2 +c11+c20+c29+c38+c47+c56 in(0,1) and c17+c26+c35+c44+c53+c62 in(0,1) and c3 +c12+c21+c30+c39+c48 in(0,1) and c25+c34+c43+c52+c61 in(0,1) and c4 +c13+c22+c31+c40 in(0,1) and c33+c42+c51+c60 in(0,1) and c5 +c14+c23+c32 in(0,1) and c41+c50+c59 in(0,1) and c6 +c15+c24 in(0,1) and c49+c58 in(0,1) and c7 +c16 in(0,1) and c57 in(0,1) and c8  in(0,1)
// and
// c1+c2+c3+c4+c5+c6+c7+c8+c9+c10+c11+c12+c13+c14+c15+c16+c17+c18+c19+c20+c21+c22+c23+c24+c25+c26+c27+c28+c29+c30+c31+c32+c33+c34+c35+c36+c37+c38+c39+c40+c41+c42+c43+c44+c45+c46+c47+c48+c49+c50+c51+c52+c53+c54+c55+c56+c57+c58+c59+c60+c61+c62+c63+c64=8
