package board

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
	"github.com/pkg/errors"
)

type playerBoardAdapter struct {
	boardPlayerAdapter player.IBoardPlayerAdapter
}

func NewPlayerBoardAdapter(boardPlayerAdapter player.IBoardPlayerAdapter) out.IBoardPort {
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

	return newB, errors.Wrap(err, "in GetBoardState")
}
