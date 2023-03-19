package game_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_domain"
	"github.com/pkg/errors"
)

type resetBoardStateUseCase struct {
	loadBoardPort game_application_port_out.ILoadBoardAdapter
}

func NewResetBoardStateUseCase(
	loadBoardPort game_application_port_out.ILoadBoardAdapter,
) game_application_port_in.IResetBoardStateUseCase {
	return &resetBoardStateUseCase{
		loadBoardPort: loadBoardPort,
	}
}

func (r *resetBoardStateUseCase) ResetBoardState() error {
	board, err := r.loadBoardPort.GetBoard()
	if err != nil {
		if errors.Is(err, game_domain.ErrGetEmpty) {
			board = game_domain.NewBoard()
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
