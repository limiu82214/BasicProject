package player_domain

type IPlayer interface {
	SetInfo(nickName string) error
	SetPlayerStore(ps *PlayerStore) error
	GetPlayerStore() (PlayerStore, error)
}
type Player struct {
	playerStore PlayerStore
}

func NewPlayer() IPlayer {
	p := &Player{
		playerStore: PlayerStore{
			NickName: "",
		},
	}

	return p
}

func (p *Player) SetInfo(nickName string) error {
	p.playerStore.NickName = nickName
	return nil
}

func (p *Player) SetPlayerStore(ps *PlayerStore) error {
	p.playerStore = *ps
	return nil
}

func (p *Player) GetPlayerStore() (PlayerStore, error) {
	return p.playerStore, nil
}
