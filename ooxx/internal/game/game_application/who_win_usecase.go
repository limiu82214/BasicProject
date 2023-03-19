package game_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

type whoWinUseCase struct {
	loadBoardPort game_application_port_out.ILoadBoardAdapter
}

func NewWhoWinUseCase(
	loadBoardPort game_application_port_out.ILoadBoardAdapter,
) game_application_port_in.IWhoWinUseCase {
	return &whoWinUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (w *whoWinUseCase) WhoWin() (shared.State, error) {
	board, err := w.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, game_domain.ErrGetEmpty) {
			board = game_domain.NewBoard()
		} else {
			return shared.Blank, errors.Wrap(err, "in service WhoWin")
		}
	}

	return board.WhoWin(), nil
}
