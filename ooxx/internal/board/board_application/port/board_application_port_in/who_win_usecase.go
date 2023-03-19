package board_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"

type IWhoWinUseCase interface {
	WhoWin() (board_domain.State, error)
}
