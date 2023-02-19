package domain

import (
	"github.com/pkg/errors"
)

type IBoard interface {
	GetBoardState() [3][3]State
	ResetBoardState()
	SetState(x, y int, s State) error
	WhoWin() State
}

type board struct {
	board [3][3]State
	rule  iRule
}

var boardOnce *board //nolint:gochecknoglobals // only allow this board

func GetBoardInst() IBoard {
	if boardOnce != nil {
		return boardOnce
	}

	b := &board{
		rule: &rule{
			lastState: 0,
		},
	}
	b.ResetBoardState()
	boardOnce = b

	return boardOnce
}

func (b *board) ResetBoardState() {
	for row := range b.board {
		for col := range b.board[row] {
			b.board[row][col].SetBlank()
		}
	}
}

func (b *board) GetBoardState() [3][3]State {
	return b.board
}

func (b *board) SetState(x, y int, s State) error {
	err := b.rule.setState(b, x, y, s)
	if err != nil {
		return errors.Wrap(err, "domain board SetState")
	}

	b.board[x][y] = s

	return nil
}

func (b *board) WhoWin() State {
	return b.rule.whoWin(b)
}
