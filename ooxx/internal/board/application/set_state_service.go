package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type setState struct {
	loadBoardPort board_application_port_out.ILoadBoardPort
}

func NewSetState(loadBoardPort board_application_port_out.ILoadBoardPort) board_application_port_in.ISetStateUseCase {
	return &setState{
		loadBoardPort: loadBoardPort,
	}
}

func (ss *setState) SetState(cmd *board_application_port_in.SetStateCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	board, err := ss.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetEmpty) {
			board = domain.NewBoard()
		} else {
			return errors.Wrap(err, "in service setState")
		}
	}

	err = board.SetState(cmd.X, cmd.Y, cmd.S)
	if err != nil {
		return errors.Wrap(err, "in service setState")
	}

	err = ss.loadBoardPort.SetBoard(board)
	if err != nil {
		return errors.Wrap(err, "in service setState")
	}

	whoWin := board.WhoWin()
	if whoWin != domain.Blank {
		board.ResetBoardState()

		return errors.Errorf("winner is %d", whoWin)
	}

	return nil
}
