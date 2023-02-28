package out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"

type ILoadPlayerPort interface {
	GetPlayer(nickname string) (domain.IPlayer, error)
	SetPlayer(domain.IPlayer) error
}
