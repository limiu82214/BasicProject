package board_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"

type IWhoWinUseCase interface {
	WhoWin() (domain.State, error)
}
