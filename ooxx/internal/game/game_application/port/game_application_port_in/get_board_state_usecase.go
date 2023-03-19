package game_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/shared"

type IGetBoardStateUseCase interface {
	GetBoardState() ([3][3]shared.State, error)
}
