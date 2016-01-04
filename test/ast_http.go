package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/Alienero/Rambo/mysql/sqlparser"
)

func main() {
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
