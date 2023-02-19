package in

type ISetStateUseCase interface {
	SetState(cmd *SetStateCmd) error
}
