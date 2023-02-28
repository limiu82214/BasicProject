package application

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/domain"
	"github.com/pkg/errors"
)

type setPlayerInfo struct {
	loadPlayerPort out.ILoadPlayerPort
}

func NewSetPlayerInfoService(
	loadPlayerPort out.ILoadPlayerPort,
) in.ISetPlayerInfoUseCase {
	return &setPlayerInfo{
		loadPlayerPort: loadPlayerPort,
	}
}

func (spis *setPlayerInfo) SetPlayerInfo(cmd *in.SetPlayerInfoCmd) error {
	if !cmd.IsValid() {
		panic("檢查是本基")
	}

	p, err := spis.loadPlayerPort.GetPlayer(cmd.Nickname)
	if err != nil && errors.Is(err, domain.ErrGetEmpty) {
		p = domain.NewPlayer()
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
