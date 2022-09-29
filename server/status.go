package server

import (
	"context"
	"fuwu/conf"
	"fuwu/utils"
	"io"
	"log"
	"sync"
)

var _ io.Writer = (*ServerStatus)(nil)

type ServerStatus struct {
	rw sync.RWMutex

	count  int32
	logMsg [][]byte

	server *conf.Server
	Ctx    context.Context
	cancel context.CancelFunc
	log    *log.Logger
}

func NewServerStatus(ctx context.Context, cf *conf.Server) *ServerStatus {
	ctx, cancel := context.WithCancel(ctx)
	wc := utils.NewLog(cf.LogDir)
	return &ServerStatus{
		rw:     sync.RWMutex{},
		server: cf,
		Ctx:    ctx,
		cancel: cancel,
		log:    wc,
	}

}

func (s ServerStatus) Write(b []byte) (int, error) {
	if s.server.LogMsgCount != 0 && len(b) > 0 {
		s.rw.Lock()
		if s.count >= s.server.LogMsgCount {
			s.logMsg = append(s.logMsg[1:], b)
		} else {
			s.logMsg = append(s.logMsg, b)
			s.count++
		}
		s.rw.Unlock()
	}

	return s.log.Writer().Write(b)
}

func (s ServerStatus) GetTempLog() [][]byte {
	s.rw.RLock()
	defer s.rw.RUnlock()
	var res [][]byte
	for i := range s.logMsg {
		res = append(res, s.logMsg[i])
	}
	return res
}

func (s ServerStatus) Close() error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}
