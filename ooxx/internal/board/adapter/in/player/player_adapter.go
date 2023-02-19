package player

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/pkg/errors"
)

type IBoardPlayerAdapter interface {
	SetState(x, y int, s domain.State) error
	WhoWin() (domain.State, error)
}

type boardPlayerAdapter struct {
	setStateUseCase in.ISetStateUseCase
	whoWinUseCase   in.IWhoWinUseCase
}

func NewBoardPlayerAdapter() IBoardPlayerAdapter {
	return &boardPlayerAdapter{}
}

func (bpa *boardPlayerAdapter) SetState(x, y int, s domain.State) error {
	ssc, err := in.NewSetStateCmd(x, y, s)
	if err != nil {
		return errors.Wrap(err, "SetState")
	}

	err = bpa.setStateUseCase.SetState(ssc)
	if err != nil {
		return errors.Wrap(err, "player_adapter setState")
	}

	return nil
}

func (bpa *boardPlayerAdapter) WhoWin() (domain.State, error) {
	return bpa.whoWinUseCase.WhoWin()
}
