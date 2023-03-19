package player_adapter_out_board

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/in/board_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_adapter_out_board")

type playerBoardAdapter struct {
	boardPlayerAdapter board_adapter_in_player.IBoardPlayerAdapter
}

func New(boardPlayerAdapter board_adapter_in_player.IBoardPlayerAdapter) player_adapter_port_out.IBoardPort {
	return &playerBoardAdapter{
		boardPlayerAdapter: boardPlayerAdapter,
	}
}

func (pba *playerBoardAdapter) GetBoardState() ([3][3]player_domain.State, error) {
	b, err := pba.boardPlayerAdapter.GetBoardState()
	newB := [3][3]player_domain.State{}

	for i := range b {
		for j := range b[i] {
			newB[i][j] = player_domain.State(b[i][j])
		}
	}

	return newB, errors.Wrap(err, errInHere.Error())
}

func (pba *playerBoardAdapter) SetBoardState(x, y, s int) error {
	err := pba.boardPlayerAdapter.SetState(x, y, s)
	return errors.Wrap(err, errInHere.Error())
}

func (pba *playerBoardAdapter) ResetBoard() error {
	err := pba.boardPlayerAdapter.ResetBoard()
	return errors.Wrap(err, errInHere.Error())
}

func (pba *playerBoardAdapter) WhoWin() (player_domain.State, error) {
	ds, err := pba.boardPlayerAdapter.WhoWin()
	return player_domain.State(ds), errors.Wrap(err, errInHere.Error())
}