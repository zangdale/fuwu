package server

import (
	"context"
	"fuwu/conf"
	"log"
	"net/http"
)

type DaemonServer struct {
	fw          *conf.FuWu
	Log         *log.Logger
	ctx         context.Context
	runningList map[string]*ServerStatus
}

var _ http.Handler = (*DaemonServer)(nil)

func NewDaemon(ctx context.Context, fw *conf.FuWu) *DaemonServer {
	if ctx == nil {
		ctx = context.Background()
	}
	return &DaemonServer{
		fw:          fw,
		ctx:         ctx,
		runningList: make(map[string]*ServerStatus),
		Log:         log.Default(),
	}
}

func (s *DaemonServer) ReloadConf() error {
	fw, err := conf.GetFuwu()
	if err != nil {
		return err
	}
	s.fw = fw
	return nil
}

func (s *DaemonServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NewServerStatus()
}
