package in

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"

type IGetBoardStateUseCase interface {
	GetBoardState() ([3][3]domain.State, error)
}
