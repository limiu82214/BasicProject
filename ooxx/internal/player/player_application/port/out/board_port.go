package out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"

type IBoardPort interface {
	GetBoardState() ([3][3]player_domain.State, error)
	SetBoardState(x, y, s int) error
	ResetBoard() error
	WhoWin() (player_domain.State, error)
}
