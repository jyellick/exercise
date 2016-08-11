package main

import (
	"fmt"
	"time"
)

type coordinate struct {
	x int
	y int
}

type game struct {
	board map[coordinate]struct{}
}

func NewGame(coords ...coordinate) *game {
	game := &game{board: make(map[coordinate]struct{})}
	for _, coord := range coords {
		game.board[coord] = struct{}{}
	}
	return game
}

type cellUpdate struct {
	aliveNeighbors int
	wasAlive       bool
}

func (cu *cellUpdate) nowAlive() bool {
	if cu.wasAlive {
		if cu.aliveNeighbors < 2 {
			return false
		}
		if cu.aliveNeighbors > 3 {
			return false
		}

		return true
	}

	if cu.aliveNeighbors == 3 {
		return true
	}

	return false
}

func (c coordinate) neighbors() []coordinate {
	neighbors := make([]coordinate, 8)
	neighbors[0] = coordinate{x: c.x + 0, y: c.y + 1}
	neighbors[1] = coordinate{x: c.x + 1, y: c.y + 1}
	neighbors[2] = coordinate{x: c.x + 1, y: c.y + 0}
	neighbors[3] = coordinate{x: c.x + 1, y: c.y - 1}
	neighbors[4] = coordinate{x: c.x + 0, y: c.y - 1}
	neighbors[5] = coordinate{x: c.x - 1, y: c.y - 1}
	neighbors[6] = coordinate{x: c.x - 1, y: c.y + 0}
	neighbors[7] = coordinate{x: c.x - 1, y: c.y + 1}
	return neighbors
}

func (g *game) isAlive(x, y int) bool {
	_, ok := g.board[coordinate{x: x, y: y}]
	return ok
}

func (g *game) tick() {
	status := make(map[coordinate]*cellUpdate)
	for coord := range g.board {
		if _, ok := status[coord]; !ok {
			status[coord] = &cellUpdate{wasAlive: true}
		}
		for _, neighbor := range coord.neighbors() {
			cu, ok := status[neighbor]
			if !ok {
				cu = &cellUpdate{}
				status[neighbor] = cu
				_, ok = g.board[neighbor]
				cu.wasAlive = ok
			}
			cu.aliveNeighbors++
		}
	}

	g.board = make(map[coordinate]struct{})
	for coord, cu := range status {
		if cu.nowAlive() {
			g.board[coord] = struct{}{}
		}
	}
}

func (g *game) display() {
	var maxX, minX, maxY, minY int
	for coord := range g.board {
		if coord.x < minX {
			minX = coord.x
		}
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.y < minY {
			minY = coord.y
		}
		if coord.y > maxY {
			maxY = coord.y
		}
	}

	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			if g.isAlive(x, y) {
				fmt.Printf("X")
			} else {
				fmt.Printf("")
			}
		}
		fmt.Println("")
	}
}

func main() {
	game := NewGame(
		coordinate{x: 0, y: 0},
		coordinate{x: 1, y: 0},
		coordinate{x: 2, y: 0},
	)
	for i := 0; i < 10; i++ {
		game.display()
		time.Sleep(time.Second)
		game.tick()
		fmt.Println("==================")
	}
}
