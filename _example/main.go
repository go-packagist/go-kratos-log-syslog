package main

import (
	"github.com/go-kratos/kratos/v2/log"
	syslog "github.com/go-packagist/go-kratos-log-syslog"
)

func main() {
	s, err := syslog.New(&syslog.Config{
		Network: "udp",
		Tag:     "test",
		Addr:    "192.168.8.92:30732",
	})
	if err != nil {
		panic(err)
	}
	defer s.Close()

	s.Log(log.LevelDebug, "test", "test")
}
