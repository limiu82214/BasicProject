package main

import (
	"log"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/adapter/in/goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application"
)

func main() {
	gp := goprompt.NewBoardGopromptAdapter(
		application.NewSetState(),
		application.NewWhoWin(),
		application.NewGetBoardState(),
		application.NewResetBoardState(),
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
		case "q":
			break leave
		default:
			log.Println("err cmd. retry")
		}
	}
}
