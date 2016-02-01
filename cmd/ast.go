package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/Alienero/Rambo/mysql/sqlparser"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var astCmd = &cobra.Command{
	Use:   "ast",
	Short: "start rambo ast web front",
	Run: func(cmd *cobra.Command, args []string) {
		listen, _ := cmd.Flags().GetString("http")

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
		fmt.Println("listen in:", listen)
		http.ListenAndServe(listen, nil)
	},
}

func init() {
	RootCmd.AddCommand(astCmd)
	astCmd.Flags().String("http", "localhost:8080", "-http=locahost:8080")
}
