package board_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/pkg/errors"
)

type getBoardStateUseCase struct {
	loadBoardPort board_application_port_out.ILoadBoardAdapter
}

func NewGetBoardStateUseCase(
	loadBoardPort board_application_port_out.ILoadBoardAdapter,
) board_application_port_in.IGetBoardStateUseCase {
	return &getBoardStateUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (g *getBoardStateUseCase) GetBoardState() ([3][3]board_domain.State, error) {
	board, err := g.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, board_domain.ErrGetEmpty) {
			board = board_domain.NewBoard()
		} else {
			return [3][3]board_domain.State{}, errors.Wrap(err, "in service resetBoardState")
		}
	}

	return board.GetBoardState(), nil
}
