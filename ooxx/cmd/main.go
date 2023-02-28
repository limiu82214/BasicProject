package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	board_in_goprompt "github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/player"
	board_in_leveldb_adapter "github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/out/leveldb"
	board_application "github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
	player_in_goprompt "github.com/limiu82214/GoBasicProject/ooxx/internal/player/adapter/in/goprompt"
	player_board "github.com/limiu82214/GoBasicProject/ooxx/internal/player/adapter/out/board"
	player_in_leveldb_adapter "github.com/limiu82214/GoBasicProject/ooxx/internal/player/adapter/out/leveldb"
	player_application "github.com/limiu82214/GoBasicProject/ooxx/internal/player/application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

func main() {
	playerX()
}
func playerX() {
	// 做一個 board adapter in player
	db := leveldb.GetInst()
	bldba := board_in_leveldb_adapter.NewBoardLevelDBAdapter(db)
	nbpa := player.NewBoardPlayerAdapter(
		board_application.NewGetBoardState(bldba),
		board_application.NewSetState(bldba),
		board_application.NewResetBoardState(bldba),
		board_application.NewWhoWin(bldba),
	)

	// 將 board adapter in player 注入到 player adapter out board
	pb := player_board.NewPlayerBoardAdapter(nbpa)
	pldba := player_in_leveldb_adapter.NewPlayerLevelDBAdapter(db)
	// 將 player adapter out board 注入到 usecase
	gp := player_in_goprompt.NewPlayerGopromptAdapter(
		player_application.NewGetBoardState(pb),
		player_application.NewPutChess(pb, pldba),
		player_application.NewResetBoard(pb),
		player_application.NewWhoWin(pb),
		player_application.NewSetPlayerInfoService(pldba),
	)
leave:
	for {
		t := prompt.Input("action: ", player_in_goprompt.Completer)
		switch t {
		case "put":
			gp.PutChess()
		case "show":
			gp.ShowBoard()
		case "reset":
			gp.ResetBoard()
		case "winner":
			gp.WhoWin()
		case "setinfo":
			gp.SetPlayerInfo()
		case "q", "exit":
			break leave
		default:
			log.Println("err cmd. retry")
		}
	}
}

func boardX() { //nolint:unused // for test
	db := leveldb.GetInst()
	ldba := board_in_leveldb_adapter.NewBoardLevelDBAdapter(db)
	gp := board_in_goprompt.NewBoardGopromptAdapter(
		board_application.NewSetState(ldba),
		board_application.NewWhoWin(ldba),
		board_application.NewGetBoardState(ldba),
		board_application.NewResetBoardState(ldba),
	)
leave:
	for {
		t := prompt.Input("action: ", board_in_goprompt.Completer)
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
