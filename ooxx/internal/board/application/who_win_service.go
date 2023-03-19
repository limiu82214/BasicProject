package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type whoWin struct {
	loadBoardPort board_application_port_out.ILoadBoardPort
}

func NewWhoWin(loadBoardPort board_application_port_out.ILoadBoardPort) board_application_port_in.IWhoWinUseCase {
	return &whoWin{
		loadBoardPort: loadBoardPort,
	}
}

func (ww *whoWin) WhoWin() (domain.State, error) {
	board, err := ww.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetEmpty) {
			board = domain.NewBoard()
		} else {
			return domain.Blank, errors.Wrap(err, "in service WhoWin")
		}
	}

	return board.WhoWin(), nil
}
