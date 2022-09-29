package controller

import (
	"fmt"
	"fuwu/server"
)

var (
	ErrServerName = fmt.Errorf("server name is required")
	ErrConfigJson = fmt.Errorf("must have config.json")
)

type CliArgs struct {
	// cli
	ShowList          bool
	RefreshServerList bool
	// cli server
	Start      bool
	Stop       bool
	Status     bool
	Restart    bool
	ServerName string
	DaemonAddr string
}

func (args CliArgs) Run() error {
	if args.DaemonAddr == "" {
		return ErrConfigJson
	}

	// cli
	switch {
	case args.ShowList:
		return server.ActionServerList.CliRun("", args.DaemonAddr)
	case args.RefreshServerList:
		return server.ActionServerRefresh.CliRun("", args.DaemonAddr)
	}

	// cli server
	if args.ServerName == "" {
		return ErrServerName
	}

	switch {
	case args.Start:
		return server.ActionServerStart.CliRun(args.ServerName, args.DaemonAddr)
	case args.Stop:
		return server.ActionServerStop.CliRun(args.ServerName, args.DaemonAddr)
	case args.Status:
		return server.ActionServerStatus.CliRun(args.ServerName, args.DaemonAddr)
	case args.Restart:
		return server.ActionServerRestart.CliRun(args.ServerName, args.DaemonAddr)
	}

	return nil
}
