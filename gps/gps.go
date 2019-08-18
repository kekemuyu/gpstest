package gps

import (
	"strings"

	log "github.com/donnie4w/go-logger/logger"
)

type Msg struct {
	MsgId              string
	Utctime            string
	State              string
	Latitude           string
	LatitudeDirection  string
	Longitude          string
	LongitudeDirecting string
	Speed              string
	SpeedDirection     string
	Utcdate            string
}
type Gps struct {
	Buf     string
	Data    Msg
	Outdata chan Msg
}

var DefaultGps = Gps{
	Outdata: make(chan Msg, 10),
}

func (g Gps) getSubStr(str, start, end string) string {
	n := strings.Index(str, start)
	if n == -1 {
		n = 0
		return ""
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
		return ""
	}
	str = string([]byte(str)[:m])
	return str
}
func (g Gps) Read(data []byte) {
	if len(DefaultGps.Buf) < 300 {
		DefaultGps.Buf = DefaultGps.Buf + string(data)
	} else {
		substr := g.getSubStr(DefaultGps.Buf, "$GNRMC", "\r\n")
		gpsstr := strings.Split(substr, ",")

		DefaultGps.Data.MsgId = gpsstr[0]
		DefaultGps.Data.Utctime = gpsstr[1]
		DefaultGps.Data.State = gpsstr[2]
		DefaultGps.Data.Latitude = gpsstr[3]
		DefaultGps.Data.LatitudeDirection = gpsstr[4]
		DefaultGps.Data.Longitude = gpsstr[5]
		DefaultGps.Data.LongitudeDirecting = gpsstr[6]
		DefaultGps.Data.Speed = gpsstr[7]
		DefaultGps.Data.SpeedDirection = gpsstr[8]
		DefaultGps.Data.Utcdate = gpsstr[9]
		log.Debug(gpsstr)
		DefaultGps.Outdata <- DefaultGps.Data
		DefaultGps.Buf = ""
	}
}

func (g Gps) Write(data []byte) {

}
