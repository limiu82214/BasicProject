package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/pkg/errors"
)

type resetBoard struct {
	boardPort player_adapter_port_out.IBoardPort
}

func NewResetBoard(boardPort player_adapter_port_out.IBoardPort) player_adapter_port_in.IResetBoardUseCase {
	return &resetBoard{
		boardPort: boardPort,
	}
}

func (rb *resetBoard) ResetBoard() error {
	err := rb.boardPort.ResetBoard()
	return errors.Wrap(err, errInHere.Error())
}
