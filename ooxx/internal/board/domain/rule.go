package domain

import "errors"

type iRule interface {
	whoWin(b *board) State
	isMyTune(b *board, s State) bool
	setState(b *board, x, y int, s State) error
}

type rule struct {
	lastState State
}

var errNotYourTurn = errors.New("is not your turn")

func (r *rule) setState(b *board, x, y int, s State) error {
	if !b.rule.isMyTune(b, s) {
		return errNotYourTurn
	}

	r.lastState = s

	return nil
}

func (r *rule) isMyTune(b *board, s State) bool {
	lastPlayState := State(r.lastState)
	if lastPlayState == Blank {
		return true
	}

	return lastPlayState != s
}

func (r *rule) whoWin(b *board) State {
	if r.isWin(b, O) {
		return O
	}

	if r.isWin(b, X) {
		return X
	}

	return Blank
}

func (r *rule) isWin(b *board, s State) bool { //nolint:cyclop // this is simplest
	// row, col
	for i := 0; i < 2; i++ {
		if b.board[i][0] == s &&
			b.board[i][1] == s &&
			b.board[i][2] == s {
			return true
		}

		if b.board[0][i] == s &&
			b.board[1][i] == s &&
			b.board[2][i] == s {
			return true
		}
	}

	if b.board[0][0] == s &&
		b.board[1][1] == s &&
		b.board[2][2] == s {
		return true
	}

	if b.board[0][2] == s &&
		b.board[1][1] == s &&
		b.board[2][0] == s {
		return true
	}

	return false
}
