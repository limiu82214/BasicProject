package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/player"
	leveldb_adapter "github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/out/leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
	player_goprompt "github.com/limiu82214/GoBasicProject/ooxx/internal/player/adapter/in/goprompt"
	player_board "github.com/limiu82214/GoBasicProject/ooxx/internal/player/adapter/out/board"
	player_application "github.com/limiu82214/GoBasicProject/ooxx/internal/player/application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

func main() {
	playerX()
}
func playerX() {
	// 做一個 board adapter in player
	db := leveldb.GetInst()
	ldba := leveldb_adapter.NewBoardLevelDBAdapter(db)
	nbpa := player.NewBoardPlayerAdapter(
		application.NewGetBoardState(ldba),
	)

	// 將 board adapter in player 注入到 player adapter out board
	pb := player_board.NewPlayerBoardAdapter(nbpa)
	// 將 player adapter out board 注入到 usecase
	gp := player_goprompt.NewPlayerGopromptAdapter(
		player_application.NewGetBoardState(pb),
	)
leave:
	for {
		t := prompt.Input("action: ", goprompt.Completer)
		switch t {
		// case "put":
		// 	gp.PutChess()
		case "show":
			gp.ShowBoard()
		// case "reset":
		// 	gp.ResetBoard()
		case "q", "exit":
			break leave
		default:
			log.Println("err cmd. retry")
		}
	}
}

func boardX() { //nolint:unused // for test
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
