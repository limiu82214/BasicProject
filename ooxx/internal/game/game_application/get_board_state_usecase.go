package game_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

type getBoardStateUseCase struct {
	loadBoardPort game_application_port_out.ILoadBoardAdapter
}

func NewGetBoardStateUseCase(
	loadBoardPort game_application_port_out.ILoadBoardAdapter,
) game_application_port_in.IGetBoardStateUseCase {
	return &getBoardStateUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (g *getBoardStateUseCase) GetBoardState() ([3][3]shared.State, error) {
	board, err := g.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, game_domain.ErrGetEmpty) {
			board = game_domain.NewBoard()
		} else {
			return [3][3]shared.State{}, errors.Wrap(err, "in service resetBoardState")
		}
	}

	return board.GetBoardState(), nil
}
