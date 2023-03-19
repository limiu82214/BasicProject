package out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"

type ILoadPlayerPort interface {
	GetPlayer(nickname string) (player_domain.IPlayer, error)
	SetPlayer(player_domain.IPlayer) error
}
