package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

type resetBoardState struct {
}

func NewResetBoardState() in.IResetBoardStateUseCase {
	return &resetBoardState{}
}

func (rbs *resetBoardState) ResetBoardState() {
	board := domain.GetBoard()
	board.ResetBoardState()
}
