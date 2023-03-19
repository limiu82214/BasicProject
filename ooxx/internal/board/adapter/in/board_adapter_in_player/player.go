package board_adapter_in_player

import (
	"log"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in board_adapter_in_player")

type IBoardPlayerAdapter interface {
	GetBoardState() ([3][3]board_domain.State, error)
	SetState(x, y, s int) error
	ResetBoard() error
	WhoWin() (board_domain.State, error)
}

type boardPlayerAdapter struct {
	getBoardUseCase        board_application_port_in.IGetBoardStateUseCase
	setStateUseCase        board_application_port_in.ISetStateUseCase
	resetBoardStateUseCase board_application_port_in.IResetBoardStateUseCase
	whoWinUseCase          board_application_port_in.IWhoWinUseCase
}

func New(
	getBoardUseCase board_application_port_in.IGetBoardStateUseCase,
	setStateUseCase board_application_port_in.ISetStateUseCase,
	resetBoardStateUseCase board_application_port_in.IResetBoardStateUseCase,
	whoWinUseCase board_application_port_in.IWhoWinUseCase,
) IBoardPlayerAdapter {
	return &boardPlayerAdapter{
		getBoardUseCase:        getBoardUseCase,
		setStateUseCase:        setStateUseCase,
		resetBoardStateUseCase: resetBoardStateUseCase,
		whoWinUseCase:          whoWinUseCase,
	}
}

func (bpa *boardPlayerAdapter) GetBoardState() ([3][3]board_domain.State, error) {
	bs, err := bpa.getBoardUseCase.GetBoardState()
	return bs, errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) SetState(x, y, s int) error {
	ss := board_domain.State(s)

	ssc, err := board_application_port_in.NewSetStateCmd(x, y, ss)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
	}

	err = bpa.setStateUseCase.SetState(ssc)

	return errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) ResetBoard() error {
	err := bpa.resetBoardStateUseCase.ResetBoardState()
	return errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) WhoWin() (board_domain.State, error) {
	ds, err := bpa.whoWinUseCase.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
