package test_test

import (
	"encoding/json"
	"log"
	"strings"
	"testing"

	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

func Test_setState_SetState(t *testing.T) { //nolint // this is test
	nss := application.NewSetState()
	ngbs := application.NewGetBoardState()
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
			showBoard(ngbs.GetBoardState())
		}
	}
}
