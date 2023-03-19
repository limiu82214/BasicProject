package board_application_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"

type ILoadBoardPort interface {
	GetBoard() (board_domain.IBoard, error)
	SetBoard(board_domain.IBoard) error
}
