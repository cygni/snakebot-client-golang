package main

import (
	"testing"
)

const MAP_WIDTH = 3

func GetSnakeOne() SnakeInfo {
	return SnakeInfo{
		Name:   "1",
		Points: 0,
		TailProtectedForGameTicks: 0,
		Positions:                 []int{TranslateCoordinate(Coordinate{X: 1, Y: 1}, MAP_WIDTH)},
		Id:                        "1",
	}
}

func GetSnakeTwo() SnakeInfo {
	return SnakeInfo{
		Name:   "2",
		Points: 0,
		TailProtectedForGameTicks: 0,
		Positions:                 []int{TranslateCoordinate(Coordinate{X: 1, Y: 2}, MAP_WIDTH)},
		Id:                        "2",
	}
}

// The map used for testing, 1 and 2 represents the snakes
//yx012
//0  F
//1 11#
//2  2
func GetTestMap() Map {
	return Map{
		Width:             MAP_WIDTH,
		Height:            MAP_WIDTH,
		WorldTick:         0,
		SnakeInfos:        []SnakeInfo{GetSnakeOne(), GetSnakeTwo()},
		FoodPositions:     []int{TranslateCoordinate(Coordinate{X: 1, Y: 0}, MAP_WIDTH)},
		ObstaclePositions: []int{TranslateCoordinate(Coordinate{X: 2, Y: 1}, MAP_WIDTH)},
	}
}

func TestSnakeCanBeFoundById(t *testing.T) {
	m := GetTestMap()
	id := GetSnakeOne().Id
	found_id := m.GetSnakeById(id).Id
	if id != found_id {
		t.Error("Expected ", id, "found ", found_id)
	}
}

func TestCanNotMoveToWalls(t *testing.T) {
	m := GetTestMap()
	id := GetSnakeTwo().Id

	if m.CanSnakeMoveInDirection(id, Down) {
		t.Error("Expected snake to not be able to move to walls")
	}
}

func TestCanNotMoveToObstacles(t *testing.T) {
	m := GetTestMap()
	for _, obstaclePos := range m.ObstaclePositions {
		coord := TranslatePosition(obstaclePos, MAP_WIDTH)
		if m.GetTileAt(coord).IsMovable() {
			t.Error("Expected obstacles not to be movable to, yet could move to", coord)
		}
	}
}

func TestCanNotMoveToSnakes(t *testing.T) {
	m := GetTestMap()
	for _, snakeInfo := range m.SnakeInfos {
		for _, snakePos := range snakeInfo.Positions {
			coord := TranslatePosition(snakePos, MAP_WIDTH)
			if m.GetTileAt(coord).IsMovable() {
				t.Error("Expected snakes not to be movable to, yet could move to", coord)
			}
		}
	}
}
