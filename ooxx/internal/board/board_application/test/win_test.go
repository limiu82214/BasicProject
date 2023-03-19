package test_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/out/board_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
)

func Test_setState_SetState(t *testing.T) {
	db := leveldb.GetInst()
	ldba := board_adapter_out_leveldb.New(db)
	nss := board_application.NewSetStateUseCase(ldba)
	ngbs := board_application.NewGetBoardStateUseCase(ldba)
	showBoard := func(bs [3][3]board_domain.State) {
		b, err := json.Marshal(bs)
		if err != nil {
			log.Printf("%v\n", nerror.PrettyError(err))
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
		cmd, err := board_application_port_in.NewSetStateCmd(t[0], t[1], board_domain.State(t[2]))
		if err != nil {
			log.Printf("%v\n", nerror.PrettyError(err))
		}

		err = nss.SetState(cmd)
		if err != nil {
			log.Printf("%v\n", nerror.PrettyError(err))
		} else {
			bs, err := ngbs.GetBoardState()
			if err != nil {
				log.Printf("%v\n", nerror.PrettyError(err))
			} else {
				showBoard(bs)
			}
		}
	}
}
