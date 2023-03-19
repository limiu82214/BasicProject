package player_adapter_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"

type IWhoWinUseCase interface {
	WhoWin() (player_domain.State, error)
}
