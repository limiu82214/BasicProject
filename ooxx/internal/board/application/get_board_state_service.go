package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type getBoardState struct {
	loadBoardPort out.ILoadBoardPort
}

func NewGetBoardState(loadBoardPort out.ILoadBoardPort) board_application_port_in.IGetBoardStateUseCase {
	return &getBoardState{
		loadBoardPort: loadBoardPort,
	}
}

func (gbs *getBoardState) GetBoardState() ([3][3]domain.State, error) {
	board, err := gbs.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetEmpty) {
			board = domain.NewBoard()
		} else {
			return [3][3]domain.State{}, errors.Wrap(err, "in service resetBoardState")
		}
	}

	return board.GetBoardState(), nil
}
