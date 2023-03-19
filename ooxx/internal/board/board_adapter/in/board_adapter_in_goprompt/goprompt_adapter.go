package board_adapter_in_goprompt

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/board_application/port/board_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
)

type IBoardGopromptAdapter interface {
	SetState()
	WhoWin()
	ShowBoard()
	ResetBoard()
}

type boardGopromptAdapter struct {
	setStateUseCase        board_application_port_in.ISetStateUseCase
	whoWinUseCase          board_application_port_in.IWhoWinUseCase
	getBoardStateUseCase   board_application_port_in.IGetBoardStateUseCase
	resetBoardStateUseCase board_application_port_in.IResetBoardStateUseCase
}

func New(
	setStateUserCase board_application_port_in.ISetStateUseCase,
	whoWinUseCase board_application_port_in.IWhoWinUseCase,
	getBoardStateUseCase board_application_port_in.IGetBoardStateUseCase,
	resetBoardStateUseCase board_application_port_in.IResetBoardStateUseCase,
) IBoardGopromptAdapter {
	return &boardGopromptAdapter{
		setStateUseCase:        setStateUserCase,
		whoWinUseCase:          whoWinUseCase,
		getBoardStateUseCase:   getBoardStateUseCase,
		resetBoardStateUseCase: resetBoardStateUseCase,
	}
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "set", Description: "set the state on board"},
		{Text: "winner", Description: "tell you O/X win or not"},
		{Text: "show", Description: "print board"},
		{Text: "reset", Description: "reset board"},
		{Text: "q", Description: "exit"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func nullCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (bpa *boardGopromptAdapter) ShowBoard() {
	bs, err := bpa.getBoardStateUseCase.GetBoardState()
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	b, err := json.Marshal(bs)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	ans := string(b[1 : len(b)-1])
	ans = "sys: board state\n" + strings.Join(strings.Split(ans, "],["), "]\n[")
	log.Println(ans)
}

func (bpa *boardGopromptAdapter) ResetBoard() {
	err := bpa.resetBoardStateUseCase.ResetBoardState()
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
	} else {
		log.Println("sys: reset board state done.")
	}
}

func (bpa *boardGopromptAdapter) SetState() {
	xStr := prompt.Input("x: ", nullCompleter)
	yStr := prompt.Input("y: ", nullCompleter)
	sStr := prompt.Input("state: ", nullCompleter)
	x, err := strconv.Atoi(xStr)

	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	s, err := strconv.Atoi(sStr)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	ss := shared.State(s)

	ssc, err := board_application_port_in.NewSetStateCmd(x, y, ss)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	err = bpa.setStateUseCase.SetState(ssc)
	if err != nil {
		log.Printf("sys: %s", err.Error())
	} else {
		bpa.ShowBoard()
		log.Printf("sys: [%d][%d] will be %d.", x, y, s)
	}
}

func (bpa *boardGopromptAdapter) WhoWin() {
	winner, err := bpa.whoWinUseCase.WhoWin()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	if winner == shared.Blank {
		log.Print("sys: nobody win")
	} else {
		log.Printf("sys: winner is %d", winner)
	}
}
