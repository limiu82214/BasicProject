package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/out"
	"github.com/pkg/errors"
)

type resetBoard struct {
	boardPort out.IBoardPort
}

func NewResetBoard(boardPort out.IBoardPort) in.IResetBoardUseCase {
	return &resetBoard{
		boardPort: boardPort,
	}
}

func (rb *resetBoard) ResetBoard() error {
	err := rb.boardPort.ResetBoard()
	return errors.Wrap(err, errInHere.Error())
}
