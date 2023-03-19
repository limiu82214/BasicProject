package player_adapter_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"

type IBoardAdapter interface {
	GetBoardState() ([3][3]player_domain.State, error)
	SetBoardState(x, y, s int) error
	ResetBoard() error
	WhoWin() (player_domain.State, error)
}
