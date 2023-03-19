package player_adapter_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/shared"

type IBoardAdapter interface {
	GetBoardState() ([3][3]shared.State, error)
	SetBoardState(x, y, s int) error
	ResetBoard() error
	WhoWin() (shared.State, error)
}
