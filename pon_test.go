package pon

import (
	"fmt"
	"testing"
)

func TestMakeHand(T *testing.T) {
	board, err := MakeBoard()
	if err != nil {
		panic(err)
	}
	players := NewPlayersFromBoard(&board)

	for i := 0; i < 13; i++ {
		err = players[0].drawTile(&board)
		if err != nil {
			panic(err)
		}
		err = players[1].drawTile(&board)
		if err != nil {
			panic(err)
		}
		err = players[2].drawTile(&board)
		if err != nil {
			panic(err)
		}
		err = players[3].drawTile(&board)
		if err != nil {
			panic(err)
		}
	}

	T.Log(fmt.Println(board))
	T.Log(fmt.Println(players))
}
