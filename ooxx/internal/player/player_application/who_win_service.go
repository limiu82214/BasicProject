package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

type whoWin struct {
	boardPort player_adapter_port_out.IBoardPort
}

func NewWhoWinUseCase(boardPort player_adapter_port_out.IBoardPort) player_application_port_in.IWhoWinUseCase {
	return &whoWin{
		boardPort: boardPort,
	}
}

func (rb *whoWin) WhoWin() (player_domain.State, error) {
	ds, err := rb.boardPort.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
