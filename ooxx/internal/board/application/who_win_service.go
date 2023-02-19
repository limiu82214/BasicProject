package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type whoWin struct {
	loadBoardPort out.ILoadBoardPort
}

func NewWhoWin(loadBoardPort out.ILoadBoardPort) in.IWhoWinUseCase {
	return &whoWin{
		loadBoardPort: loadBoardPort,
	}
}

func (ww *whoWin) WhoWin() (domain.State, error) {
	board, err := ww.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, domain.ErrGetBoardEmpty) {
			board = domain.NewBoard()
		} else {
			return domain.Blank, errors.Wrap(err, "in service WhoWin")
		}
	}

	return board.WhoWin(), nil
}
