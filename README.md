# go-kratos-log-syslog

[![Go Version](https://badgen.net/github/release/go-packagist/go-kratos-log-syslog/stable)](https://github.com/go-packagist/go-kratos-log-syslog/releases)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-packagist/go-kratos-log-syslog)](https://pkg.go.dev/github.com/go-packagist/go-kratos-log-syslog)
[![codecov](https://codecov.io/gh/go-packagist/go-kratos-log-syslog/branch/master/graph/badge.svg?token=5TWGQ9DIRU)](https://codecov.io/gh/go-packagist/go-kratos-log-syslog)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-packagist/go-kratos-log-syslog)](https://goreportcard.com/report/github.com/go-packagist/go-kratos-log-syslog)
[![tests](https://github.com/go-packagist/go-kratos-log-syslog/actions/workflows/go.yml/badge.svg)](https://github.com/go-packagist/go-kratos-log-syslog/actions/workflows/go.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/go-packagist/go-kratos-log-syslog
```

## Usage

```go
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
```

## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.