package scavenge

import (
	"github.com/yxuco/cosmos-scavenge/x/scavenge/keeper"
	"github.com/yxuco/cosmos-scavenge/x/scavenge/types"
)

const (
	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	// QueryParams       = types.QueryParams
	QuerierRoute = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc

	NewMsgCreateScavenge = types.NewMsgCreateScavenge
	NewMsgCommitSolution = types.NewMsgCommitSolution
	NewMsgRevealSolution = types.NewMsgRevealSolution
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgCreateScavenge = types.MsgCreateScavenge
	MsgCommitSolution = types.MsgCommitSolution
	MsgRevealSolution = types.MsgRevealSolution
)
