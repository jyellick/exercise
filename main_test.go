package main

import (
	"testing"
)

func TestBoardCreation(t *testing.T) {
	game := NewGame(
		&coordinate{x: 2, y: 3},
		&coordinate{x: -3, y: 2},
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
