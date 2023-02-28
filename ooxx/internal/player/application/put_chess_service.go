package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_application")

type putChess struct {
	boardPort out.IBoardPort
}

func NewPutChess(
	boardPort out.IBoardPort,
) in.IPutChessUseCase {
	return &putChess{
		boardPort: boardPort,
	}
}

func (pc *putChess) PutChess(cmd *in.PutChessCmd) error {
	err := pc.boardPort.SetBoardState(cmd.X, cmd.Y, int(cmd.S))
	return errors.Wrap(err, errInHere.Error())
}
