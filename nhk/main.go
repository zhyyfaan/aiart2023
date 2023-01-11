package main

import (
	"aiart2023/pkg/global"
	"aiart2023/pkg/initialization"
	"aiart2023/pkg/util"
)

func main() {
	initialization.Init()

	if err := global.GlobalEngine.Run(":8080"); err != nil {
		util.ErrorLogger.Panicf("启动异常, err: %s", err.Error())
	}
}
