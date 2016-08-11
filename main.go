package main

import (
	"fmt"
)

type coordinate struct {
	x int
	y int
}

type gameTile struct {
	alive bool
}

type updatingTile struct {
	wasAlive       bool
	aliveNeighbors int
	deadNeighbors  int
}

func (ut *updatingTile) gameTile() *gameTile {
	if ut.wasAlive {
		if ut.aliveNeighbors < 2 {
			return &gameTile{false}
		}
		if ut.aliveNeighbors > 3 {
			return &gameTile{false}
		}
		return &gameTile{true}
	}

	if ut.aliveNeighbors == 3 {
		return &gameTile{true}
	}

	return nil

}

type tilePair struct {
	coord *coordinate
	tile  *gameTile
}

type game struct {
	board map[coordinate]struct{}
}

func NewGame(tiles ...coordinate) *game {
	game := &game{board: make(map[coordinate]struct{})}
	for _, coord := range tiles {
		game.board[coord] = struct{}{}
	}
	return game
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

func (g *game) tick() {
	nextGen := make(map[coordinate]*updatingTile)
	for key, val := range g.board {
		neighbors := key.neighbors()
		for _, neighbor := range neighbors {
			ut, ok := nextGen[neighbor]
			if !ok {
				ut = &updatingTile{}
				nextGen[neighbor] = ut
				ival, ok := g.board[neighbor]
				if ok && ival.alive {
					ut.wasAlive = true
				}
			}
			if val.alive {
				ut.aliveNeighbors++
			} else {
				ut.deadNeighbors++
			}
		}
	}

	for key, val := range nextGen {
		// Temp
	}
}

func main() {
	fmt.Println("Testing")
}
