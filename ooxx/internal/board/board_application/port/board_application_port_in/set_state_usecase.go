package board_application_port_in

type ISetStateUseCase interface {
	SetState(cmd *SetStateCmd) error
}