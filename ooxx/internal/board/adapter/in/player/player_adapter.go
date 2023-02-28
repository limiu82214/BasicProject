package player

import (
	"log"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type IBoardPlayerAdapter interface {
	GetBoardState() ([3][3]domain.State, error)
	SetState(x, y, s int) error
}

type boardPlayerAdapter struct {
	getBoardUseCase in.IGetBoardStateUseCase
	setStateUseCase in.ISetStateUseCase
}

func NewBoardPlayerAdapter(
	getBoardUseCase in.IGetBoardStateUseCase,
	setStateUseCase in.ISetStateUseCase,
) IBoardPlayerAdapter {
	return &boardPlayerAdapter{
		getBoardUseCase: getBoardUseCase,
		setStateUseCase: setStateUseCase,
	}
}

func (bpa *boardPlayerAdapter) GetBoardState() ([3][3]domain.State, error) {
	bs, err := bpa.getBoardUseCase.GetBoardState()
	return bs, errors.Wrap(err, "in board_adapter_in_player GetBoardState")
}

func (bpa *boardPlayerAdapter) SetState(x, y, s int) error {
	ss := domain.State(s)

	ssc, err := in.NewSetStateCmd(x, y, ss)
	if err != nil {
		log.Println(err.Error())
	}

	err = bpa.setStateUseCase.SetState(ssc)

	return errors.Wrap(err, "in board_adapter_in_player SetState")
}
