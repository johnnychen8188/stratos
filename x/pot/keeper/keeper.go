package keeper

import (
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/stratosnet/stratos-chain/x/register"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

// Keeper of the pot store
type Keeper struct {
	storeKey         sdk.StoreKey
	cdc              *codec.Codec
	paramSpace       params.Subspace
	feeCollectorName string // name of the FeeCollector ModuleAccount
	BankKeeper       bank.Keeper
	SupplyKeeper     supply.Keeper
	AccountKeeper    auth.AccountKeeper
	StakingKeeper    staking.Keeper
	RegisterKeeper   register.Keeper
}

// NewKeeper creates a pot keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramSpace params.Subspace, feeCollectorName string,
	bankKeeper bank.Keeper, supplyKeeper supply.Keeper, accountKeeper auth.AccountKeeper, stakingKeeper staking.Keeper,
	registerKeeper register.Keeper,
) Keeper {
	keeper := Keeper{
		cdc:              cdc,
		storeKey:         key,
		paramSpace:       paramSpace.WithKeyTable(types.ParamKeyTable()),
		feeCollectorName: feeCollectorName,
		BankKeeper:       bankKeeper,
		SupplyKeeper:     supplyKeeper,
		AccountKeeper:    accountKeeper,
		StakingKeeper:    stakingKeeper,
		RegisterKeeper:   registerKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetVolumeReport(ctx sdk.Context, epoch sdk.Int) (res types.ReportRecord, err error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.VolumeReportStoreKey(epoch))
	if bz == nil {
		return types.ReportRecord{}, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"key %s does not exist", epoch)
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return res, nil
}

func (k Keeper) SetVolumeReport(ctx sdk.Context, epoch sdk.Int, reportRecord types.ReportRecord) {
	store := ctx.KVStore(k.storeKey)
	//storeKey := types.VolumeReportStoreKey(reporter)
	storeKey := types.VolumeReportStoreKey(epoch)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(reportRecord)
	store.Set(storeKey, bz)
}

func (k Keeper) DeleteVolumeReport(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(key)
}

func (k Keeper) IsSPNode(ctx sdk.Context, addr sdk.AccAddress) (found bool) {
	_, found = k.RegisterKeeper.GetIndexingNode(ctx, addr)
	return found
}

func getNodeOwnerMap(ctx sdk.Context, registerKeeper register.Keeper) map[string]sdk.AccAddress {

	resourceNodes := registerKeeper.GetAllResourceNodes(ctx)
	indexingNodes := registerKeeper.GetAllIndexingNodes(ctx)
	nodeOwnerMap := make(map[string]sdk.AccAddress)

	for _, v := range indexingNodes {
		nodeAddr := sdk.AccAddress(v.PubKey.Address())
		nodeOwnerMap[nodeAddr.String()] = v.OwnerAddress
	}

	for _, v := range resourceNodes {
		nodeAddr := sdk.AccAddress(v.PubKey.Address())
		nodeOwnerMap[nodeAddr.String()] = v.OwnerAddress
	}
	return nodeOwnerMap
}

func (k Keeper) setPotRewardRecordByOwnerHeight(ctx sdk.Context, nodeOwnerMap map[string]sdk.AccAddress, epoch sdk.Int, value []NodeRewardsInfo) {
	potRewardsRecordWithOwnerAddr := make(map[string][]NodeRewardsInfo)
	ctx.Logger().Info("In setPotRewardRecordByOwnerHeight")
	for _, v := range value {
		ctx.Logger().Info("In setPotRewardRecordByOwnerHeight first loop", "v", v)
		if ownerAddr, ok := nodeOwnerMap[v.NodeAddress.String()]; ok {
			ctx.Logger().Info("In setPotRewardRecordByOwnerHeight first loop", "ownerAddr", ownerAddr)
			if _, ok := potRewardsRecordWithOwnerAddr[ownerAddr.String()]; !ok {
				ctx.Logger().Info("In setPotRewardRecordByOwnerHeight first loop", "potRewardsRecordWithOwnerAddr[ownerAddr.String()]", ok)
				potRewardsRecordWithOwnerAddr[ownerAddr.String()] = []NodeRewardsInfo{}
			}

			ctx.Logger().Info("In setPotRewardRecordByOwnerHeight first loop", "potRewardsRecordWithOwnerAddr[ownerAddr.String()]", potRewardsRecordWithOwnerAddr[ownerAddr.String()])
			potRewardsRecordWithOwnerAddr[ownerAddr.String()] = append(potRewardsRecordWithOwnerAddr[ownerAddr.String()], v)
			ctx.Logger().Info("In setPotRewardRecordByOwnerHeight first loop", "potRewardsRecordWithOwnerAddrKey", ownerAddr, "Value", potRewardsRecordWithOwnerAddr[ownerAddr.String()])
		}
	}

	for key, val := range potRewardsRecordWithOwnerAddr {
		ctx.Logger().Info("In setPotRewardRecordByOwnerHeight second loop", "key", key, "val", val)
		k.setPotRewardRecord(ctx, key, epoch, val)
	}
}
