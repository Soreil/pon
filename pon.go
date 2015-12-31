package pon

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
	east rank = iota
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

type player struct {
	name    string
	hand    *hand
	discard *hand
	points  int
}
