package pon

import (
	"errors"
	"fmt"
	"strings"
)

const tileSetSize = 136

const (
	manCount    = 9
	pinCount    = 9
	circleCount = 9
	windCount   = 4
	dragonCount = 3
)

type suit int

const (
	man suit = iota
	pin
	circle
	wind
	dragon
)

func (s suit) String() string {
	switch s {
	case man:
		return "man"
	case pin:
		return "pin"
	case circle:
		return "circle"
	case wind:
		return "wind"
	case dragon:
		return "dragon"
	default:
		return string(s)
	}
}

type rank int

const (
	east rank = 10 + iota
	south
	west
	north

	white
	red
	green
)

func (r rank) String() string {
	switch r {
	case east:
		return "east"
	case south:
		return "south"
	case west:
		return "west"
	case north:
		return "north"
	case white:
		return "white"
	case red:
		return "red"
	case green:
		return "green"
	default:
		return string(r + '0')
	}
}

type tile struct {
	suit
	rank
	isRed bool
}

func (t tile) String() string {
	//First unicode character,they are stacked back to back
	base := '\U0001F000'
	var char rune
	switch t.suit {
	case man:
		char = base + windCount + dragonCount
	case pin:
		char = base + windCount + dragonCount + manCount
	case circle:
		char = base + windCount + dragonCount + manCount + pinCount
	case wind:
		char = base
	case dragon:
		char = base
	default:
		char = base
	}
	return string(char + rune(t.rank))
}

type hand []struct {
	tile
	open bool
}

type board struct {
	playerHands    [4]hand
	playerDiscards [4]hand
	liveWall       hand
	deadWall       hand
}

func (b board) String() string {
	var out []string
	out = append(out, "Hands:")
	out = append(out, fmt.Sprint(b.playerHands[0]))
	out = append(out, fmt.Sprint(b.playerHands[1]))
	out = append(out, fmt.Sprint(b.playerHands[2]))
	out = append(out, fmt.Sprint(b.playerHands[3]))
	out = append(out, "Discards:")
	out = append(out, fmt.Sprint(b.playerDiscards[0]))
	out = append(out, fmt.Sprint(b.playerDiscards[1]))
	out = append(out, fmt.Sprint(b.playerDiscards[2]))
	out = append(out, fmt.Sprint(b.playerDiscards[3]))
	out = append(out, "Live wall:")
	out = append(out, fmt.Sprint(b.liveWall))
	out = append(out, "Dead wall:")
	out = append(out, fmt.Sprint(b.deadWall))
	return strings.Join(out, "\n")
}

type player struct {
	name    string
	hand    *hand
	discard *hand
	points  int
}

func (b *board) MakeDeadWall() error {
	if len(b.deadWall) != 0 || len(b.liveWall) != tileSetSize {
		return errors.New("Walls already in a broken state!")
	}
	b.deadWall = b.liveWall[:14]
	b.liveWall = b.liveWall[14:]
	return nil
}

//Creates a new board
func MakeBoard() (board, error) {
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

	var count int
	for suit := range tileSet {
		for value := range tileSet[suit] {
			for i := range tileSet[suit][value] {
				count++
				tileSet[suit][value][i].rank = rank(i)
			}
		}
	}

	if count != tileSetSize {
		return board{}, errors.New(fmt.Sprintf("Tilecount wrong, expected:%d, got: %d\n", tileSetSize, count))
	}

	//Make Red tiles
	tileSet[0][man][4].isRed = true
	tileSet[0][pin][4].isRed = true
	tileSet[0][circle][4].isRed = true

	var b board

	for suit := range tileSet {
		for value := range tileSet[suit] {
			for _, v := range tileSet[suit][value] {
				b.liveWall = append(b.liveWall, struct {
					tile
					open bool
				}{v, false})
			}
		}
	}
	if err := b.MakeDeadWall(); err != nil {
		return board{}, err
	}
	return b, nil
}
