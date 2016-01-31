package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alienero/Rambo/admin"
	"github.com/Alienero/Rambo/meta"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "manage user",
}

func init() {
	RootCmd.AddCommand(adminCmd)
	adminCmd.AddCommand(addUserCmd)
	addUserCmd.Flags().StringVar(&etcd, "etcd", "http://localhost:2379", "etcd machines (default is http://localhost:2379)")
}

var etcd string

var addUserCmd = &cobra.Command{
	Use:     "addUser",
	Short:   "add user",
	Example: "rambo admin addUser user 123",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			panic("bad input")
		}
		meta.InitMetaDB(strings.Split(etcd, ","))
		manage := admin.Admin{}
		if err := manage.AddUser(args[0], args[1]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		result, err := manage.GetUser(args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
	},
}
