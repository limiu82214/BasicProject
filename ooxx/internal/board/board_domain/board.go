package board_domain

import (
	"github.com/pkg/errors"
)

type IBoard interface {
	GetBoardState() [3][3]State
	ResetBoardState()
	SetState(x, y int, s State) error
	WhoWin() State
	SetBoardStatus(bs *BoardStatus) error
	GetBoardStatus() (BoardStatus, error)
}

type Board struct {
	boardStatus *BoardStatus
	rule        iRule
}

func NewBoard() IBoard {
	b := &Board{
		boardStatus: &BoardStatus{},
		rule:        &rule{},
	}
	b.ResetBoardState()

	return b
}

func (b *Board) ResetBoardState() {
	for row := range b.boardStatus.Board {
		for col := range b.boardStatus.Board[row] {
			b.boardStatus.Board[row][col].SetBlank()
		}
	}

	b.boardStatus.LastState = Blank
}

func (b *Board) GetBoardState() [3][3]State {
	return b.boardStatus.Board
}

func (b *Board) GetBoardStatus() (BoardStatus, error) {
	return *b.boardStatus, nil
}

func (b *Board) SetBoardStatus(bs *BoardStatus) error {
	b.boardStatus = bs

	return nil
}

func (b *Board) SetState(x, y int, s State) error {
	err := b.rule.setState(b, x, y, s)
	if err != nil {
		return errors.Wrap(err, ErrSetState.Error())
	}

	b.boardStatus.Board[x][y] = s

	return nil
}

func (b *Board) WhoWin() State {
	return b.rule.whoWin(b)
}
