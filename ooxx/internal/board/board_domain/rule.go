package board_domain

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/pkg/errors"
)

type iRule interface {
	whoWin(b *Board) shared.State
	isMyTune(b *Board, s shared.State) bool
	setState(b *Board, x, y int, s shared.State) error
}

type rule struct {
}

var errNotYourTurn = errors.New("is not your turn")
var errIsNotBlankPos = errors.New("this pos is not Blank")

func (r *rule) setState(b *Board, x, y int, s shared.State) error {
	if !r.isMyTune(b, s) {
		return errNotYourTurn
	}

	if !r.isBlankPos(b, x, y) {
		return errIsNotBlankPos
	}

	b.boardStatus.LastState = s

	return nil
}

func (r *rule) isBlankPos(b *Board, x, y int) bool {
	return b.boardStatus.Board[x][y] == shared.Blank
}

func (r *rule) isMyTune(b *Board, s shared.State) bool {
	lastPlayState := shared.State(b.boardStatus.LastState)
	if lastPlayState == shared.Blank {
		return true
	}

	return lastPlayState != s
}

func (r *rule) whoWin(b *Board) shared.State {
	if r.isWin(b, shared.O) {
		return shared.O
	}

	if r.isWin(b, shared.X) {
		return shared.X
	}

	return shared.Blank
}

func (r *rule) isWin(b *Board, s shared.State) bool { //nolint:cyclop // this is simplest
	// row, col
	for i := 0; i < 2; i++ {
		if b.boardStatus.Board[i][0] == s &&
			b.boardStatus.Board[i][1] == s &&
			b.boardStatus.Board[i][2] == s {
			return true
		}

		if b.boardStatus.Board[0][i] == s &&
			b.boardStatus.Board[1][i] == s &&
			b.boardStatus.Board[2][i] == s {
			return true
		}
	}

	if b.boardStatus.Board[0][0] == s &&
		b.boardStatus.Board[1][1] == s &&
		b.boardStatus.Board[2][2] == s {
		return true
	}

	if b.boardStatus.Board[0][2] == s &&
		b.boardStatus.Board[1][1] == s &&
		b.boardStatus.Board[2][0] == s {
		return true
	}

	return false
}
