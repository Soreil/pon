package pon

const initialPoints = 25000
const initialName = "CPU"

type player struct {
	name    string
	hand    *hand
	discard *hand
	points  int
}

type players [4]player

func NewPlayersFromBoard(b *board) players {
	var p players
	for i := range p {
		p[i].hand = &b.playerHands[i]
		p[i].discard = &b.playerDiscards[i]
		p[i].name = initialName + string('0'+i)
		p[i].points = initialPoints
	}
	return p
}

func (p *player) drawTile(b *board) error {
	t, err := b.liveWall.draw()
	if err != nil {
		return err
	}
	p.hand.add(t)
	return nil
}
