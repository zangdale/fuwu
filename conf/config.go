package conf

import (
	"io"
	"log"
)

type Config struct {
	DaemonAddr string         `json:"daemon_addr"`
	logWriter  io.WriteCloser `json:"-"`
	Log        *log.Logger    `json:"-"`
}

func GetConfig() (*Config, error) {
	return &Config{
		Log:        log.Default(),
		DaemonAddr: ":555",
	}, nil
}

func (c Config) SetLogWriter(w io.WriteCloser) {
	c.logWriter = w
	c.Log.SetOutput(c.logWriter)
}
