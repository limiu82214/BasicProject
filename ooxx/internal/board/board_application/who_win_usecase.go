package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type whoWinUseCase struct {
	loadBoardPort board_application_port_out.ILoadBoardAdapter
}

func NewWhoWinUseCase(
	loadBoardPort board_application_port_out.ILoadBoardAdapter,
) board_application_port_in.IWhoWinUseCase {
	return &whoWinUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (w *whoWinUseCase) WhoWin() (board_domain.State, error) {
	board, err := w.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return board_domain.Blank, errors.Wrap(err, "in service WhoWin")
		}
	}

	return board.WhoWin(), nil
}
