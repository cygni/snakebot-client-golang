package main

import (
	"math"
)

func TranslatePosition(position int, mapWidth int) Coordinate {
	y := math.Floor(float64(position / mapWidth))
	x := math.Abs(float64(position) - y*float64(mapWidth))
	return Coordinate{
		X: int(x),
		Y: int(y),
	}
}

func TranslatePositions(positions []int, mapWidth int) []Coordinate {
	var l []Coordinate
	for _, pos := range positions {
		l = append(l, TranslatePosition(pos, mapWidth))
	}

	return l
}

func TranslateCoordinate(coordinate Coordinate, mapWidth int) int {
	return coordinate.X + coordinate.Y*mapWidth
}

func TranslateCoordinates(coordinates []Coordinate, mapWidth int) []int {
	var l []int
	for _, coord := range coordinates {
		l = append(l, TranslateCoordinate(coord, mapWidth))
	}

	return l
}

func GetManhattanDistance(start Coordinate, goal Coordinate) int {
	x := math.Abs(float64(start.X - goal.X))
	y := math.Abs(float64(start.Y - goal.Y))

	return int(x + y)
}

func GetEuclidianDistance(start Coordinate, goal Coordinate) int {
	x := math.Pow(float64(start.X-goal.X), 2)
	y := math.Pow(float64(start.Y-goal.Y), 2)

	return int(math.Floor(math.Sqrt(x + y)))
}

func IsWithinSquare(coord Coordinate, ne_coord Coordinate, sw_coord Coordinate) bool {
	return coord.X < ne_coord.X || coord.X > sw_coord.X || coord.Y < sw_coord.Y || coord.Y > ne_coord.Y
}

type Coordinate struct {
	X int
	Y int
}

type Tile struct {
	TileType   string
	Coordinate Coordinate
}

const (
	empty     = "EMPTY"
	food      = "FOOD"
	wall      = "WALL"
	obstacle  = "OBSTACLE"
	snakeHead = "SNAKE_HEAD"
	snakeBody = "SNAKE_BODY"
	snakeTail = "SNAKE_TAIL"
)

func (t Tile) IsMovable() bool {
	switch t.TileType {
	case empty:
		return true
	case food:
		return true
	default:
		return false
	}
}

var (
	Up    = Direction{Name: "UP"}
	Down  = Direction{Name: "DOWN"}
	Left  = Direction{Name: "LEFT"}
	Right = Direction{Name: "RIGHT"}
)

type Direction struct {
	Name string
}

func (d Direction) CoordInDirection(coord Coordinate) Coordinate {
	switch d.Name {
	case Down.Name:
		return Coordinate{
			coord.X,
			coord.Y + 1,
		}
	case Up.Name:
		return Coordinate{
			coord.X,
			coord.Y - 1,
		}
	case Left.Name:
		return Coordinate{
			coord.X - 1,
			coord.Y,
		}
	case Right.Name:
		return Coordinate{
			coord.X + 1,
			coord.Y,
		}
	default:
		return coord
	}
}

func (m Map) GetSnakeById(snakeID string) *SnakeInfo {
	for _, snake := range m.SnakeInfos {
		if snake.Id == snakeID {
			return &snake
		}
	}

	return nil
}

func (m Map) GetTileAt(coordinate Coordinate) Tile {
	return Tile{
		TileType:   m.getTileType(coordinate),
		Coordinate: coordinate,
	}
}

func (m Map) getTileType(coordinate Coordinate) string {
	if m.IsCoordinateOutOfBounds(coordinate) {
		return wall
	}

	position := TranslateCoordinate(coordinate, m.Width)
	for _, snake := range m.SnakeInfos {
		for idx, pos := range snake.Positions {
			if position == pos {
				if idx == 0 {
					return snakeHead
				} else if idx == len(snake.Positions)-1 {
					return snakeTail
				} else {
					return snakeBody
				}
			}
		}
	}

	for _, obstaclePos := range m.ObstaclePositions {
		if obstaclePos == position {
			return obstacle
		}
	}

	for _, foodPos := range m.FoodPositions {
		if foodPos == position {
			return food
		}
	}

	return empty
}

func (m Map) CanSnakeMoveInDirection(snakeID string, direction Direction) bool {
	snake := m.GetSnakeById(snakeID)
	coord := TranslatePosition(snake.Positions[0], m.Width)
	coord = direction.CoordInDirection(coord)
	return m.GetTileAt(coord).IsMovable()
}

func (m Map) IsCoordinateOutOfBounds(coordinate Coordinate) bool {
	return coordinate.X < 0 || coordinate.X >= m.Width || coordinate.Y < 0 || coordinate.Y >= m.Height
}
