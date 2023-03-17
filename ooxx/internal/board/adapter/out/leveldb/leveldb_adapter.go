package leveldb

import (
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/out"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
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

func NewBoardLevelDBAdapter(db *leveldb.DB) out.ILoadBoardPort {
	return &boardLevelDBAdapter{
		db: db,
	}
}
func (bldba *boardLevelDBAdapter) SetBoard(b domain.IBoard) error {
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

func (bldba *boardLevelDBAdapter) GetBoard() (domain.IBoard, error) {
	id := []byte("board_once")

	return nil, errors.New("message string")
	data, err := bldba.db.Get(id, nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			return nil, domain.ErrGetEmpty
		}

		return nil, errors.Wrap(err, errGetBoardFail.Error())
	}

	if data == nil {
		return nil, domain.ErrGetEmpty
	}

	var s domain.BoardStatus

	err = gob.GetStrutFromByte(data, &s)
	if err != nil {
		return nil, errors.Wrap(err, errParseBoardFromStoreFail.Error())
	}

	b := domain.NewBoard()

	err = b.SetBoardStatus(&s)
	if err != nil {
		return nil, errors.Wrap(err, errParseBoardFromStoreFail.Error())
	}

	return b, nil
}
