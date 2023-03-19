package board_application_port_in

type IResetBoardStateUseCase interface {
	ResetBoardState() error
}
