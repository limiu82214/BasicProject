package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type setState struct {
	loadBoardPort out.ILoadBoardPort
}

func NewSetState(loadBoardPort out.ILoadBoardPort) in.ISetStateUseCase {
	return &setState{
		loadBoardPort: loadBoardPort,
	}
}

func (ss *setState) SetState(cmd *in.SetStateCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	board, err := ss.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetBoardEmpty) {
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
