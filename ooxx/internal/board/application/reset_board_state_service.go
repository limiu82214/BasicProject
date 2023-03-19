package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type resetBoardState struct {
	loadBoardPort board_application_port_out.ILoadBoardPort
}

func NewResetBoardState(loadBoardPort board_application_port_out.ILoadBoardPort) board_application_port_in.IResetBoardStateUseCase {
	return &resetBoardState{
		loadBoardPort: loadBoardPort,
	}
}

func (rbs *resetBoardState) ResetBoardState() error {
	board, err := rbs.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetEmpty) {
			board = domain.NewBoard()
		} else {
			return errors.Wrap(err, "in service resetBoardState")
		}
	}

	board.ResetBoardState()

	err = rbs.loadBoardPort.SetBoard(board)
	if err != nil {
		return errors.Wrap(err, "in service resetBoardState")
	}

	return nil
}
