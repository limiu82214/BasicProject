package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
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

func (gbs *getBoardState) GetBoardState() ([3][3]domain.State, error) {
	tmpB, err := gbs.boardPort.GetBoardState()
	// 不在這裡轉換，那是adapter該做的事情
	// if err != nil {
	// 	return [3][3]domain.State{}, errors.Wrap(err, "in service resetBoardState")
	// }

	// test

	// for i := range tmpB {
	// 	for j := range tmpB[i] {
	// 		b[i][j] = domain.State(tmpB[i][j])
	// 	}
	// }

	return tmpB, errors.Wrap(err, "in GetBoardState")
}
