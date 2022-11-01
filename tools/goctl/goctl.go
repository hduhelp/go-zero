package main

import (
	"github.com/hduhelp/go-zero/core/load"
	"github.com/hduhelp/go-zero/core/logx"
	"github.com/hduhelp/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
