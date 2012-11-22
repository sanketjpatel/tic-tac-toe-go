package tictactoe

import "github.com/sdegutis/go.assert"
import "testing"

func TestGameOver( t *testing.T ) {
  var game Game = NewGame()

  t.Log( "New Game is not over" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with two-in-a-row is not over" )
  addMarks( game.Board, []int{ 4, 8 }, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"X\" is over" )
  game.Board.Mark( 0, "X" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row \"O\" is over" )
  game.Board.Reset()
  addMarks( game.Board, []int{ 2, 4, 6 }, "O" )
  assert.True( t, game.IsOver() )

  t.Log( "Game with three-in-a-row mismatched is not over" )
  game.Board.Mark( 2, "X" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with nearly-full, non-winning board is not over" )
  game.Board.Reset()
  addMarks( game.Board, []int{ 0, 1, 4, 5 }, "X" )
  addMarks( game.Board, []int{ 2, 3, 7, 8 }, "O" )
  assert.False( t, game.IsOver() )

  t.Log( "Game with full, non-winning board is over" )
  game.Board.Mark( 6, "X" )
  assert.True( t, game.IsOver() )
}

func addMarks( b Board, set []int, mark string ) {
  for _,p := range set {
    b.Mark( p, mark )
  }
}
