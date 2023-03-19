package user_adapter_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_domain"

type ILoadPlayerAdapter interface {
	GetPlayer(nickname string) (user_domain.IPlayer, error)
	SetPlayer(user_domain.IPlayer) error
}
