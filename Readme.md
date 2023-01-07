# yaglogger (Yet another Go Logger)
[![GoDoc](https://godoc.org/github.com/jlentink/yaglogger?status.svg)](https://godoc.org/github.com/0xAX/yaglogger)
[![Go Report Card](https://goreportcard.com/badge/github.com/jlentink/yaglogger)](https://goreportcard.com/report/github.com/jlentink/yaglogger)
[![Coverage Status](https://coveralls.io/repos/github/jlentink/yaglogger/badge.svg?branch=master)](https://coveralls.io/github/jlentink/yaglogger?branch=master)

yaglogger is a simple logger for Go. It allows to log in color to screen and to a file.

## Installation

```bash
 go get -u github.com/jlentink/yaglogger
```

## Direct usage
```go
package main

import (
	log "github.com/jlentink/yaglogger"
)

func main() {
	log.Debug("Debug")
}
```
 
## Default logger

```go
package main

import (
	"github.com/jlentink/yaglogger"
)

func main() {
    logger := yaglogger.NewLogger()
    logger.Info("Hello world")
}
```

## Custom logger

Or adjust where needed by initializing a new logger with the following options. Referer NewLogger for more information.

```go
package main

import (
	"github.com/jlentink/yaglogger"
)
func main() {
        yalogger.Logger {
            Level:       0,
            OutErr:      nil,
            OutNormal:   nil,
            Output:      yalogger.LevelOutput{},
            ShowLevel:   false,
            ShowDate:    false,
            Color:       false,
            LogToScreen: false,
            LogFilePath: "",
            LogFile:     nil,
        }
	}
```