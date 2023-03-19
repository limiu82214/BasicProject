package player_adapter_out_board

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_adapter_out_board")

type playerBoardAdapter struct {
	boardPlayerAdapter game_adapter_in_player.IBoardPlayerAdapter
}

func New(boardPlayerAdapter game_adapter_in_player.IBoardPlayerAdapter) player_adapter_port_out.IBoardAdapter {
	return &playerBoardAdapter{
		boardPlayerAdapter: boardPlayerAdapter,
	}
}

func (pba *playerBoardAdapter) GetBoardState() ([3][3]shared.State, error) {
	b, err := pba.boardPlayerAdapter.GetBoardState()
	newB := [3][3]shared.State{}

	for i := range b {
		for j := range b[i] {
			newB[i][j] = shared.State(b[i][j])
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

func (pba *playerBoardAdapter) WhoWin() (shared.State, error) {
	ds, err := pba.boardPlayerAdapter.WhoWin()
	return shared.State(ds), errors.Wrap(err, errInHere.Error())
}
