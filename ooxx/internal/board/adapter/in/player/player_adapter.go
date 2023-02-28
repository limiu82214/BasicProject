package player

import (
	"log"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in board_adapter_in_player")

type IBoardPlayerAdapter interface {
	GetBoardState() ([3][3]domain.State, error)
	SetState(x, y, s int) error
	ResetBoard() error
}

type boardPlayerAdapter struct {
	getBoardUseCase        in.IGetBoardStateUseCase
	setStateUseCase        in.ISetStateUseCase
	resetBoardStateUseCase in.IResetBoardStateUseCase
}

func NewBoardPlayerAdapter(
	getBoardUseCase in.IGetBoardStateUseCase,
	setStateUseCase in.ISetStateUseCase,
	resetBoardStateUseCase in.IResetBoardStateUseCase,
) IBoardPlayerAdapter {
	return &boardPlayerAdapter{
		getBoardUseCase:        getBoardUseCase,
		setStateUseCase:        setStateUseCase,
		resetBoardStateUseCase: resetBoardStateUseCase,
	}
}

func (bpa *boardPlayerAdapter) GetBoardState() ([3][3]domain.State, error) {
	bs, err := bpa.getBoardUseCase.GetBoardState()
	return bs, errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) SetState(x, y, s int) error {
	ss := domain.State(s)

	ssc, err := in.NewSetStateCmd(x, y, ss)
	if err != nil {
		log.Println(err.Error())
	}

	err = bpa.setStateUseCase.SetState(ssc)

	return errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) ResetBoard() error {
	err := bpa.resetBoardStateUseCase.ResetBoardState()
	return errors.Wrap(err, errInHere.Error())
}
