package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"github.com/hokaccha/go-prettyjson"
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"gopkg.in/alecthomas/kingpin.v2"
	"net/url"
	"time"
)

var (
	host     = kingpin.Flag("host", "Host to connect to").Short('r').Default("localhost").String()
	port     = kingpin.Flag("port", "Port to connect to").Short('p').Default("8080").String()
	venue    = kingpin.Flag("venue", "Venue to connect to").Short('v').Default("training").Enum("training", "tournament")
	logLevel = kingpin.Flag("log-level", "Logging level to use").Short('l').Default("info").Enum(log.DebugLevel.String(), log.InfoLevel.String(), log.WarnLevel.String(), log.ErrorLevel.String(), log.FatalLevel.String())
)

func main() {
	kingpin.Parse()

	logLevel, _ := log.ParseLevel(*logLevel)
	log.SetLevel(logLevel)
	formatter := new(prefixed.TextFormatter)
	log.SetFormatter(formatter)

	conn := Connection{}
	conn.Start(*host, *port, *venue)

	log.Info("All done, closing down")
}

type Connection struct {
	logger     *log.Entry
	conn       *websocket.Conn
	isTraining bool
	done       chan struct{}
	messages   chan Message
	heartBeat  chan string
}

type Message struct {
	t     interface{}
	bytes []byte
}

func (c *Connection) connect(host string, port string, venue string) {
	c.logger = log.WithFields(log.Fields{"host": host, "port": port, "venue": venue, "prefix": "client"})

	addr := flag.String("addr", host+":"+port, "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/" + venue}

	c.logger.Info("Connecting to websocket")
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		c.logger.WithError(err).Fatal("Error connecting to websocket")
	}

	c.conn = conn
	c.done = make(chan struct{}, 1)
}

func (c *Connection) Start(host string, port string, venue string) {
	c.connect(host, port, venue)
	c.messages = make(chan Message, 1)

	c.conn.WriteJSON(RegisterPlayerMessage("go-snake"))
	c.conn.WriteJSON(ClientInfoMessage())
	go c.readMessages()

	for {
		select {
		case message := <-c.messages:
			c.routeMessage(message)
		case playerId := <-c.heartBeat:
			c.sendHeartBeat(playerId)
		case <-c.done:
			return
		}
	}
}

func (c *Connection) readMessages() {
	for {
		msg := c.nextMessage()
		if msg != nil {
			c.messages <- *msg
		}
	}
}

func (c Connection) nextMessage() *Message {
	msgType, bytes, err := c.conn.ReadMessage()
	if err != nil {
		c.logger.WithError(err).Error("Got an error reading the next message from websocket.")
		c.close()
		return nil
	}

	if msgType == websocket.BinaryMessage {
		c.logger.Error("Received unexpected binary message!")
		return nil
	}

	var f interface{}
	err = json.Unmarshal(bytes, &f)
	if err != nil {
		c.logger.WithError(err).WithFields(log.Fields{
			"msg": string(bytes),
		}).Error("Unable to parse incoming message as JSON")
		return nil
	}

	m := f.(map[string]interface{})

	return &Message{
		t:     m["type"],
		bytes: bytes,
	}

}

func (c *Connection) routeMessage(msg Message) {
	switch msg.t {
	case PLAYER_REGISTERED:
		var m PlayerRegistered
		c.messageToJson(&m, msg.bytes)
		c.conn.WriteJSON(StartGameMessage())

		c.isTraining = m.GameMode == "TRAINING"
		if !c.isTraining {
			go c.startHeartBeat(m.ReceivingPlayerId)
		}
	case GAME_LINK_EVENT:
		var m GameLink
		c.messageToJson(&m, msg.bytes)
		c.logger.WithFields(log.Fields{
			"url": m.Url,
		}).Info("Received link to watch game")
	case GAME_STARTING:
		var m GameStarting
		c.messageToJson(&m, msg.bytes)
	case MAP_UPDATE:
		var m MapUpdate
		c.messageToJson(&m, msg.bytes)
		move := "DOWN"
		c.logger.WithFields(log.Fields{
			"move": move,
		}).Debug("Registering move to client")
		c.conn.WriteJSON(RegisterMoveMessage(move, m))
	case GAME_RESULT_EVENT:
		var m GameResult
		c.messageToJson(&m, msg.bytes)
	case SNAKE_DEAD:
		var m SnakeDead
		c.messageToJson(&m, msg.bytes)
	case INVALID_PLAYER_NAME:
		var m InvalidPlayerName
		c.messageToJson(&m, msg.bytes)
	case HEART_BEAT_RESPONSE:
		// do nothing
	case GAME_ENDED:
		var m GameEnded
		c.messageToJson(&m, msg.bytes)
		if c.isTraining {
			c.close()
		}
	case TOURNAMENT_ENDED:
		var m TournamentEnded
		c.messageToJson(&m, msg.bytes)
		c.close()
	default:
		c.logger.WithFields(log.Fields{
			"type": msg.t,
			"msg":  string(msg.bytes),
		}).Warn("Unknown message type received")
	}
}

func (c Connection) messageToJson(s interface{}, bytes []byte) {
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		c.logger.WithFields(log.Fields{
			"msg": string(bytes),
		}).Error("Unable to parse JSON message to matching struct", err)
	}

	json, _ := prettyjson.Marshal(s)
	c.logger.WithFields(log.Fields{
		"json": string(json),
		"msg":  string(bytes),
	}).Debug("Successfully parsed JSON message to struct")
}

func (c Connection) startHeartBeat(playerId string) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.heartBeat <- playerId
		case <-c.done:
			return
		}
	}
}

func (c Connection) sendHeartBeat(playerId string) {
	c.conn.WriteJSON(HeartBeatMessage(playerId))
}

func (c Connection) close() {
	defer c.conn.Close()
	c.conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	close(c.done)
}
