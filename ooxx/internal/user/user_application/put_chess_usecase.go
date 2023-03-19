package user_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_domain"
	"github.com/pkg/errors"
)

var errInHere = errors.New("in player_application")
var errShouldSetNicknameFirst = errors.New("should set nickname first")

type putChessUseCase struct {
	boardPort      user_adapter_port_out.IBoardAdapter
	loadPlayerPort user_adapter_port_out.ILoadPlayerAdapter
}

func NewPutChessUseCase(
	boardPort user_adapter_port_out.IBoardAdapter,
	loadPlayerPort user_adapter_port_out.ILoadPlayerAdapter,
) user_application_port_in.IPutChessUseCase {
	return &putChessUseCase{
		boardPort:      boardPort,
		loadPlayerPort: loadPlayerPort,
	}
}

func (p *putChessUseCase) PutChess(cmd *user_application_port_in.PutChessCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	_, err := p.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, user_domain.ErrGetEmpty) {
		return errShouldSetNicknameFirst
	}

	err = p.boardPort.SetBoardState(cmd.X, cmd.Y, int(cmd.S))

	return errors.Wrap(err, errInHere.Error())
}
