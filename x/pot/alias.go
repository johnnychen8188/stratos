package pot

import (
	"github.com/stratosnet/stratos-chain/x/pot/keeper"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

const (
	DefaultParamSpace = types.DefaultParamSpace
	ModuleName        = types.ModuleName
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	FoundationAccount = types.FoundationAccount
)

var (
	NewKeeper               = keeper.NewKeeper
	RegisterCodec           = types.RegisterLegacyAminoCodec
	ParamKeyTable           = types.ParamKeyTable
	NewGenesisState         = types.NewGenesisState
	NewMsgFoundationDeposit = types.NewMsgFoundationDeposit
)

type (
	Keeper = keeper.Keeper
)
