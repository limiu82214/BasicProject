package user_adapter_in_goprompt

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/shared"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application/port/user_application_port_in"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/nerror"
)

type IUserGopromptAdapter interface {
	PutChess()
	ShowBoard()
	ResetBoard()
	WhoWin()
	SetPlayerInfo()
}

type userGopromptAdapter struct {
	getBoardStateUseCase user_application_port_in.IGetBoardStateUseCase
	putChessUseCase      user_application_port_in.IPutChessUseCase
	resetBoardUseCase    user_application_port_in.IResetBoardUseCase
	whoWinUseCase        user_application_port_in.IWhoWinUseCase
	setPlayerInfoUseCase user_application_port_in.ISetPlayerInfoUseCase
}

func New(
	getBoardStateUseCase user_application_port_in.IGetBoardStateUseCase,
	putChessUseCase user_application_port_in.IPutChessUseCase,
	resetBoardUseCase user_application_port_in.IResetBoardUseCase,
	whoWinUseCase user_application_port_in.IWhoWinUseCase,
	setPlayerInfoUseCase user_application_port_in.ISetPlayerInfoUseCase,
) IUserGopromptAdapter {
	return &userGopromptAdapter{
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

func (bpa *userGopromptAdapter) ShowBoard() {
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

func (bpa *userGopromptAdapter) ResetBoard() {
	err := bpa.resetBoardUseCase.ResetBoard()
	if err != nil {
		log.Printf("%v\n", nerror.PrettyError(err))
	} else {
		log.Println("sys: reset board state done.")
	}
}

func (bpa *userGopromptAdapter) PutChess() {
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

	ss := shared.State(s)

	ssc, err := user_application_port_in.NewPutChessCmd(nickname, x, y, ss)
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

func (bpa *userGopromptAdapter) WhoWin() {
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

func (bpa *userGopromptAdapter) SetPlayerInfo() {
	nickname := prompt.Input("Nickname len(1~3): ", nullCompleter)

	cmd, err := user_application_port_in.NewSetPlayerInfoCmd(nickname)
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