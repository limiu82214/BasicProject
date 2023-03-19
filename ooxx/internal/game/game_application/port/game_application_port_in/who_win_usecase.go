package game_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/shared"

type IWhoWinUseCase interface {
	WhoWin() (shared.State, error)
}
