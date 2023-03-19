package board

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/board_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_adapter_out_board")

type playerBoardAdapter struct {
	boardPlayerAdapter board_adapter_in_player.IBoardPlayerAdapter
}

func NewPlayerBoardAdapter(boardPlayerAdapter board_adapter_in_player.IBoardPlayerAdapter) out.IBoardPort {
	return &playerBoardAdapter{
		boardPlayerAdapter: boardPlayerAdapter,
	}
}

func (pba *playerBoardAdapter) GetBoardState() ([3][3]domain.State, error) {
	b, err := pba.boardPlayerAdapter.GetBoardState()
	newB := [3][3]domain.State{}

	for i := range b {
		for j := range b[i] {
			newB[i][j] = domain.State(b[i][j])
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

func (pba *playerBoardAdapter) WhoWin() (domain.State, error) {
	ds, err := pba.boardPlayerAdapter.WhoWin()
	return domain.State(ds), errors.Wrap(err, errInHere.Error())
}
