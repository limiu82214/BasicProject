package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type getBoardState struct {
	loadBoardPort board_application_port_out.ILoadBoardPort
}

func NewGetBoardStateUseCase(loadBoardPort board_application_port_out.ILoadBoardPort) board_application_port_in.IGetBoardStateUseCase {
	return &getBoardState{
		loadBoardPort: loadBoardPort,
	}
}

func (gbs *getBoardState) GetBoardState() ([3][3]board_domain.State, error) {
	board, err := gbs.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return [3][3]board_domain.State{}, errors.Wrap(err, "in service resetBoardState")
		}
	}

	return board.GetBoardState(), nil
}
