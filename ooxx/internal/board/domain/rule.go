package domain

import "errors"

type iRule interface {
	whoWin(b *Board) State
	isMyTune(b *Board, s State) bool
	setState(b *Board, x, y int, s State) error
}

type rule struct {
}

var errNotYourTurn = errors.New("is not your turn")
var errIsNotBlankPos = errors.New("this pos is not Blank")

func (r *rule) setState(b *Board, x, y int, s State) error {
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
	return b.boardStatus.Board[x][y] == Blank
}

func (r *rule) isMyTune(b *Board, s State) bool {
	lastPlayState := State(b.boardStatus.LastState)
	if lastPlayState == Blank {
		return true
	}

	return lastPlayState != s
}

func (r *rule) whoWin(b *Board) State {
	if r.isWin(b, O) {
		return O
	}

	if r.isWin(b, X) {
		return X
	}

	return Blank
}

func (r *rule) isWin(b *Board, s State) bool { //nolint:cyclop // this is simplest
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
