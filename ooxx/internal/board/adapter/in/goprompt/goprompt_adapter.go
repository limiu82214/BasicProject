package goprompt

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/board/domain"
)

type IBoardGopromptAdapter interface {
	SetState()
	WhoWin()
	ShowBoard()
}

type boardGopromptAdapter struct {
	setStateUseCase        in.ISetStateUseCase
	whoWinUseCase          in.IWhoWinUseCase
	getBoardStateUseCase   in.IGetBoardStateUseCase
	resetBoardStateUseCase in.IResetBoardStateUseCase
}

func NewBoardGopromptAdapter(
	setStateUserCase in.ISetStateUseCase,
	whoWinUseCase in.IWhoWinUseCase,
	getBoardStateUseCase in.IGetBoardStateUseCase,
	resetBoardStateUseCase in.IResetBoardStateUseCase,
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
		{Text: "q", Description: "exit"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
func nullCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (bpa *boardGopromptAdapter) ShowBoard() {
	bs := bpa.getBoardStateUseCase.GetBoardState()
	// TODO: 檢查回傳的bs修改其值會不會影響到整個board

	b, err := json.Marshal(bs)
	if err != nil {
		log.Println(err.Error())
	}

	ans := string(b[1 : len(b)-1])
	ans = "sys: board state\n" + strings.Join(strings.Split(ans, "],["), "]\n[")
	log.Println(ans)
}

func (bpa *boardGopromptAdapter) ResetBoard() {
	bpa.resetBoardStateUseCase.ResetBoardState()
	log.Println("sys: reset board state done.")
}

func (bpa *boardGopromptAdapter) SetState() {
	xStr := prompt.Input("x: ", nullCompleter)
	yStr := prompt.Input("y: ", nullCompleter)
	sStr := prompt.Input("state: ", nullCompleter)
	x, err := strconv.Atoi(xStr)

	if err != nil {
		log.Println(err.Error())
	}

	y, err := strconv.Atoi(yStr)
	if err != nil {
		log.Println(err.Error())
	}

	s, err := strconv.Atoi(sStr)
	if err != nil {
		log.Println(err.Error())
	}

	ss := domain.State(s)

	ssc, err := in.NewSetStateCmd(x, y, ss)
	if err != nil {
		log.Println(err.Error())
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
	winner := bpa.whoWinUseCase.WhoWin()
	if winner == domain.Blank {
		log.Print("sys: nobody win")
	} else {
		log.Printf("sys: winner is %d", winner)
	}
}
