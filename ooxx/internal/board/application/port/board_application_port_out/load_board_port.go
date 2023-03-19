package board_application_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"

type ILoadBoardPort interface {
	GetBoard() (domain.IBoard, error)
	SetBoard(domain.IBoard) error
}
