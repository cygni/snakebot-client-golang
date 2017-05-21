package main

import (
	"flag"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/url"
)

var (
	host  = kingpin.Flag("host", "Host to connect to").Short('r').Default("localhost").String()
	port  = kingpin.Flag("port", "Port to connect to").Short('p').Default("8080").String()
	venue = kingpin.Flag("venue", "Venue to connect to").Short('v').Default("training").Enum("training", "tournament")
)

func main() {
	kingpin.Parse()
	host := *host
	port := *port
	venue := *venue

	log.SetLevel(log.DebugLevel)
	formatter := new(prefixed.TextFormatter)

	log.SetFormatter(formatter)

	clientLogger := log.WithFields(log.Fields{"host": host, "port": port, "venue": venue, "prefix": "client"})
	clientLogger.Info("Connecting to websocket")

	addr := flag.String("addr", host+":"+port, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/" + venue}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		clientLogger.Fatal("Error connecting to websocket:", err)
	}

	clientLogger.Debug("Websocket is connected!")

	defer c.Close()
	clientLogger.Info("All done, closing down")
}
