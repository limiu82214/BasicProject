package in

import (
	"github.com/pkg/errors"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
)

type PutChessCmd struct {
	isValid bool

	Nickname string
	X        int
	Y        int
	S        domain.State
}

var errOutOfRange = errors.New("must 0~2")
var errNicknameErr = errors.New("nickname error")

func NewPutChessCmd(nickname string, x, y int, s domain.State) (*PutChessCmd, error) {
	l := len(nickname)
	if l < 1 || l > 3 {
		return nil, errNicknameErr
	}

	if x < 0 || x > 2 {
		return nil, errOutOfRange
	}

	if y < 0 || y > 2 {
		return nil, errOutOfRange
	}

	if s < domain.Blank || s > domain.X {
		return nil, errOutOfRange
	}

	ssc := &PutChessCmd{
		Nickname: nickname,
		X:        x,
		Y:        y,
		S:        s,
		isValid:  true,
	}

	return ssc, nil
}

func (ssc *PutChessCmd) IsValid() bool {
	return ssc.isValid
}
