package syslog

import (
	"testing"

	"github.com/go-kratos/kratos/v2/log"
)

func TestSyslogLogger(t *testing.T) {
	logger, err := New(&Config{
		Network: "udp",
		Tag:     "test",
	})
	defer logger.Close()

	if err != nil {
		t.Fatal(err)
	}

	err = logger.Log(log.LevelDebug, "test", "test")
	if err != nil {
		t.Fatal(err)
	}
}
