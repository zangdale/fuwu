package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type ActionArgType string

func (a ActionArgType) String() string {
	return string(a)
}

const (
	ActionServerList    ActionArgType = "list"
	ActionServerRefresh ActionArgType = "refresh"

	ActionServerStart   ActionArgType = "start"
	ActionServerStop    ActionArgType = "stop"
	ActionServerStatus  ActionArgType = "status"
	ActionServerRestart ActionArgType = "restart"
)

func (a ActionArgType) CliRun(serverName string, daemonAddr string) error {
	resp, err := http.DefaultClient.Get(func() string {
		if serverName == "" {
			return fmt.Sprintf("http://%s/%s", daemonAddr, a.String())
		}
		return fmt.Sprintf("http://%s/%s/%s", daemonAddr, a.String(), serverName)
	}())
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	return err
}
