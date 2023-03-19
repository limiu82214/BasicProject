package in

type ISetPlayerInfoUseCase interface {
	SetPlayerInfo(cmd *SetPlayerInfoCmd) error
}
