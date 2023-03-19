package user_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/pkg/errors"
)

type resetBoardUseCase struct {
	boardPort user_adapter_port_out.IBoardAdapter
}

func NewResetBoardUseCase(
	boardPort user_adapter_port_out.IBoardAdapter,
) user_application_port_in.IResetBoardUseCase {
	return &resetBoardUseCase{
		boardPort: boardPort,
	}
}

func (r *resetBoardUseCase) ResetBoard() error {
	err := r.boardPort.ResetBoard()
	return errors.Wrap(err, errInHere.Error())
}
