package main

import (
	log "github.com/sirupsen/logrus"
	"time"
	"math/rand"
)

type Snake struct {
	Name         string
	GameSettings GameSettings
	ID           string
}

func (snake Snake) GetNextMove(m Map) string {

	directions := []Direction{Up, Down, Left, Right}
	var possibleDirections []Direction

	for _, direction := range directions {
		if m.CanSnakeMoveInDirection(snake.ID, direction) {
			possibleDirections = append(possibleDirections, direction)
		}
	}

	nbrOfDirections := len(possibleDirections)
	if nbrOfDirections <= 0 {
		return Down.Name
	}
	return possibleDirections[rand.Intn(nbrOfDirections)].Name
}

func (snake *Snake) OnPlayerRegistered(s GameSettings, snakeID string) {
	snake.GameSettings = s
	snake.ID = snakeID
	log.Debug("Player registered.")
}

func (snake Snake) OnSnakeDead(reason string) {
	log.WithField("reason", reason).Debug("The snake died")
}

func (snake Snake) OnGameStarting() {
	rand.Seed(time.Now().Unix())
	log.Debug("All snakes are ready to rock. Game is starting.")
}

func (snake Snake) OnInvalidPlayername(reasonCode int) {
	log.Debug("Player name is invalid.")
}

func GetSnake() Snake {
	return Snake{
		Name: "golang",
		ID:   "",
	}
}
