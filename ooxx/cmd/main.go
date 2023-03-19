package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/in/board_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/in/board_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/out/board_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/in/player_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/out/player_adapter_out_board"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/out/player_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

func main() {
	playerX()
}
func playerX() {
	// 做一個 board adapter in player
	db := leveldb.GetInst()
	bldba := board_adapter_out_leveldb.New(db)
	nbpa := board_adapter_in_player.New(
		board_application.NewGetBoardState(bldba),
		board_application.NewSetState(bldba),
		board_application.NewResetBoardState(bldba),
		board_application.NewWhoWin(bldba),
	)

	// 將 board adapter in player 注入到 player adapter out board
	pb := player_adapter_out_board.NewPlayerBoardAdapter(nbpa)
	pldba := player_adapter_out_leveldb.NewPlayerLevelDBAdapter(db)
	// 將 player adapter out board 注入到 usecase
	gp := player_adapter_in_goprompt.NewPlayerGopromptAdapter(
		player_application.NewGetBoardState(pb),
		player_application.NewPutChess(pb, pldba),
		player_application.NewResetBoard(pb),
		player_application.NewWhoWin(pb),
		player_application.NewSetPlayerInfoService(pldba),
	)
leave:
	for {
		t := prompt.Input("action: ", player_adapter_in_goprompt.Completer)
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
	ldba := board_adapter_out_leveldb.New(db)
	gp := board_adapter_in_goprompt.New(
		board_application.NewSetState(ldba),
		board_application.NewWhoWin(ldba),
		board_application.NewGetBoardState(ldba),
		board_application.NewResetBoardState(ldba),
	)
leave:
	for {
		t := prompt.Input("action: ", board_adapter_in_goprompt.Completer)
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
