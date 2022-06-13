package register

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/stratosnet/stratos-chain/x/register/keeper"
	"github.com/stratosnet/stratos-chain/x/register/types"
)

// InitGenesis initialize default parameters
// and the keeper's address to pubkey map
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data *types.GenesisState) {
	keeper.SetParams(ctx, *data.Params)

	initialStakeTotal := sdk.ZeroInt()
	lenOfGenesisBondedResourceNode := int64(0)
	for _, resourceNode := range data.GetResourceNodes() {
		if resourceNode.GetStatus() == stakingtypes.Bonded {
			lenOfGenesisBondedResourceNode++
			initialStakeTotal = initialStakeTotal.Add(resourceNode.Tokens)

		}
		keeper.SetResourceNode(ctx, resourceNode)
	}

	// set initial genesis number of resource nodes
	keeper.SetInitialGenesisBondedResourceNodeCnt(ctx, sdk.NewInt(lenOfGenesisBondedResourceNode))

	lenOfGenesisBondedMetaNode := int64(0)
	for _, metaNode := range data.GetMetaNodes() {
		if metaNode.GetStatus() == stakingtypes.Bonded {
			lenOfGenesisBondedResourceNode++
			initialStakeTotal = initialStakeTotal.Add(metaNode.Tokens)
		}
		keeper.SetMetaNode(ctx, metaNode)
	}
	// set initial genesis number of meta nodes
	keeper.SetInitialGenesisBondedMetaNodeCnt(ctx, sdk.NewInt(lenOfGenesisBondedMetaNode))

	totalUnissuedPrepay := keeper.GetTotalUnissuedPrepay(ctx).Amount
	initialUOzonePrice := sdk.ZeroDec()
	initialUOzonePrice = initialUOzonePrice.Add(data.InitialUozPrice)
	keeper.SetInitialGenesisStakeTotal(ctx, initialStakeTotal)
	keeper.SetInitialUOzonePrice(ctx, initialUOzonePrice)
	initOzoneLimit := initialStakeTotal.Add(totalUnissuedPrepay).ToDec().Quo(initialUOzonePrice).TruncateInt()
	keeper.SetRemainingOzoneLimit(ctx, initOzoneLimit)

	for _, slashing := range data.Slashing {
		walletAddress, err := sdk.AccAddressFromBech32(slashing.GetWalletAddress())
		if err != nil {
			panic(err)
		}

		keeper.SetSlashing(ctx, walletAddress, sdk.NewInt(slashing.Value))
	}
	return
}

// ExportGenesis writes the current store values
// to a genesis file, which can be imported again
// with InitGenesis
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) (data *types.GenesisState) {
	params := keeper.GetParams(ctx)

	resourceNodes := keeper.GetAllResourceNodes(ctx)
	metaNodes := keeper.GetAllMetaNodes(ctx)
	initialUOzonePrice := keeper.CurrUozPrice(ctx)

	var slashingInfo []*types.Slashing
	keeper.IteratorSlashingInfo(ctx, func(walletAddress sdk.AccAddress, val sdk.Int) (stop bool) {
		if val.GT(sdk.ZeroInt()) {
			slashing := types.NewSlashing(walletAddress, val)
			slashingInfo = append(slashingInfo, slashing)
		}
		return false
	})

	return &types.GenesisState{
		Params:          &params,
		ResourceNodes:   resourceNodes,
		MetaNodes:       metaNodes,
		InitialUozPrice: initialUOzonePrice,
		Slashing:        slashingInfo,
	}
}
