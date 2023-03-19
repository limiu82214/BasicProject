package player_application_port_in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"

type IGetBoardStateUseCase interface {
	GetBoardState() ([3][3]player_domain.State, error)
}
