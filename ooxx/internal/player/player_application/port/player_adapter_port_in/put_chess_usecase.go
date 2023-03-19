package player_adapter_port_in

type IPutChessUseCase interface {
	PutChess(cmd *PutChessCmd) error
}
