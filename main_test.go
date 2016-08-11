package main

import (
	"testing"
)

func TestBoardCreation(t *testing.T) {
	game := NewGame(
		coordinate{x: 2, y: 3},
		coordinate{x: -3, y: 2},
	)

	_, ok := game.board[coordinate{2, 3}]
	if !ok {
		t.Fatalf("Board should have had an alive cell at 2,3")
	}

	_, ok = game.board[coordinate{-3, 2}]
	if !ok {
		t.Fatalf("Board should have had an dead cell at -3,2")
	}
}

func TestDuplicateCoordinates(t *testing.T) {
	game := NewGame(
		coordinate{x: 1, y: 1},
		coordinate{x: 1, y: 1},
	)
	if len(game.board) != 1 {
		t.Fatalf("Board should only one cell")
	}
}

func TestSingleCellGame(t *testing.T) {
	game := NewGame(coordinate{x: 0, y: 0})
	game.tick()
	if len(game.board) != 0 {
		t.Fatalf("Board should have no living cells.")
	}
}

func TestOscilatingPattern(t *testing.T) {
	game := NewGame(
		coordinate{x: 0, y: 0},
		coordinate{x: 1, y: 0},
		coordinate{x: 2, y: 0},
	)

	game.tick()

	if game.isAlive(0, 1) {
		t.Errorf("Failed 0,1")
	}
	if !game.isAlive(1, 1) &&
		game.isAlive(2, 1) &&
		game.isAlive(0, 0) &&
		!game.isAlive(1, 0) &&
		game.isAlive(2, 0) &&
		game.isAlive(0, 2) &&
		!game.isAlive(1, 2) &&
		game.isAlive(2, 2) {
		t.Fatalf("Fail")
	}

}
