package in

type IPutChessUseCase interface {
	PutChess(cmd *PutChessCmd) error
}
