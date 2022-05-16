package pot

//
//import (
//	"testing"
//
//	stratos "github.com/stratosnet/stratos-chain/types"
//	"github.com/stretchr/testify/require"
//
//	abci "github.com/tendermint/tendermint/abci/types"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/bank"
//	"github.com/cosmos/cosmos-sdk/x/mock"
//	"github.com/cosmos/cosmos-sdk/x/staking"
//	"github.com/cosmos/cosmos-sdk/x/supply"
//	supplyexported "github.com/cosmos/cosmos-sdk/x/supply/exported"
//	"github.com/stratosnet/stratos-chain/x/pot/types"
//	"github.com/stratosnet/stratos-chain/x/register"
//)
//
//const (
//	stopFlagOutOfTotalMiningReward = true
//	stopFlagSpecificMinedReward    = false
//	stopFlagSpecificEpoch          = true
//)
//
//var (
//	paramSpecificMinedReward = sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(160000000000)))
//	paramSpecificEpoch       = sdk.NewInt(10)
//)
//
//// initialize data of volume report
//func setupMsgVolumeReport(newEpoch int64) types.MsgVolumeReport {
//	volume1 := types.NewSingleWalletVolume(resOwner1, resourceNodeVolume1)
//	volume2 := types.NewSingleWalletVolume(resOwner2, resourceNodeVolume2)
//	volume3 := types.NewSingleWalletVolume(resOwner3, resourceNodeVolume3)
//
//	nodesVolume := []types.SingleWalletVolume{volume1, volume2, volume3}
//	reporter := idxNodeNetworkId1
//	epoch := sdk.NewInt(newEpoch)
//	reportReference := "report for epoch " + epoch.String()
//	reporterOwner := idxOwner1
//
//	pubKeys := make([][]byte, 1)
//	for i := range pubKeys {
//		pubKeys[i] = make([]byte, 1)
//	}
//
//	signature := types.NewBLSSignatureInfo(pubKeys, []byte("signature"), []byte("txData"))
//
//	volumeReportMsg := types.NewMsgVolumeReport(nodesVolume, reporter, epoch, reportReference, reporterOwner, signature)
//
//	return volumeReportMsg
//}
//
//func setupSlashingMsg() types.MsgSlashingResourceNode {
//	reporters := make([]stratos.SdsAddress, 0)
//	reporters = append(reporters, idxNodeNetworkId1)
//	reportOwner := make([]sdk.AccAddress, 0)
//	reportOwner = append(reportOwner, idxOwner1)
//
//	slashingMsg := types.NewMsgSlashingResourceNode(reporters, reportOwner, resNodeNetworkId1, resOwner1, resNodeSlashingUOZAmt1, true)
//	return slashingMsg
//}
//
//// Test case termination conditions
//// modify stop flag & variable could make the test case stop when reach a specific condition
//func isNeedStop(ctx sdk.Context, k Keeper, epoch sdk.Int, minedToken sdk.Coin) bool {
//
//	if stopFlagOutOfTotalMiningReward && (minedToken.Amount.GT(foundationDeposit.AmountOf(k.BondDenom(ctx))) ||
//		minedToken.Amount.GT(foundationDeposit.AmountOf(k.RewardDenom(ctx)))) {
//		return true
//	}
//	if stopFlagSpecificMinedReward && minedToken.Amount.GT(paramSpecificMinedReward.AmountOf(k.BondDenom(ctx))) {
//		return true
//	}
//	if stopFlagSpecificEpoch && epoch.GT(paramSpecificEpoch) {
//		return true
//	}
//	return false
//}
//
//func TestPotVolumeReportMsgs(t *testing.T) {
//	/********************* initialize mock app *********************/
//	mApp, k, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper := getMockApp(t)
//	accs := setupAccounts(mApp)
//	mock.SetGenesis(mApp, accs)
//
//	/********************* foundation account deposit *********************/
//	header := abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx := mApp.BaseApp.NewContext(true, header)
//	foundationDepositMsg := NewMsgFoundationDeposit(foundationDeposit, foundationDepositorAccAddr)
//	foundationDepositorAcc := mApp.AccountKeeper.GetAccount(ctx, foundationDepositorAccAddr)
//	accNum := foundationDepositorAcc.GetAccountNumber()
//	accSeq := foundationDepositorAcc.GetSequence()
//	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{foundationDepositMsg}, []uint64{accNum}, []uint64{accSeq}, true, true, foundationDepositorPrivKey)
//	foundationAccAddr := supplyKeeper.GetModuleAddress(types.FoundationAccount)
//	mock.CheckBalance(t, mApp, foundationAccAddr, foundationDeposit)
//
//	/********************* create validator with 50% commission *********************/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//
//	commission := staking.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
//	description := staking.NewDescription("foo_moniker", "", "", "", "")
//	createValidatorMsg := staking.NewMsgCreateValidator(valOpValAddr1, valConsPubk1, sdk.NewCoin("ustos", valInitialStake), description, commission, sdk.OneInt())
//
//	valOpAcc1 := mApp.AccountKeeper.GetAccount(ctx, valOpAccAddr1)
//	accNum = valOpAcc1.GetAccountNumber()
//	accSeq = valOpAcc1.GetSequence()
//	mock.SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{createValidatorMsg}, []uint64{accNum}, []uint64{accSeq}, true, true, valOpPrivKey1)
//	mock.CheckBalance(t, mApp, valOpAccAddr1, nil)
//
//	/********************** commit **********************/
//	header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//	ctx = mApp.BaseApp.NewContext(true, header)
//
//	mApp.BeginBlock(abci.RequestBeginBlock{Header: header})
//	stakingKeeper.ApplyAndReturnValidatorSetUpdates(mApp.BaseApp.NewContext(true, header))
//	validator := checkValidator(t, mApp, stakingKeeper, valOpValAddr1, true)
//
//	require.Equal(t, valOpValAddr1, validator.OperatorAddress)
//	require.Equal(t, sdk.Bonded, validator.Status)
//	require.True(sdk.IntEq(t, valInitialStake, validator.BondedTokens()))
//
//	/********************** loop sending volume report **********************/
//	var i int64
//	var slashingAmtSetup sdk.Int
//	i = 0
//	slashingAmtSetup = sdk.ZeroInt()
//	for {
//
//		/********************* test slashing msg when i==2 *********************/
//		if i == 2 {
//			ctx.Logger().Info("********************************* Deliver Slashing Tx START ********************************************")
//			slashingMsg := setupSlashingMsg()
//			/********************* deliver tx *********************/
//
//			idxOwnerAcc1 := mApp.AccountKeeper.GetAccount(ctx, idxOwner1)
//			ownerAccNum := idxOwnerAcc1.GetAccountNumber()
//			ownerAccSeq := idxOwnerAcc1.GetSequence()
//
//			SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{slashingMsg}, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, idxOwnerPrivKey1)
//			/********************* commit & check result *********************/
//			header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//			ctx = mApp.BaseApp.NewContext(true, header)
//
//			slashingAmtSetup = registerKeeper.GetSlashing(ctx, resOwner1)
//
//			_, slashingAmtCheck := k.GetTrafficReward(ctx, []types.SingleWalletVolume{{
//				WalletAddress: resOwner1,
//				Volume:        resNodeSlashingUOZAmt1,
//			}})
//			println("slashingAmtSetup=" + slashingAmtSetup.String())
//			require.Equal(t, slashingAmtSetup, slashingAmtCheck.TruncateInt())
//
//			ctx.Logger().Info("********************************* Deliver Slashing Tx END ********************************************")
//		}
//
//		ctx.Logger().Info("*****************************************************************************")
//		/********************* prepare tx data *********************/
//		volumeReportMsg := setupMsgVolumeReport(i + 1)
//
//		lastTotalMinedToken := k.GetTotalMinedTokens(ctx)
//		ctx.Logger().Info("last committed mined token = " + lastTotalMinedToken.String())
//		if isNeedStop(ctx, k, volumeReportMsg.Epoch, lastTotalMinedToken) {
//			break
//		}
//
//		/********************* print info *********************/
//		ctx.Logger().Info("epoch " + volumeReportMsg.Epoch.String())
//		S := k.RegisterKeeper.GetInitialGenesisStakeTotal(ctx).ToDec()
//		Pt := k.RegisterKeeper.GetTotalUnissuedPrepay(ctx).Amount.ToDec()
//		Y := k.GetTotalConsumedUoz(volumeReportMsg.WalletVolumes).ToDec()
//		Lt := k.RegisterKeeper.GetRemainingOzoneLimit(ctx).ToDec()
//		R := S.Add(Pt).Mul(Y).Quo(Lt.Add(Y))
//		//ctx.Logger().Info("R = (S + Pt) * Y / (Lt + Y)")
//		ctx.Logger().Info("S=" + S.String() + "\nPt=" + Pt.String() + "\nY=" + Y.String() + "\nLt=" + Lt.String() + "\nR=" + R.String() + "\n")
//
//		ctx.Logger().Info("---------------------------")
//		distributeGoal := types.InitDistributeGoal()
//		_, distributeGoal, err := k.CalcTrafficRewardInTotal(ctx, volumeReportMsg.WalletVolumes, distributeGoal)
//		require.NoError(t, err)
//
//		//TODO: recovery when shift to main net
//		/********************************************************** Main net part Start *********************************************************************/
//		distributeGoal, err = k.CalcMiningRewardInTotal(ctx, distributeGoal) //for main net
//		require.NoError(t, err)
//		ctx.Logger().Info(distributeGoal.String())
//
//		ctx.Logger().Info("---------------------------")
//		distributeGoalBalance := distributeGoal
//		rewardDetailMap := make(map[string]types.Reward)
//		rewardDetailMap, distributeGoalBalance = k.CalcRewardForResourceNode(ctx, volumeReportMsg.WalletVolumes, distributeGoalBalance, rewardDetailMap)
//		/********************************************************** Main net part End *********************************************************************/
//
//		//TODO: remove when shift to main net
//		/********************************************************** Incentive testnet part Start *********************************************************************/
//		//distributeGoal, idxNodeCnt, resNodeCnt, err := k.CalcMiningRewardInTotalForTestnet(ctx, distributeGoal) //for incentive test net
//		//require.NoError(t, err)
//		//ctx.Logger().Info(distributeGoal.String())
//		//ctx.Logger().Info("---------------------------")
//		//distributeGoalBalance := distributeGoal
//		//rewardDetailMap := make(map[string]types.Reward)
//		//
//		//rewardDetailMap, distributeGoalBalance = k.CalcRewardForResourceNodeForTestnet(ctx, volumeReportMsg.WalletVolumes, distributeGoalBalance, rewardDetailMap, resNodeCnt)
//		//rewardDetailMap, distributeGoalBalance = k.CalcRewardForIndexingNodeForTestnet(ctx, distributeGoalBalance, rewardDetailMap, idxNodeCnt)
//		//
//		////calc mining reward to distribute to validators
//		//rewardFromMiningPool := distributeGoal.BlockChainRewardToValidatorFromMiningPool
//		//usedRewardFromMiningPool := sdk.NewCoin(k.RewardDenom(ctx), sdk.ZeroInt())
//		//validatorWalletList := make([]sdk.AccAddress, 0)
//		//validators := k.StakingKeeper.GetAllValidators(ctx)
//		//for _, validator := range validators {
//		//	if validator.IsBonded() && !validator.IsJailed() {
//		//		validatorWalletList = append(validatorWalletList, sdk.AccAddress(validator.GetOperator()))
//		//	}
//		//}
//		//rewardPerValidator := sdk.NewCoin(k.RewardDenom(ctx), rewardFromMiningPool.Amount.ToDec().Quo(sdk.NewDec(int64(len(validatorWalletList)))).TruncateInt())
//		//usedRewardFromMiningPool = sdk.NewCoin(k.RewardDenom(ctx), rewardPerValidator.Amount.Mul(sdk.NewInt(int64(len(validatorWalletList)))))
//		/********************************************************** Incentive testnet part End *********************************************************************/
//
//		ctx.Logger().Info("resource_wallet1:  address = " + resOwner1.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[resOwner1.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[resOwner1.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("resource_wallet2:  address = " + resOwner2.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[resOwner2.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[resOwner2.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("resource_wallet3:  address = " + resOwner3.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[resOwner3.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[resOwner3.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("resource_wallet4:  address = " + resOwner4.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[resOwner4.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[resOwner4.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("resource_wallet5:  address = " + resOwner5.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[resOwner5.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[resOwner5.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("indexing_wallet1:  address = " + idxOwner1.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[idxOwner1.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[idxOwner1.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("indexing_wallet2:  address = " + idxOwner2.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[idxOwner2.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[idxOwner2.String()].RewardFromTrafficPool.String())
//
//		ctx.Logger().Info("indexing_wallet3:  address = " + idxOwner3.String())
//		ctx.Logger().Info("           miningReward = " + rewardDetailMap[idxOwner3.String()].RewardFromMiningPool.String())
//		ctx.Logger().Info("          trafficReward = " + rewardDetailMap[idxOwner3.String()].RewardFromTrafficPool.String())
//		ctx.Logger().Info("---------------------------")
//
//		/********************* record data before delivering tx  *********************/
//		feePoolAccAddr := supplyKeeper.GetModuleAddress(auth.FeeCollectorName)
//		lastFoundationAccBalance := bankKeeper.GetCoins(ctx, foundationAccAddr)
//		lastFeePool := bankKeeper.GetCoins(ctx, feePoolAccAddr)
//		lastUnissuedPrepay := k.RegisterKeeper.GetTotalUnissuedPrepay(ctx)
//		lastMatureTotalOfResNode1 := k.GetMatureTotalReward(ctx, resOwner2)
//
//		/********************* deliver tx *********************/
//
//		idxOwnerAcc1 := mApp.AccountKeeper.GetAccount(ctx, idxOwner1)
//		ownerAccNum := idxOwnerAcc1.GetAccountNumber()
//		ownerAccSeq := idxOwnerAcc1.GetSequence()
//
//		SignCheckDeliver(t, mApp.Cdc, mApp.BaseApp, header, []sdk.Msg{volumeReportMsg}, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, idxOwnerPrivKey1)
//
//		/********************* commit & check result *********************/
//		header = abci.Header{Height: mApp.LastBlockHeight() + 1}
//		ctx = mApp.BaseApp.NewContext(true, header)
//
//		//TODO: recovery when shift to main net
//		checkResult(t, ctx, k, registerKeeper,
//			volumeReportMsg.Epoch,
//			lastFoundationAccBalance,
//			lastUnissuedPrepay,
//			lastFeePool,
//			lastMatureTotalOfResNode1,
//			slashingAmtSetup,
//		) // Main net
//
//		//TODO: remove when shift to main net
//		//checkResultForIncentiveTestnet(
//		//	t, ctx, k,
//		//	volumeReportMsg.Epoch,
//		//	lastFoundationAccBalance,
//		//	lastUnissuedPrepay,
//		//	lastFeePool,
//		//	usedRewardFromMiningPool,
//		//	slashingAmtSetup,
//		//) //Incentive test net
//
//		i++
//	}
//}
//
////for incentive test net
////func checkResultForIncentiveTestnet(t *testing.T, ctx sdk.Context, k Keeper,
////	currentEpoch sdk.Int,
////	lastFoundationAccBalance sdk.Coins,
////	lastUnissuedPrepay sdk.Coin,
////	lastFeePool sdk.Coins,
////	validatorDirectDeposited sdk.Coin,
////	slashingAmtSetup sdk.Int) {
////
////	currentSlashing := k.RegisterKeeper.GetSlashing(ctx, resNodeAddr2)
////	println("currentSlashing=" + currentSlashing.String())
////
////	individualRewardTotal := sdk.Coins{}
////	newMatureEpoch := currentEpoch.Add(sdk.NewInt(k.MatureEpoch(ctx)))
////	rewardAddrList := k.GetRewardAddressPool(ctx)
////	for _, addr := range rewardAddrList {
////		individualReward, found := k.GetIndividualReward(ctx, addr, newMatureEpoch)
////		if found {
////			individualRewardTotal = individualRewardTotal.Add(individualReward.RewardFromTrafficPool...).Add(individualReward.RewardFromMiningPool...)
////		}
////
////		ctx.Logger().Info("individualReward of [" + addr.String() + "] = " + individualReward.String())
////	}
////
////	feePoolAccAddr := k.SupplyKeeper.GetModuleAddress(auth.FeeCollectorName)
////	foundationAccAddr := k.SupplyKeeper.GetModuleAddress(types.FoundationAccount)
////	newFoundationAccBalance := k.BankKeeper.GetCoins(ctx, foundationAccAddr)
////	newUnissuedPrepay := sdk.NewCoins(k.RegisterKeeper.GetTotalUnissuedPrepay(ctx))
////
////	slashingChange := slashingAmtSetup.Sub(k.RegisterKeeper.GetSlashing(ctx, resOwner1))
////	ctx.Logger().Info("resource node 1 slashing change	= " + slashingChange.String())
////	matureTotal := k.GetMatureTotalReward(ctx, resOwner1)
////	immatureTotal := k.GetImmatureTotalReward(ctx, resOwner1)
////	ctx.Logger().Info("resource node 1 matureTotal		= " + matureTotal.String())
////	ctx.Logger().Info("resource node 1 immatureTotal	= " + immatureTotal.String())
////
////	rewardSrcChange := lastFoundationAccBalance.
////		Sub(newFoundationAccBalance).
////		Add(lastUnissuedPrepay).
////		Sub(newUnissuedPrepay)
////
////	ctx.Logger().Info("rewardSrcChange					= " + rewardSrcChange.String())
////
////	rewardSrcChangeSubSlashing := deductSlashingAmt(ctx, rewardSrcChange, slashingChange)
////
////	newFeePool := k.BankKeeper.GetCoins(ctx, feePoolAccAddr)
////	ctx.Logger().Info("lastFeePool	= " + lastFeePool.String())
////	ctx.Logger().Info("newFeePool	= " + newFeePool.String())
////
////	feePoolValChange := newFeePool.Sub(lastFeePool)
////	ctx.Logger().Info("reward send to validator fee pool= " + feePoolValChange.String())
////	rewardDestChange := feePoolValChange.Add(individualRewardTotal...).Add(validatorDirectDeposited)
////
////	rewardDestChange = k.RegisterKeeper.DeductSlashing(ctx, resOwner1, rewardDestChange)
////
////	ctx.Logger().Info("rewardDestChange	= " + rewardDestChange.String())
////	require.Equal(t, rewardSrcChangeSubSlashing, rewardDestChange)
////
////}
//
//// return : coins - slashing
//func deductSlashingAmt(ctx sdk.Context, coins sdk.Coins, slashing sdk.Int) sdk.Coins {
//	ret := sdk.Coins{}
//	for _, coin := range coins {
//		if coin.Amount.GTE(slashing) {
//			coin = coin.Sub(sdk.NewCoin(coin.Denom, slashing))
//			ret = ret.Add(coin)
//			slashing = sdk.ZeroInt()
//		} else {
//			slashing = slashing.Sub(coin.Amount)
//			coin = sdk.NewCoin(coin.Denom, sdk.ZeroInt())
//			ret = ret.Add(coin)
//		}
//	}
//	return ret
//}
//
////for main net
//func checkResult(t *testing.T, ctx sdk.Context, k Keeper, registerKeeper register.Keeper,
//	currentEpoch sdk.Int,
//	lastFoundationAccBalance sdk.Coins,
//	lastUnissuedPrepay sdk.Coin,
//	lastFeePool sdk.Coins,
//	lastMatureTotalOfResNode1 sdk.Coins,
//	slashingAmtSetup sdk.Int) {
//
//	currentSlashing := registerKeeper.GetSlashing(ctx, resNodeAddr2)
//	println("currentSlashing							=" + currentSlashing.String())
//
//	individualRewardTotal := sdk.Coins{}
//	newMatureEpoch := currentEpoch.Add(sdk.NewInt(k.MatureEpoch(ctx)))
//
//	k.IteratorIndividualReward(ctx, newMatureEpoch, func(walletAddress sdk.AccAddress, individualReward types.Reward) (stop bool) {
//		individualRewardTotal = individualRewardTotal.Add(individualReward.RewardFromTrafficPool...).Add(individualReward.RewardFromMiningPool...)
//		ctx.Logger().Info("individualReward of [" + walletAddress.String() + "] = " + individualReward.String())
//		return false
//	})
//
//	feePoolAccAddr := k.SupplyKeeper.GetModuleAddress(auth.FeeCollectorName)
//	foundationAccAddr := k.SupplyKeeper.GetModuleAddress(types.FoundationAccount)
//	newFoundationAccBalance := k.BankKeeper.GetCoins(ctx, foundationAccAddr)
//	newUnissuedPrepay := sdk.NewCoins(registerKeeper.GetTotalUnissuedPrepay(ctx))
//
//	slashingChange := slashingAmtSetup.Sub(registerKeeper.GetSlashing(ctx, resOwner1))
//	ctx.Logger().Info("resource node 1 slashing change	= " + slashingChange.String())
//	matureTotal := k.GetMatureTotalReward(ctx, resOwner1)
//	immatureTotal := k.GetImmatureTotalReward(ctx, resOwner1)
//	ctx.Logger().Info("resource node 1 matureTotal		= " + matureTotal.String())
//	ctx.Logger().Info("resource node 1 immatureTotal	= " + immatureTotal.String())
//
//	rewardSrcChange := lastFoundationAccBalance.
//		Sub(newFoundationAccBalance).
//		Add(lastUnissuedPrepay).
//		Sub(newUnissuedPrepay)
//	ctx.Logger().Info("rewardSrcChange					= " + rewardSrcChange.String())
//
//	// get fee pool changes
//	newFeePool := k.BankKeeper.GetCoins(ctx, feePoolAccAddr)
//	ctx.Logger().Info("lastFeePool						= " + lastFeePool.String())
//	ctx.Logger().Info("newFeePool						= " + newFeePool.String())
//
//	feePoolValChange := newFeePool.Sub(lastFeePool)
//	ctx.Logger().Info("reward send to validator fee pool= " + feePoolValChange.String())
//
//	rewardDestChange := feePoolValChange.Add(individualRewardTotal...)
//	ctx.Logger().Info("rewardDestChange					= " + rewardDestChange.String())
//
//	require.Equal(t, rewardSrcChange, rewardDestChange)
//
//	ctx.Logger().Info("************************ slashing test***********************************")
//	ctx.Logger().Info("slashing change					= " + slashingChange.String())
//
//	upcomingMaturedIndividual := sdk.Coins{}
//	individualReward, found := k.GetIndividualReward(ctx, resOwner1, currentEpoch)
//	if found {
//		tmp := individualReward.RewardFromTrafficPool.Add(individualReward.RewardFromMiningPool...)
//		upcomingMaturedIndividual = deductSlashingAmt(ctx, tmp, slashingChange)
//	}
//	ctx.Logger().Info("upcomingMaturedIndividual		= " + upcomingMaturedIndividual.String())
//
//	// get mature total changes
//	newMatureTotalOfResNode1 := k.GetMatureTotalReward(ctx, resOwner1)
//	matureTotalOfResNode1Change, _ := newMatureTotalOfResNode1.SafeSub(lastMatureTotalOfResNode1)
//
//	if upcomingMaturedIndividual == nil {
//		upcomingMaturedIndividual = sdk.Coins{}
//	}
//	if matureTotalOfResNode1Change == nil || matureTotalOfResNode1Change.IsAnyNegative() {
//		matureTotalOfResNode1Change = sdk.Coins{}
//	}
//
//	ctx.Logger().Info("matureTotalOfResNode1Change				= " + matureTotalOfResNode1Change.String())
//	require.Equal(t, matureTotalOfResNode1Change, upcomingMaturedIndividual)
//}
//
//func checkValidator(t *testing.T, mApp *mock.App, stakingKeeper staking.Keeper,
//	addr sdk.ValAddress, expFound bool) staking.Validator {
//
//	ctxCheck := mApp.BaseApp.NewContext(true, abci.Header{})
//	validator, found := stakingKeeper.GetValidator(ctxCheck, addr)
//
//	require.Equal(t, expFound, found)
//	return validator
//}
//
//func getMockApp(t *testing.T) (*mock.App, Keeper, staking.Keeper, bank.Keeper, supply.Keeper, register.Keeper) {
//	mApp := mock.NewApp()
//
//	RegisterCodec(mApp.Cdc)
//	supply.RegisterCodec(mApp.Cdc)
//	staking.RegisterCodec(mApp.Cdc)
//	register.RegisterCodec(mApp.Cdc)
//
//	keySupply := sdk.NewKVStoreKey(supply.StoreKey)
//	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
//	keyRegister := sdk.NewKVStoreKey(register.StoreKey)
//	keyPot := sdk.NewKVStoreKey(StoreKey)
//
//	feeCollector := supply.NewEmptyModuleAccount(auth.FeeCollectorName)
//	notBondedPool := supply.NewEmptyModuleAccount(staking.NotBondedPoolName, supply.Burner, supply.Staking)
//	bondPool := supply.NewEmptyModuleAccount(staking.BondedPoolName, supply.Burner, supply.Staking)
//	foundationAccount := supply.NewEmptyModuleAccount(types.FoundationAccount)
//
//	blacklistedAddrs := make(map[string]bool)
//	blacklistedAddrs[feeCollector.GetAddress().String()] = true
//	blacklistedAddrs[notBondedPool.GetAddress().String()] = true
//	blacklistedAddrs[bondPool.GetAddress().String()] = true
//	blacklistedAddrs[foundationAccount.GetAddress().String()] = true
//
//	bankKeeper := bank.NewBaseKeeper(mApp.AccountKeeper, mApp.ParamsKeeper.Subspace(bank.DefaultParamspace), blacklistedAddrs)
//	maccPerms := map[string][]string{
//		auth.FeeCollectorName:     nil,
//		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
//		staking.BondedPoolName:    {supply.Burner, supply.Staking},
//		types.FoundationAccount:   nil,
//	}
//	supplyKeeper := supply.NewKeeper(mApp.Cdc, keySupply, mApp.AccountKeeper, bankKeeper, maccPerms)
//	stakingKeeper := staking.NewKeeper(mApp.Cdc, keyStaking, supplyKeeper, mApp.ParamsKeeper.Subspace(staking.DefaultParamspace))
//	registerKeeper := register.NewKeeper(mApp.Cdc, keyRegister, mApp.ParamsKeeper.Subspace(register.DefaultParamSpace), mApp.AccountKeeper, bankKeeper)
//
//	keeper := NewKeeper(mApp.Cdc, keyPot, mApp.ParamsKeeper.Subspace(DefaultParamSpace), auth.FeeCollectorName, bankKeeper, supplyKeeper, mApp.AccountKeeper, stakingKeeper, registerKeeper)
//
//	mApp.Router().AddRoute(staking.RouterKey, staking.NewHandler(stakingKeeper))
//	mApp.Router().AddRoute(RouterKey, NewHandler(keeper))
//	mApp.SetEndBlocker(getEndBlocker(keeper))
//	mApp.SetInitChainer(getInitChainer(mApp, keeper, mApp.AccountKeeper, supplyKeeper,
//		[]supplyexported.ModuleAccountI{feeCollector, notBondedPool, bondPool}, stakingKeeper, registerKeeper))
//
//	err := mApp.CompleteSetup(keyStaking, keySupply, keyRegister, keyPot)
//	require.NoError(t, err)
//
//	return mApp, keeper, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper
//}
//
//// getInitChainer initializes the chainer of the mock app and sets the genesis
//// state. It returns an empty ResponseInitChain.
//func getInitChainer(mapp *mock.App, keeper Keeper, accountKeeper auth.AccountKeeper, supplyKeeper supply.Keeper,
//	blacklistedAddrs []supplyexported.ModuleAccountI, stakingKeeper staking.Keeper, registerKeeper register.Keeper) sdk.InitChainer {
//	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//
//		mapp.InitChainer(ctx, req)
//
//		resourceNodes := setupAllResourceNodes()
//		indexingNodes := setupAllIndexingNodes()
//
//		registerGenesis := register.NewGenesisState(
//			register.DefaultParams(),
//			resourceNodes,
//			indexingNodes,
//			initialUOzonePrice,
//			sdk.ZeroInt(),
//			make([]register.Slashing, 0),
//		)
//
//		register.InitGenesis(ctx, registerKeeper, registerGenesis)
//
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//
//		stakingGenesis := staking.NewGenesisState(staking.NewParams(staking.DefaultUnbondingTime, staking.DefaultMaxValidators, staking.DefaultMaxEntries, 0, "ustos"), nil, nil)
//
//		totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100000000000000)))
//		supplyKeeper.SetSupply(ctx, supply.NewSupply(totalSupply))
//
//		// set module accounts
//		for _, macc := range blacklistedAddrs {
//			supplyKeeper.SetModuleAccount(ctx, macc)
//		}
//
//		validators := staking.InitGenesis(ctx, stakingKeeper, accountKeeper, supplyKeeper, stakingGenesis)
//
//		//preset
//		keeper.RegisterKeeper.SetTotalUnissuedPrepay(ctx, totalUnissuedPrepay)
//
//		//pot genesis data load
//		InitGenesis(ctx, keeper, NewGenesisState(
//			types.DefaultParams(),
//			sdk.NewCoin(types.DefaultRewardDenom, sdk.ZeroInt()),
//			0,
//			make([]types.ImmatureTotal, 0),
//			make([]types.MatureTotal, 0),
//			make([]types.Reward, 0),
//		))
//
//		return abci.ResponseInitChain{
//			Validators: validators,
//		}
//	}
//
//}
//
//// getEndBlocker returns a staking endblocker.
//func getEndBlocker(keeper Keeper) sdk.EndBlocker {
//	return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
//		validatorUpdates := keeper.StakingKeeper.BlockValidatorUpdates(ctx)
//
//		return abci.ResponseEndBlock{
//			ValidatorUpdates: validatorUpdates,
//		}
//	}
//	return nil
//}
