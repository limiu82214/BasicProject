package in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"

type IWhoWinUseCase interface {
	WhoWin() (domain.State, error)
}
