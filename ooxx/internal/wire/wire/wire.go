//go:build wireinject
// +build wireinject

// wire ./internal/wire/wire/wire.go; mv ./internal/wire/wire/wire_gen.go ./internal/wire

package wire

import (
	"github.com/google/wire"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/in/game_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_adapter/out/game_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/game_application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/in/user_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/out/user_adapter_out_game"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_adapter/out/user_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/user/user_application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

var DBSet = wire.NewSet(
	leveldb.GetInst,
)

var BoardApplicationSet = wire.NewSet(
	game_application.NewGetBoardStateUseCase,
	game_application.NewSetStateUseCase,
	game_application.NewResetBoardStateUseCase,
	game_application.NewWhoWinUseCase,
)

var UserApplicationSet = wire.NewSet(
	user_application.NewGetBoardStateUseCase,
	user_application.NewPutChessUseCase,
	user_application.NewResetBoardUseCase,
	user_application.NewWhoWinUseCase,
	user_application.NewSetPlayerInfoUseCase,
)

func InitUserAdapterInGoPrompt() user_adapter_in_goprompt.IUserGopromptAdapter {
	panic(wire.Build(
		DBSet,
		game_adapter_out_leveldb.New,
		BoardApplicationSet,
		game_adapter_in_player.New,

		user_adapter_out_game.New,
		user_adapter_out_leveldb.New,
		UserApplicationSet,
		user_adapter_in_goprompt.New,
	))
}

func InitBoardAdapterInGoPrompt() game_adapter_in_goprompt.IBoardGopromptAdapter {
	panic(wire.Build(
		DBSet,
		game_adapter_out_leveldb.New,
		BoardApplicationSet,
		game_adapter_in_goprompt.New,
	))
}
