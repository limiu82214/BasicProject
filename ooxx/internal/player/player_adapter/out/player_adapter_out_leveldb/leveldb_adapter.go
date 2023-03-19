package player_adapter_out_leveldb

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/player_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/gob"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

var errInHere = errors.New("in player_adapter_out_leveldb")
var errGetFail = errors.New("get fail")
var errSetFail = errors.New("set fail")
var errParseToStoreFail = errors.New("parse to store fail")
var errParseFromStoreFail = errors.New("parse from store fail")

type playerLevelDBAdapter struct {
	db *leveldb.DB
}

func NewPlayerLevelDBAdapter(db *leveldb.DB) player_adapter_port_out.ILoadPlayerPort {
	return &playerLevelDBAdapter{
		db: db,
	}
}

func (pla *playerLevelDBAdapter) GetPlayer(nickname string) (player_domain.IPlayer, error) {
	id := []byte(nickname)

	data, err := pla.db.Get(id, nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, player_domain.ErrGetEmpty
		}

		return nil, errors.Wrap(err, errGetFail.Error())
	}

	if data == nil {
		return nil, player_domain.ErrGetEmpty
	}

	var ps player_domain.PlayerStore

	err = gob.GetStrutFromByte(data, &ps)
	if err != nil {
		return nil, errors.Wrap(err, errParseFromStoreFail.Error())
	}

	p := player_domain.NewPlayer()

	err = p.SetPlayerStore(&ps)
	if err != nil {
		return nil, errors.Wrap(err, errParseFromStoreFail.Error())
	}

	return p, nil
}

func (pla *playerLevelDBAdapter) SetPlayer(p player_domain.IPlayer) error {
	ps, err := p.GetPlayerStore()
	if err != nil {
		return errors.Wrap(err, errInHere.Error())
	}

	id := []byte(ps.NickName)

	value, err := gob.StoreStructToByte(ps)
	if err != nil {
		return errors.Wrap(err, errParseToStoreFail.Error())
	}

	err = pla.db.Put(id, value, nil)
	if err != nil {
		return errSetFail
	}

	return nil
}
