package player

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type IBoardPlayerAdapter interface {
	GetBoardState() ([3][3]domain.State, error)
}

type boardPlayerAdapter struct {
	getBoardUseCase in.IGetBoardStateUseCase
}

func NewBoardPlayerAdapter(getBoardUseCase in.IGetBoardStateUseCase) IBoardPlayerAdapter {
	return &boardPlayerAdapter{
		getBoardUseCase: getBoardUseCase,
	}
}

func (bpa *boardPlayerAdapter) GetBoardState() ([3][3]domain.State, error) {
	bs, err := bpa.getBoardUseCase.GetBoardState()
	return bs, errors.Wrap(err, "in player_adapter GetBoardState")
}
