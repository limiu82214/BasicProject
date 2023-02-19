package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

type whoWin struct {
}

func NewWhoWin() in.IWhoWinUseCase {
	return &whoWin{}
}

func (ww *whoWin) WhoWin() domain.State {
	board := domain.GetBoardInst()

	return board.WhoWin()
}
