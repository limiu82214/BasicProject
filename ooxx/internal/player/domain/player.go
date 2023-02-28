package domain

type Player struct {
	playerStore PlayerStore
}

func (p *Player) SetInfo(nickName string) {
	p.playerStore.NickName = nickName
}

func (p *Player) PutChess(x, y int, s State) {
}
