package user_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

type whoWinUseCase struct {
	boardPort user_adapter_port_out.IBoardAdapter
}

func NewWhoWinUseCase(boardPort user_adapter_port_out.IBoardAdapter) user_application_port_in.IWhoWinUseCase {
	return &whoWinUseCase{
		boardPort: boardPort,
	}
}

func (w *whoWinUseCase) WhoWin() (shared.State, error) {
	ds, err := w.boardPort.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
