package domain

type IBoard interface {
	GetBoardState() [3][3]State
	ResetBoardState()
	SetState(x, y int, s State)
	WhoWin() State
}

type board struct {
	board [3][3]State
}

var boardOnce board //nolint:gochecknoglobals // only allow this board

func GetBoardInst() IBoard {
	b := board{}
	b.ResetBoardState()
	boardOnce = b

	return &boardOnce
}

func GetBoard() IBoard {
	return &boardOnce
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

func (b *board) SetState(x, y int, s State) {
	b.board[x][y] = s
}

func (b *board) WhoWin() State {
	if b.isWin(O) {
		return O
	}

	if b.isWin(X) {
		return X
	}

	return Blank
}

func (b *board) isWin(s State) bool { //nolint:cyclop // this is simplest
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
