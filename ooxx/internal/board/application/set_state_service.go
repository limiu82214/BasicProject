package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

type setState struct {
}

func NewGetSiteInfo() in.ISetStateUseCase {
	return &setState{}
}

func (ss *setState) SetState(cmd *in.SetStateCmd) {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	board := domain.GetBoard()
	board.SetState(cmd.X, cmd.Y, cmd.S)
}
