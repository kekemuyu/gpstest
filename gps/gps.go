package gps

import (
	log "github.com/donnie4w/go-logger/logger"
)

type Gps struct {
}

var DefaultGps = Gps{}

func (g Gps) Read(data []byte) {
	log.Debug(data)
}

func (g Gps) Write(data []byte) {

}
