package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_adapter/in/board_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/in/player_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/wire"
)

func main() {
	playerX()
}
func playerX() {
	gp := wire.InitPlayerAdapterInGoPrompt()
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
	gp := wire.InitBoardAdapterInGoPrompt()
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
