package game_adapter_in_player

import (
	"log"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in board_adapter_in_player")

type IBoardPlayerAdapter interface {
	GetBoardState() ([3][3]shared.State, error)
	SetState(x, y, s int) error
	ResetBoard() error
	WhoWin() (shared.State, error)
}

type boardPlayerAdapter struct {
	getBoardUseCase        game_application_port_in.IGetBoardStateUseCase
	setStateUseCase        game_application_port_in.ISetStateUseCase
	resetBoardStateUseCase game_application_port_in.IResetBoardStateUseCase
	whoWinUseCase          game_application_port_in.IWhoWinUseCase
}

func New(
	getBoardUseCase game_application_port_in.IGetBoardStateUseCase,
	setStateUseCase game_application_port_in.ISetStateUseCase,
	resetBoardStateUseCase game_application_port_in.IResetBoardStateUseCase,
	whoWinUseCase game_application_port_in.IWhoWinUseCase,
) IBoardPlayerAdapter {
	return &boardPlayerAdapter{
		getBoardUseCase:        getBoardUseCase,
		setStateUseCase:        setStateUseCase,
		resetBoardStateUseCase: resetBoardStateUseCase,
		whoWinUseCase:          whoWinUseCase,
	}
}

func (bpa *boardPlayerAdapter) GetBoardState() ([3][3]shared.State, error) {
	bs, err := bpa.getBoardUseCase.GetBoardState()
	return bs, errors.Wrap(err, errInHere.Error())
}

func (bpa *boardPlayerAdapter) SetState(x, y, s int) error {
	ss := shared.State(s)

	ssc, err := game_application_port_in.NewSetStateCmd(x, y, ss)
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

func (bpa *boardPlayerAdapter) WhoWin() (shared.State, error) {
	ds, err := bpa.whoWinUseCase.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
