package main

import (
	"runtime"
)

// Inbound
var (
	GAME_ENDED          = "se.cygni.snake.api.event.GameEndedEvent"
	TOURNAMENT_ENDED    = "se.cygni.snake.api.event.TournamentEndedEvent"
	MAP_UPDATE          = "se.cygni.snake.api.event.MapUpdateEvent"
	SNAKE_DEAD          = "se.cygni.snake.api.event.SnakeDeadEvent"
	GAME_STARTING       = "se.cygni.snake.api.event.GameStartingEvent"
	PLAYER_REGISTERED   = "se.cygni.snake.api.response.PlayerRegistered"
	INVALID_PLAYER_NAME = "se.cygni.snake.api.exception.InvalidPlayerName"
	HEART_BEAT_RESPONSE = "se.cygni.snake.api.response.HeartBeatResponse"
	GAME_LINK_EVENT     = "se.cygni.snake.api.event.GameLinkEvent"
	GAME_RESULT_EVENT   = "se.cygni.snake.api.event.GameResultEvent"
)

// Outbound
var (
	REGISTER_PLAYER_MESSAGE_TYPE = "se.cygni.snake.api.request.RegisterPlayer"
	START_GAME                   = "se.cygni.snake.api.request.StartGame"
	REGISTER_MOVE                = "se.cygni.snake.api.request.RegisterMove"
	HEART_BEAT_REQUEST           = "se.cygni.snake.api.request.HeartBeatRequest"
	CLIENT_INFO                  = "se.cygni.snake.api.request.ClientInfo"
)

func RegisterPlayerMessage(playerName string) RegisterPlayer {
	return RegisterPlayer{
		Type:       "se.cygni.snake.api.request.RegisterPlayer",
		PlayerName: playerName,
	}
}

func StartGameMessage() StartGame {
	return StartGame{
		Type: "se.cygni.snake.api.request.StartGame",
	}
}

func RegisterMoveMessage(direction string, msg MapUpdate) RegisterMove {
	return RegisterMove{
		Type:              "se.cygni.snake.api.request.RegisterMove",
		Direction:         direction,
		GameTick:          msg.GameTick,
		ReceivingPlayerId: msg.ReceivingPlayerId,
		GameId:            msg.GameId,
	}
}

func HeartBeatMessage(receivingPlayerId string) HeartBeat {
	return HeartBeat{
		Type:              "se.cygni.snake.api.request.HeartBeatRequest",
		ReceivingPlayerId: receivingPlayerId,
	}
}

func ClientInfoMessage() ClientInfo {
	os := runtime.GOOS
	if os == "darwin" {
		os = "macOS"
	}

	return ClientInfo{
		Type:                   "se.cygni.snake.api.request.ClientInfo",
		Language:               "go",
		LanguageVersion:        runtime.Version(),
		OperatingSystem:        os,
		OperatingSystemVersion: "???", // seems impossible to do reliably
		ClientVersion:          "0.1",
	}
}
