package board_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"

type IGetBoardStateUseCase interface {
	GetBoardState() ([3][3]board_domain.State, error)
}
