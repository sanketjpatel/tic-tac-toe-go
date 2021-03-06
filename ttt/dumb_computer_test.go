package ttt

import (
	"testing"

	"github.com/sdegutis/go.assert"
)

func TestDumbComputer_Move(t *testing.T) {
	var TIMES = 200
	var computer Player = NewDumbComputer()
	var board = NewBoard()

	t.Log("chooses a move at random")
	moves := make([]int, 9)
	for i := 0; i < TIMES; i++ {
		moves[computer.Move(*board)] = 1
	}
	assert.DeepEquals(t, moves, []int{1, 1, 1, 1, 1, 1, 1, 1, 1})

	t.Log("chooses only unmarked spaces")
	AddMarks(board, "X", 1, 3)
	AddMarks(board, "O", 5, 8)
	moves = make([]int, 9)
	for i := 0; i < TIMES; i++ {
		moves[computer.Move(*board)] = 1
	}
	assert.DeepEquals(t, moves, []int{1, 0, 1, 0, 1, 0, 1, 1, 0})
}

func TestDumbComputer_Mark(t *testing.T) {
	var computer = NewDumbComputer()

	t.Log("implements GetMark and SetMark")
	computer.SetMark("X")
	assert.Equals(t, computer.GetMark(), "X")
	computer.SetMark("O")
	assert.Equals(t, computer.GetMark(), "O")
}
