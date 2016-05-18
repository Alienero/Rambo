package cmd

import (
	"flag"
	"os"
	"os/signal"
	"path"
	"syscall"

	"github.com/Alienero/Rambo/config"
	"github.com/Alienero/Rambo/meta"
	"github.com/Alienero/Rambo/server"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "start rambo server",
	Run: func(cmd *cobra.Command, args []string) {
		// initConfig reads in config file.
		if err := config.InitConfig(cfgFile); err != nil {
			panic(err)
		}
		flag.Parse()
		if config.Config.IsDev {
			flag.Set("logtostderr", "true")
			flag.Set("v", "10")
		}
		defer glog.Flush()
		flag.Set("log_dir", config.Config.LogFileDir)

		s := server.NewSever()
		// listenning
		go s.Run()
		go s.StartDDLMange()
		// register this server
		stop, err := s.GetInfo().CreateAndHeartBeat(path.Join(meta.ProxyNodes, config.Config.Server.RPCAddr), config.Config.Server.RPCAddr,
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

	serverCmd.Flags().StringVar(&cfgFile, "config", "config.toml", "config file (default is rambo.toml)")
}
