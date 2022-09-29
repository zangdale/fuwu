package controller

import (
	"context"
	"fuwu/conf"
	"fuwu/server"
	"log"
	"net/http"
	"runtime"
)

var (
	fw *conf.FuWu
)

func DaemonServer(ctx context.Context, cf *conf.Config) {
	var err error
	fw, err = conf.GetFuwu()
	if err != nil {
		log.Fatal(err.Error())
	}

	hander := server.NewDaemon(ctx, fw)
	hander.Log = cf.Log
	for {
		reserver(ctx, cf, hander)
	}

}

func reserver(ctx context.Context, conf *conf.Config, hander http.Handler) {
	defer func() {
		if e := recover(); e != nil {
			buf := make([]byte, 1024*2)
			runtime.Stack(buf, true)
			conf.Log.Printf("panic reserver error code: %v\n%s", e, string(buf))
		}
	}()
	if err := http.ListenAndServe(conf.DaemonAddr, hander); err != nil {
		conf.Log.Println(err.Error())
	}
}
