package player_application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/pkg/errors"
)

type setPlayerInfo struct {
	loadPlayerPort player_adapter_port_out.ILoadPlayerPort
}

func NewSetPlayerInfoService(
	loadPlayerPort player_adapter_port_out.ILoadPlayerPort,
) player_application_port_in.ISetPlayerInfoUseCase {
	return &setPlayerInfo{
		loadPlayerPort: loadPlayerPort,
	}
}

func (spis *setPlayerInfo) SetPlayerInfo(cmd *player_application_port_in.SetPlayerInfoCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	p, err := spis.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, player_domain.ErrGetEmpty) {
		p = player_domain.NewPlayer()
	} else if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	err = p.SetInfo(cmd.Nickname)
	if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	err = spis.loadPlayerPort.SetPlayer(p)
	if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	return nil
}
