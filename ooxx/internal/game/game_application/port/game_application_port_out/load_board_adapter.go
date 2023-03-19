package game_application_port_out

import "github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_domain"

type ILoadBoardAdapter interface {
	GetBoard() (game_domain.IBoard, error)
	SetBoard(game_domain.IBoard) error
}
