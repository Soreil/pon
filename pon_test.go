package pon

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/gotk3/gotk3/gtk"
)

const setSize = 136

func (b *board) MakeDeadWall() error {
	if len(b.deadWall) != 0 || len(b.liveWall) != setSize {
		return errors.New("Walls already in a broken state!")
	}
	b.deadWall = b.liveWall[:14]
	b.liveWall = b.liveWall[14:]
	return nil
}

func TestMakehand(t *testing.T) {
	//TODO(sjon): reuse code for other tests
	tiles := make(map[suit][]tile)
	var tileSet [4]map[suit][]tile

	tiles[man] = make([]tile, manCount)
	for i := range tiles[man] {
		tiles[man][i].suit = man
	}
	tiles[pin] = make([]tile, pinCount)
	for i := range tiles[pin] {
		tiles[pin][i].suit = pin
	}
	tiles[circle] = make([]tile, circleCount)
	for i := range tiles[circle] {
		tiles[circle][i].suit = circle
	}
	tiles[wind] = make([]tile, windCount)
	for i := range tiles[wind] {
		tiles[wind][i].suit = wind
	}
	tiles[dragon] = make([]tile, dragonCount)
	for i := range tiles[dragon] {
		tiles[dragon][i].suit = dragon
	}

	tileSet[0] = tiles
	tileSet[1] = tiles
	tileSet[2] = tiles
	tileSet[3] = tiles

	//TODO(sjon): Make seperate test
	var count int
	for suit := range tileSet {
		for value := range tileSet[suit] {
			for i := range tileSet[suit][value] {
				count++
				tileSet[suit][value][i].rank = rank(i)
			}
		}
	}

	//Make Red tiles
	//TODO(sjon): refactor in to real function
	tileSet[0][man][4].isRed = true
	tileSet[0][pin][4].isRed = true
	tileSet[0][circle][4].isRed = true

	if count != setSize {
		t.Errorf("Tilecount wrong, expected:%d, got: %d\n", setSize, count)
	}

	var board board

	for suit := range tileSet {
		for value := range tileSet[suit] {
			for _, v := range tileSet[suit][value] {
				board.liveWall = append(board.liveWall, struct {
					tile
					open bool
				}{v, false})
			}
		}
	}
	if err := board.MakeDeadWall(); err != nil {
		t.Error(err)
	}
	log.Println(board.deadWall, board.liveWall, len(board.deadWall), len(board.liveWall))

	gtk.Init(&os.Args)
	builder, err := gtk.BuilderNew()
	if err != nil {
		panic(err)
	}
	builder.AddFromFile("ui.glade")

	var window *gtk.Window
	obj, err := builder.GetObject("window1")
	if err != nil {
		panic(err)
	}
	if w, ok := obj.(*gtk.Window); ok {
		window = w
	}
	window.Show()
}
