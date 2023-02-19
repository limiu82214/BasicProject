package test_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	leveldb_adapter "github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/out/leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

func Test_setState_SetState(t *testing.T) { //nolint // this is test
	db := leveldb.GetInst()
	ldba := leveldb_adapter.NewBoardLevelDBAdapter(db)
	nss := application.NewSetState(ldba)
	ngbs := application.NewGetBoardState(ldba)
	showBoard := func(bs [3][3]domain.State) {
		b, err := json.Marshal(bs)
		if err != nil {
			log.Println(err.Error())
		}

		ans := string(b[1 : len(b)-1])
		ans = "\n" + strings.Join(strings.Split(ans, "],["), "]\n[")
		log.Println(ans)
	}

	tt := [5][3]int{
		{0, 0, 1},
		{0, 1, 2},
		{1, 1, 1},
		{0, 2, 2},
		{2, 2, 1},
	}
	for _, t := range tt {
		cmd, err := in.NewSetStateCmd(t[0], t[1], domain.State(t[2]))
		if err != nil {
			log.Println(err.Error())
		}

		err = nss.SetState(cmd)
		if err != nil {
			log.Println(err.Error())
		} else {
			bs, err := ngbs.GetBoardState()
			if err != nil {
				log.Println(err.Error())
			} else {
				showBoard(bs)
			}
		}
	}
}
