package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_application")
var errShouldSetNicknameFirst = errors.New("should set nickname first")

type putChess struct {
	boardPort      out.IBoardPort
	loadPlayerPort out.ILoadPlayerPort
}

func NewPutChess(
	boardPort out.IBoardPort,
	loadPlayerPort out.ILoadPlayerPort,
) in.IPutChessUseCase {
	return &putChess{
		boardPort:      boardPort,
		loadPlayerPort: loadPlayerPort,
	}
}

func (pc *putChess) PutChess(cmd *in.PutChessCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	_, err := pc.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, domain.ErrGetEmpty) {
		return errShouldSetNicknameFirst
	}

	err = pc.boardPort.SetBoardState(cmd.X, cmd.Y, int(cmd.S))

	return errors.Wrap(err, errInHere.Error())
}
