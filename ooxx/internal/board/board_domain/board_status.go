package board_domain

import "github.com/limiu82214/GoBasicProject/ooxx/internal/shared"

type BoardStatus struct {
	Board     [3][3]shared.State
	LastState shared.State
}
