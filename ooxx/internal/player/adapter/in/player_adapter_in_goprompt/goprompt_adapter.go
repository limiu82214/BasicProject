package player_adapter_in_goprompt

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application/port/in"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_domain"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
)

type IPlayerGopromptAdapter interface {
	PutChess()
	ShowBoard()
	ResetBoard()
	WhoWin()
	SetPlayerInfo()
}

type playerGopromptAdapter struct {
	getBoardStateUseCase in.IGetBoardStateUseCase
	putChessUseCase      in.IPutChessUseCase
	resetBoardUseCase    in.IResetBoardUseCase
	whoWinUseCase        in.IWhoWinUseCase
	setPlayerInfoUseCase in.ISetPlayerInfoUseCase
}

func NewPlayerGopromptAdapter(
	getBoardStateUseCase in.IGetBoardStateUseCase,
	putChessUseCase in.IPutChessUseCase,
	resetBoardUseCase in.IResetBoardUseCase,
	whoWinUseCase in.IWhoWinUseCase,
	setPlayerInfoUseCase in.ISetPlayerInfoUseCase,
) IPlayerGopromptAdapter {
	return &playerGopromptAdapter{
		getBoardStateUseCase: getBoardStateUseCase,
		putChessUseCase:      putChessUseCase,
		resetBoardUseCase:    resetBoardUseCase,
		whoWinUseCase:        whoWinUseCase,
		setPlayerInfoUseCase: setPlayerInfoUseCase,
	}
}

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "put", Description: "put chess"},
		{Text: "show", Description: "print board"},
		{Text: "reset", Description: "reset board"},
		{Text: "winner", Description: "show who win"},
		{Text: "setinfo", Description: "set info"},
		{Text: "q", Description: "exit"},
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func nullCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func (bpa *playerGopromptAdapter) ShowBoard() {
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

func (bpa *playerGopromptAdapter) ResetBoard() {
	err := bpa.resetBoardUseCase.ResetBoard()
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
	} else {
		log.Println("sys: reset board state done.")
	}
}

func (bpa *playerGopromptAdapter) PutChess() {
	nickname := prompt.Input("nickname: ", nullCompleter)
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

	ss := player_domain.State(s)

	ssc, err := in.NewPutChessCmd(nickname, x, y, ss)
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
		return
	}

	err = bpa.putChessUseCase.PutChess(ssc)
	if err != nil {
		log.Printf("sys: %s", err.Error())
	} else {
		log.Printf("sys: %s put chess, [%d][%d] will be %d.", nickname, x, y, s)
	}
}

func (bpa *playerGopromptAdapter) WhoWin() {
	winner, err := bpa.whoWinUseCase.WhoWin()
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	if winner == player_domain.Blank {
		log.Print("sys: nobody win")
	} else {
		log.Printf("sys: winner is %d", winner)
	}
}

func (bpa *playerGopromptAdapter) SetPlayerInfo() {
	nickname := prompt.Input("Nickname len(1~3): ", nullCompleter)

	cmd, err := in.NewSetPlayerInfoCmd(nickname)
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	err = bpa.setPlayerInfoUseCase.SetPlayerInfo(cmd)
	if err != nil {
		log.Panicln(err.Error())
		return
	}

	log.Printf("sys: set nickname(%s) success.", nickname)
}
