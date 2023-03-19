package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/in/user_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/wire"
)

func main() {
	playerX()
}
func playerX() {
	gp := wire.InitUserAdapterInGoPrompt()
leave:
	for {
		t := prompt.Input("action: ", user_adapter_in_goprompt.Completer)
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
		t := prompt.Input("action: ", game_adapter_in_goprompt.Completer)
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
