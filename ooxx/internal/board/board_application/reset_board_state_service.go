package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
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
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
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
