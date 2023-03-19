package user_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_domain"
	"github.com/pkg/errors"
)

type setPlayerInfoUseCase struct {
	loadPlayerPort user_adapter_port_out.ILoadPlayerAdapter
}

func NewSetPlayerInfoUseCase(
	loadPlayerPort user_adapter_port_out.ILoadPlayerAdapter,
) user_application_port_in.ISetPlayerInfoUseCase {
	return &setPlayerInfoUseCase{
		loadPlayerPort: loadPlayerPort,
	}
}

func (s *setPlayerInfoUseCase) SetPlayerInfo(cmd *user_application_port_in.SetPlayerInfoCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	p, err := s.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, user_domain.ErrGetEmpty) {
		p = user_domain.NewPlayer()
	} else if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	err = p.SetInfo(cmd.Nickname)
	if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	err = s.loadPlayerPort.SetPlayer(p)
	if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	return nil
}
