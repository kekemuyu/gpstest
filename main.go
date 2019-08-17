//go:generate go run -tags generate gen.go
package main

import (
	"test/gpstest/com"
	"test/gpstest/config"
	"test/gpstest/gps"

	log "github.com/donnie4w/go-logger/logger"
)

func main() {
	portName := config.Cfg.Section("serial").Key("PortName").MustString("com1")
	baund := config.Cfg.Section("serial").Key("BaudRate").MustInt()
	defaultCom, err := com.New(portName, uint(baund))

	if err != nil {
		log.Error("打开串口错误：", err)
		//		return
	} else {
		go defaultCom.Run(gps.DefaultGps)
	}

	myweb := New(800, 560)
	myweb.Run()
}
