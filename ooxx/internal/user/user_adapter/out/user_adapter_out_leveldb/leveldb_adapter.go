package user_adapter_out_leveldb

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_adapter_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_domain"
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

func New(db *leveldb.DB) user_adapter_port_out.ILoadPlayerAdapter {
	return &playerLevelDBAdapter{
		db: db,
	}
}

func (pla *playerLevelDBAdapter) GetPlayer(nickname string) (user_domain.IPlayer, error) {
	id := []byte(nickname)

	data, err := pla.db.Get(id, nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, user_domain.ErrGetEmpty
		}

		return nil, errors.Wrap(err, errGetFail.Error())
	}

	if data == nil {
		return nil, user_domain.ErrGetEmpty
	}

	var ps user_domain.PlayerStore

	err = gob.GetStrutFromByte(data, &ps)
	if err != nil {
		return nil, errors.Wrap(err, errParseFromStoreFail.Error())
	}

	p := user_domain.NewPlayer()

	err = p.SetPlayerStore(&ps)
	if err != nil {
		return nil, errors.Wrap(err, errParseFromStoreFail.Error())
	}

	return p, nil
}

func (pla *playerLevelDBAdapter) SetPlayer(p user_domain.IPlayer) error {
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
