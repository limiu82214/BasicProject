package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

type getBoardState struct {
	boardPort out.IBoardPort
}

func NewGetBoardState(boardPort out.IBoardPort) in.IGetBoardStateUseCase {
	return &getBoardState{
		boardPort: boardPort,
	}
}

func (gbs *getBoardState) GetBoardState() ([3][3]player_domain.State, error) {
	tmpB, err := gbs.boardPort.GetBoardState()
	return tmpB, errors.Wrap(err, "in GetBoardState")
}
