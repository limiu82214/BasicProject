//go:build wireinject
// +build wireinject

// wire ./internal/wire/wire/wire.go; mv ./internal/wire/wire/wire_gen.go ./internal/wire

package wire

import (
	"github.com/google/wire"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/board_adapter/in/board_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/board_adapter/in/board_adapter_in_player"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/board_adapter/out/board_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/game/board_application"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/in/player_adapter_in_goprompt"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/out/player_adapter_out_board"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_adapter/out/player_adapter_out_leveldb"
	"github.com/limiu82214/GoBasicProject/ooxx/internal/player/player_application"
	"github.com/limiu82214/GoBasicProject/ooxx/pkg/leveldb"
)

var DBSet = wire.NewSet(
	leveldb.GetInst,
)

var BoardApplicationSet = wire.NewSet(
	board_application.NewGetBoardStateUseCase,
	board_application.NewSetStateUseCase,
	board_application.NewResetBoardStateUseCase,
	board_application.NewWhoWinUseCase,
)

var PlayerApplicationSet = wire.NewSet(
	player_application.NewGetBoardStateUseCase,
	player_application.NewPutChessUseCase,
	player_application.NewResetBoardUseCase,
	player_application.NewWhoWinUseCase,
	player_application.NewSetPlayerInfoUseCase,
)

func InitPlayerAdapterInGoPrompt() player_adapter_in_goprompt.IPlayerGopromptAdapter {
	panic(wire.Build(
		DBSet,
		board_adapter_out_leveldb.New,
		BoardApplicationSet,
		board_adapter_in_player.New,

		player_adapter_out_board.New,
		player_adapter_out_leveldb.New,
		PlayerApplicationSet,
		player_adapter_in_goprompt.New,
	))
}

func InitBoardAdapterInGoPrompt() board_adapter_in_goprompt.IBoardGopromptAdapter {
	panic(wire.Build(
		DBSet,
		board_adapter_out_leveldb.New,
		BoardApplicationSet,
		board_adapter_in_goprompt.New,
	))
}
