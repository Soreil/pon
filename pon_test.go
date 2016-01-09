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
	T.Log(fmt.Println(board))
}
