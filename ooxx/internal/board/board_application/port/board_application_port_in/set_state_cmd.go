package board_application_port_in

import (
	"github.com/pkg/errors"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
)

type SetStateCmd struct {
	isValid bool

	X int
	Y int
	S board_domain.State
}

var errOutOfRange = errors.New("must 0~2")

func NewSetStateCmd(x, y int, s board_domain.State) (*SetStateCmd, error) {
	if x < 0 || x > 2 {
		return nil, errOutOfRange
	}

	if y < 0 || y > 2 {
		return nil, errOutOfRange
	}

	if s < board_domain.Blank || s > board_domain.X {
		return nil, errOutOfRange
	}

	ssc := &SetStateCmd{
		X:       x,
		Y:       y,
		S:       s,
		isValid: true,
	}

	return ssc, nil
}

func (ssc *SetStateCmd) IsValid() bool {
	return ssc.isValid
}
