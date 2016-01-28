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
	players := makePlayers()
	players[0].initPlayer("One", &board.playerHands[0], &board.playerDiscards[0])
	players[1].initPlayer("Two", &board.playerHands[1], &board.playerDiscards[1])
	players[2].initPlayer("Three", &board.playerHands[2], &board.playerDiscards[2])
	players[3].initPlayer("Four", &board.playerHands[3], &board.playerDiscards[3])
	T.Log(fmt.Println(board))
	T.Log(fmt.Println(players))
}
