package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

type getBoardStateUseCase struct {
	boardPort player_adapter_port_out.IBoardAdapter
}

func NewGetBoardStateUseCase(boardPort player_adapter_port_out.IBoardAdapter) player_application_port_in.IGetBoardStateUseCase {
	return &getBoardStateUseCase{
		boardPort: boardPort,
	}
}

func (g *getBoardStateUseCase) GetBoardState() ([3][3]player_domain.State, error) {
	tmpB, err := g.boardPort.GetBoardState()
	return tmpB, errors.Wrap(err, "in GetBoardState")
}
