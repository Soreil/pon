package pon

const initialPoints = 25000
const initialName = "CPU"

type players [4]player

func makePlayers() players {
	var players players
	players[0].name = initialName + "0"
	players[0].points = initialPoints
	players[1].name = initialName + "1"
	players[1].points = initialPoints
	players[2].name = initialName + "2"
	players[2].points = initialPoints
	players[3].name = initialName + "3"
	players[3].points = initialPoints
	return players
}

func (p *player) setHand(h *hand) {
	p.hand = h
}

func (p *player) setDiscard(h *hand) {
	p.discard = h
}

func (p *player) initPlayer(name string, hand *hand, discard *hand) {
	p.name = name
	p.setHand(hand)
	p.setDiscard(discard)
}
