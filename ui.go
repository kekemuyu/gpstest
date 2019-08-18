package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"

	log "github.com/donnie4w/go-logger/logger"

	"test/gpstest/gps"

	"github.com/zserge/lorca"
)

func updateGps(ui lorca.UI, msg gps.Msg) {
	jsStr := fmt.Sprintf(`$('#utctime').text('%s');
					   $('#state').text('%s');
					   $('#latitude').text('%s');
					   $('#latidudeDirection').text('%s');
					   $('#longitude').text('%s');
					   $('#longitudeDirection').text('%s');
					   $('#speed').text('%s');
					   $('#speedDirection').text('%s');
					   $('#utcdate').text('%s');`,
		msg.Utctime, msg.State, msg.Latitude, msg.LatitudeDirection, msg.Longitude, msg.LongitudeDirecting,
		msg.Speed, msg.SpeedDirection, msg.Utcdate)
	ui.Eval(jsStr)
}

func updateUI(ui lorca.UI) {
	for {
		select {
		case msg := <-gps.DefaultGps.Outdata:
			updateGps(ui, msg)
		}
	}
}

type Myweb struct {
	UI lorca.UI
}

func New(width, height int) Myweb {
	var myweb Myweb
	var err error
	myweb.UI, err = lorca.New("", "", width, height, "--no-sandbox")
	if err != nil {
		log.Fatal(err)
	}

	return myweb
}

func (m *Myweb) Run() {

	ui := m.UI
	defer ui.Close()

	// A simple way to know when UI is ready (uses body.onload even in JS)
	ui.Bind("start", func() {
		log.Debug("UI is ready")
	})

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(FS))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))
	go updateUI(ui)
	//	go recieve(ui)
	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	// ui.Eval(`
	// 	console.log("Hello, world!");
	// 	console.log('Multiple values:', [1, false, {"x":5}]);
	// `)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Debug("exiting...")
}
