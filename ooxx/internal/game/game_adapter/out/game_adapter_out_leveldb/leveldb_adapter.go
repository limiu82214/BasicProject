package game_adapter_out_leveldb

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application/port/game_application_port_out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/gob"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
)

var errGetBoardFail = errors.New("get board fail")
var errSetBoardFail = errors.New("set board fail")
var errParseBoardToStoreFail = errors.New("parse board to store fail")
var errParseBoardFromStoreFail = errors.New("parse board from store fail")

type boardLevelDBAdapter struct {
	db *leveldb.DB
}

func New(db *leveldb.DB) game_application_port_out.ILoadBoardAdapter {
	return &boardLevelDBAdapter{
		db: db,
	}
}
func (bldba *boardLevelDBAdapter) SetBoard(b game_domain.IBoard) error {
	id := []byte("board_once")

	bs, err := b.GetBoardStatus()
	if err != nil {
		return errors.Wrap(err, "in leveldb_adapter SetBoard")
	}

	value, err := gob.StoreStructToByte(bs)
	if err != nil {
		return errors.Wrap(err, errParseBoardToStoreFail.Error())
	}

	err = bldba.db.Put(id, value, nil)
	if err != nil {
		return errSetBoardFail
	}

	return nil
}

func (bldba *boardLevelDBAdapter) GetBoard() (game_domain.IBoard, error) {
	id := []byte("board_once")

	data, err := bldba.db.Get(id, nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, game_domain.ErrGetEmpty
		}

		return nil, errors.Wrap(err, errGetBoardFail.Error())
	}

	if data == nil {
		return nil, game_domain.ErrGetEmpty
	}

	var s game_domain.BoardStatus

	err = gob.GetStrutFromByte(data, &s)
	if err != nil {
		return nil, errors.Wrap(err, errParseBoardFromStoreFail.Error())
	}

	b := game_domain.NewBoard()

	err = b.SetBoardStatus(&s)
	if err != nil {
		return nil, errors.Wrap(err, errParseBoardFromStoreFail.Error())
	}

	return b, nil
}
