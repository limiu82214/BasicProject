package player_adapter_port_in

type ISetPlayerInfoUseCase interface {
	SetPlayerInfo(cmd *SetPlayerInfoCmd) error
}
