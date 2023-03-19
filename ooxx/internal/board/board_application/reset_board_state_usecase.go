package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type resetBoardStateUseCase struct {
	loadBoardPort board_application_port_out.ILoadBoardAdapter
}

func NewResetBoardStateUseCase(
	loadBoardPort board_application_port_out.ILoadBoardAdapter,
) board_application_port_in.IResetBoardStateUseCase {
	return &resetBoardStateUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (r *resetBoardStateUseCase) ResetBoardState() error {
	board, err := r.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return errors.Wrap(err, "in service resetBoardState")
		}
	}

	board.ResetBoardState()

	err = r.loadBoardPort.SetBoard(board)
	if err != nil {
		return errors.Wrap(err, "in service resetBoardState")
	}

	return nil
}
