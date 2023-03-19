package player_application_port_in

type ISetPlayerInfoUseCase interface {
	SetPlayerInfo(cmd *SetPlayerInfoCmd) error
}
