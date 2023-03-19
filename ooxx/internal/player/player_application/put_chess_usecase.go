package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_application")
var errShouldSetNicknameFirst = errors.New("should set nickname first")

type putChessUseCase struct {
	boardPort      player_adapter_port_out.IBoardAdapter
	loadPlayerPort player_adapter_port_out.ILoadPlayerAdapter
}

func NewPutChessUseCase(
	boardPort player_adapter_port_out.IBoardAdapter,
	loadPlayerPort player_adapter_port_out.ILoadPlayerAdapter,
) player_application_port_in.IPutChessUseCase {
	return &putChessUseCase{
		boardPort:      boardPort,
		loadPlayerPort: loadPlayerPort,
	}
}

func (p *putChessUseCase) PutChess(cmd *player_application_port_in.PutChessCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	_, err := p.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, player_domain.ErrGetEmpty) {
		return errShouldSetNicknameFirst
	}

	err = p.boardPort.SetBoardState(cmd.X, cmd.Y, int(cmd.S))

	return errors.Wrap(err, errInHere.Error())
}
