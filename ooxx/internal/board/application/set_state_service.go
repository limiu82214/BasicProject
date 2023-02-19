package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type setState struct {
}

func NewGetSiteInfo() in.ISetStateUseCase {
	return &setState{}
}

func (ss *setState) SetState(cmd *in.SetStateCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	board := domain.GetBoardInst()

	err := board.SetState(cmd.X, cmd.Y, cmd.S)
	if err != nil {
		return errors.Wrap(err, "application service setState")
	}

	return nil
}
