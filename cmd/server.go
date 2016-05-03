package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/server"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start rambo server",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		if debug {
			flag.Set("logtostderr", "true")
			flag.Set("v", "10")
		}
		defer glog.Flush()
		// default etcd addrs
		config.Config.Etcd.EtcdAddr = []string{"http://192.168.99.100:4001"}
		// default mysql port
		config.Config.Server.ListenAddr = "localhost:3306"
		s := server.NewSever()
		// listenning
		go s.Run()
		go s.StartDDLMange()
		// register this server
		stop, err := s.GetInfo().CreateAndHeartBeat(meta.ProxyNodes, config.Config.Server.ListenAddr,
			config.Config.Etcd.UpdateTTL, true)
		if err != nil {
			glog.Fatal(err)
		}
		defer stop()
		// signal
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
		<-signalCh
		glog.Info("Signal received, initializing clean shutdown...")
	},
}

var (
	cfgFile string
	debug   = true
)

func init() {
	RootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	serverCmd.Flags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Rambo.yaml)")

	initConfig()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".Rambo") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
