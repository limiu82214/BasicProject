package user_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

type getBoardStateUseCase struct {
	boardPort user_adapter_port_out.IBoardAdapter
}

func NewGetBoardStateUseCase(
	boardPort user_adapter_port_out.IBoardAdapter,
) user_application_port_in.IGetBoardStateUseCase {
	return &getBoardStateUseCase{
		boardPort: boardPort,
	}
}

func (g *getBoardStateUseCase) GetBoardState() ([3][3]shared.State, error) {
	tmpB, err := g.boardPort.GetBoardState()
	return tmpB, errors.Wrap(err, "in GetBoardState")
}
