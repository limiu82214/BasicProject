package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type whoWin struct {
	loadBoardPort board_application_port_out.ILoadBoardPort
}

func NewWhoWinUseCase(loadBoardPort board_application_port_out.ILoadBoardPort) board_application_port_in.IWhoWinUseCase {
	return &whoWin{
		loadBoardPort: loadBoardPort,
	}
}

func (ww *whoWin) WhoWin() (board_domain.State, error) {
	board, err := ww.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return board_domain.Blank, errors.Wrap(err, "in service WhoWin")
		}
	}

	return board.WhoWin(), nil
}
