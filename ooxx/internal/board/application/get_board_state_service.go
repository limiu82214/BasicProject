package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

type getBoardState struct {
}

func NewGetBoardState() in.IGetBoardStateUseCase {
	return &getBoardState{}
}

func (gbs *getBoardState) GetBoardState() [3][3]domain.State {
	board := domain.GetBoardInst()

	return board.GetBoardState()
}
