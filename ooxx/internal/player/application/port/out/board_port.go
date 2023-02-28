package out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"

type IBoardPort interface {
	GetBoardState() ([3][3]domain.State, error)
	SetBoardState(x, y, s int) error
	ResetBoard() error
}
