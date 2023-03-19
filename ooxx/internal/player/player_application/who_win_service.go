package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

type whoWinUseCase struct {
	boardPort player_adapter_port_out.IBoardAdapter
}

func NewWhoWinUseCase(boardPort player_adapter_port_out.IBoardAdapter) player_application_port_in.IWhoWinUseCase {
	return &whoWinUseCase{
		boardPort: boardPort,
	}
}

func (w *whoWinUseCase) WhoWin() (player_domain.State, error) {
	ds, err := w.boardPort.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
