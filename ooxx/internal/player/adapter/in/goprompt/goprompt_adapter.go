package goprompt

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/application/port/in"
)

type IPlayerGopromptAdapter interface {
	PutChess()
	ShowBoard()
	ResetBoard()
}

type playerGopromptAdapter struct {
	getBoardStateUseCase in.IGetBoardStateUseCase
}

func NewPlayerGopromptAdapter(getBoardStateUseCase in.IGetBoardStateUseCase) IPlayerGopromptAdapter {
	return &playerGopromptAdapter{
		getBoardStateUseCase: getBoardStateUseCase,
	}
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "put", Description: "put chess"},
		{Text: "show", Description: "print board"},
		{Text: "reset", Description: "reset board"},
		{Text: "q", Description: "exit"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (bpa *playerGopromptAdapter) ShowBoard() {
	bs, err := bpa.getBoardStateUseCase.GetBoardState()
	if err != nil {
		log.Println(err.Error())

		return
	}

	b, err := json.Marshal(bs)
	if err != nil {
		log.Println(err.Error())
	}

	ans := string(b[1 : len(b)-1])
	ans = "sys: board state\n" + strings.Join(strings.Split(ans, "],["), "]\n[")
	log.Println(ans)
}

func (bpa *playerGopromptAdapter) ResetBoard() {
	// err := bpa.resetPlayerStateUseCase.ResetPlayerState()
	// if err != nil {
	// 	log.Println(err.Error())
	// } else {
	// 	log.Println("sys: reset board state done.")
	// }
}

func (bpa *playerGopromptAdapter) PutChess() {
	// xStr := prompt.Input("x: ", nullCompleter)
	// yStr := prompt.Input("y: ", nullCompleter)
	// sStr := prompt.Input("state: ", nullCompleter)
	// x, err := strconv.Atoi(xStr)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// y, err := strconv.Atoi(yStr)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// s, err := strconv.Atoi(sStr)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// ss := domain.State(s)

	// ssc, err := in.NewSetStateCmd(x, y, ss)
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// err = bpa.setStateUseCase.SetState(ssc)
	// if err != nil {
	// 	log.Printf("sys: %s", err.Error())
	// } else {
	// 	bpa.ShowPlayer()
	// 	log.Printf("sys: [%d][%d] will be %d.", x, y, s)
	// }
}
