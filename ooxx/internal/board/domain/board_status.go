package domain

type BoardStatus struct {
	Board     [3][3]State
	LastState State
}