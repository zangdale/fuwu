package main

import (
	"context"
	"flag"
	"fuwu/conf"
	"fuwu/controller"
	"fuwu/server"
	"log"
	"os"
)

var (
	// daemon
	serverDaemon = flag.Bool("server", false, "start daemon server")
	serverLogger = flag.Bool("log", false, "show terminal log")
)

var (
	// cli
	cliList    = flag.Bool("list", false, "show all serverName and status is running")
	cliRefresh = flag.Bool("refresh", false, "read fuwu.yml and refresh server list")
	// cli server
	cliStart   = flag.String(server.ActionServerStart.String(), "", "start serverName")
	cliStop    = flag.String(server.ActionServerStop.String(), "", "stop serverName")
	cliStatus  = flag.String(server.ActionServerStatus.String(), "", "status serverName")
	cliRestart = flag.String(server.ActionServerRestart.String(), "", "restart serverName")
)

var (
	cf *conf.Config
)

func init() {
	var err error
	cf, err = conf.GetConfig()
	if err != nil {
		log.Panicln(err.Error())
	}
}

func main() {
	flag.Parse()
	//daemon
	daemonServer()
	// cli
	err := cli()
	if err != nil {
		log.Println(err.Error())
	}

}
func daemonServer() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if *serverDaemon {
		if *serverLogger {
			cf.SetLogWriter(os.Stdout)
		}
		controller.DaemonServer(ctx, cf)
		os.Exit(1)
	}
	return nil
}

func cli() error {
	args := controller.CliArgs{}

	switch {
	case *cliList:
		args.ShowList = true
	case *cliRefresh:
		args.RefreshServerList = true
	case *cliStart != "":
		args.Start = true
		args.ServerName = *cliStart
	case *cliStop != "":
		args.Stop = true
		args.ServerName = *cliStop
	case *cliStatus != "":
		args.Status = true
		args.ServerName = *cliStatus
	case *cliRestart != "":
		args.Restart = true
		args.ServerName = *cliRestart
	}

	// cli server
	args.DaemonAddr = cf.DaemonAddr

	return args.Run()
}
