package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/goprompt"
	leveldb_adapter "github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/out/leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

func main() {
	db := leveldb.GetInst()
	ldba := leveldb_adapter.NewBoardLevelDBAdapter(db)
	gp := goprompt.NewBoardGopromptAdapter(
		application.NewSetState(ldba),
		application.NewWhoWin(ldba),
		application.NewGetBoardState(ldba),
		application.NewResetBoardState(ldba),
	)
leave:
	for {
		t := prompt.Input("action: ", goprompt.Completer)
		switch t {
		case "set":
			gp.SetState()
		case "winner":
			gp.WhoWin()
		case "show":
			gp.ShowBoard()
		case "reset":
			gp.ResetBoard()
		case "q", "exit":
			break leave
		default:
			log.Println("err cmd. retry")
		}
	}
}
