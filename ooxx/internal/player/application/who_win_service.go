package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
	"github.com/pkg/errors"
)

type whoWin struct {
	boardPort out.IBoardPort
}

func NewWhoWin(boardPort out.IBoardPort) in.IWhoWinUseCase {
	return &whoWin{
		boardPort: boardPort,
	}
}

func (rb *whoWin) WhoWin() (domain.State, error) {
	ds, err := rb.boardPort.WhoWin()
	return ds, errors.Wrap(err, errInHere.Error())
}
