package sds_test

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	registertypes "github.com/stratosnet/stratos-chain/x/register/types"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/stratosnet/stratos-chain/app"
	stratos "github.com/stratosnet/stratos-chain/types"
	potKeeper "github.com/stratosnet/stratos-chain/x/pot/keeper"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"
	registerKeeper "github.com/stratosnet/stratos-chain/x/register/keeper"
	sdstypes "github.com/stratosnet/stratos-chain/x/sds/types"
)

/**
Test scenarios:
1. intt chain

*/

const (
	chainID    = "testchain_1-1"
	stos2ustos = 1000000000
)

var (
	paramSpecificMinedReward = sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(160000000000)))
	paramSpecificEpoch       = sdk.NewInt(10)

	resNodeSlashingUOZAmt1 = sdk.NewInt(100000000000)

	resourceNodeVolume1 = sdk.NewInt(500000)
	resourceNodeVolume2 = sdk.NewInt(300000)
	resourceNodeVolume3 = sdk.NewInt(200000)

	depositForSendingTx, _ = sdk.NewIntFromString("100000000000000000000000000000")
	totalUnissuedPrepayVal = sdk.ZeroInt()
	totalUnissuedPrepay    = sdk.NewCoin("ustos", totalUnissuedPrepayVal)
	initialUOzonePrice     = sdk.NewDecWithPrec(10000000, 9) // 0.001 ustos -> 1 uoz

	foundationDepositorPrivKey = secp256k1.GenPrivKey()
	foundationDepositorAccAddr = sdk.AccAddress(foundationDepositorPrivKey.PubKey().Address())
	foundationDeposit          = sdk.NewCoins(sdk.NewCoin("utros", sdk.NewInt(40000000000000000)))

	resOwnerPrivKey1 = secp256k1.GenPrivKey()
	resOwnerPrivKey2 = secp256k1.GenPrivKey()
	resOwnerPrivKey3 = secp256k1.GenPrivKey()
	resOwnerPrivKey4 = secp256k1.GenPrivKey()
	resOwnerPrivKey5 = secp256k1.GenPrivKey()
	idxOwnerPrivKey1 = secp256k1.GenPrivKey()
	idxOwnerPrivKey2 = secp256k1.GenPrivKey()
	idxOwnerPrivKey3 = secp256k1.GenPrivKey()

	resOwner1 = sdk.AccAddress(resOwnerPrivKey1.PubKey().Address())
	resOwner2 = sdk.AccAddress(resOwnerPrivKey2.PubKey().Address())
	resOwner3 = sdk.AccAddress(resOwnerPrivKey3.PubKey().Address())
	resOwner4 = sdk.AccAddress(resOwnerPrivKey4.PubKey().Address())
	resOwner5 = sdk.AccAddress(resOwnerPrivKey5.PubKey().Address())
	idxOwner1 = sdk.AccAddress(idxOwnerPrivKey1.PubKey().Address())
	idxOwner2 = sdk.AccAddress(idxOwnerPrivKey2.PubKey().Address())
	idxOwner3 = sdk.AccAddress(idxOwnerPrivKey3.PubKey().Address())

	resNodeInitialStakeForMultipleNodes = sdk.NewInt(3 * stos2ustos)

	resNodePubKey1       = ed25519.GenPrivKey().PubKey()
	resNodeAddr1         = sdk.AccAddress(resNodePubKey1.Address())
	resNodeNetworkId1    = stratos.SdsAddress(resNodePubKey1.Address())
	resNodeInitialStake1 = sdk.NewInt(3 * stos2ustos)

	resNodePubKey2       = ed25519.GenPrivKey().PubKey()
	resNodeAddr2         = sdk.AccAddress(resNodePubKey2.Address())
	resNodeNetworkId2    = stratos.SdsAddress(resNodePubKey2.Address())
	resNodeInitialStake2 = sdk.NewInt(3 * stos2ustos)

	resNodePubKey3       = ed25519.GenPrivKey().PubKey()
	resNodeAddr3         = sdk.AccAddress(resNodePubKey3.Address())
	resNodeNetworkId3    = stratos.SdsAddress(resNodePubKey3.Address())
	resNodeInitialStake3 = sdk.NewInt(3 * stos2ustos)

	resNodePubKey4       = ed25519.GenPrivKey().PubKey()
	resNodeAddr4         = sdk.AccAddress(resNodePubKey4.Address())
	resNodeNetworkId4    = stratos.SdsAddress(resNodePubKey4.Address())
	resNodeInitialStake4 = sdk.NewInt(3 * stos2ustos)

	resNodePubKey5       = ed25519.GenPrivKey().PubKey()
	resNodeAddr5         = sdk.AccAddress(resNodePubKey5.Address())
	resNodeNetworkId5    = stratos.SdsAddress(resNodePubKey5.Address())
	resNodeInitialStake5 = sdk.NewInt(3 * stos2ustos)

	idxNodePrivKey1      = ed25519.GenPrivKey()
	idxNodePubKey1       = idxNodePrivKey1.PubKey()
	idxNodeAddr1         = sdk.AccAddress(idxNodePubKey1.Address())
	idxNodeNetworkId1    = stratos.SdsAddress(idxNodePubKey1.Address())
	idxNodeInitialStake1 = sdk.NewInt(50 * stos2ustos)

	idxNodePubKey2       = ed25519.GenPrivKey().PubKey()
	idxNodeAddr2         = sdk.AccAddress(idxNodePubKey2.Address())
	idxNodeNetworkId2    = stratos.SdsAddress(idxNodePubKey2.Address())
	idxNodeInitialStake2 = sdk.NewInt(50 * stos2ustos)

	idxNodePubKey3       = ed25519.GenPrivKey().PubKey()
	idxNodeAddr3         = sdk.AccAddress(idxNodePubKey3.Address())
	idxNodeNetworkId3    = stratos.SdsAddress(idxNodePubKey3.Address())
	idxNodeInitialStake3 = sdk.NewInt(50 * stos2ustos)

	valOpPrivKey1 = secp256k1.GenPrivKey()
	valOpPubKey1  = valOpPrivKey1.PubKey()
	valOpValAddr1 = sdk.ValAddress(valOpPubKey1.Address())
	valOpAccAddr1 = sdk.AccAddress(valOpPubKey1.Address())

	valConsPrivKey1 = ed25519.GenPrivKey()
	valConsPubk1    = valConsPrivKey1.PubKey()
	valInitialStake = sdk.NewInt(15 * stos2ustos)
)

type UozPriceFactors struct {
	UOzonePrice        sdk.Dec
	InitialTotalStakes sdk.Int
	//EffectiveTotalStakes sdk.Int
	TotalUnissuedPrepay sdk.Int
	StakeAndPrepay      sdk.Int
	OzoneLimit          sdk.Int
	UozSupply           sdk.Int
}

func TestPriceCurve(t *testing.T) {

	NUM_TESTS := 100

	initFactorsBefore := &UozPriceFactors{
		UOzonePrice:        uozPrice,
		InitialTotalStakes: initialTotalStakesStore,
		//EffectiveTotalStakes: registerKeeper.GetEffectiveGenesisStakeTotal(ctx),
		TotalUnissuedPrepay: totalUnissuedPrepayStore,
		StakeAndPrepay:      initialTotalStakesStore.Add(totalUnissuedPrepayStore),
		OzoneLimit:          initialTotalStakesStore.ToDec().Quo(uozPrice).TruncateInt(),
	}

	initFactorsBefore, _, _ = simulatePriceChange(&PriceChangeEvent{
		stakeDelta:          sdk.ZeroInt(),
		unissuedPrepayDelta: sdk.ZeroInt(),
	}, initFactorsBefore)

	stakeChangePerm := rand.Perm(NUM_TESTS)
	prepayChangePerm := rand.Perm(NUM_TESTS)

	for i := 0; i < NUM_TESTS; i++ {
		tempStakeSign := 1
		if rand.Intn(5) >= 3 {
			tempStakeSign = -1
		}
		tempPrepaySign := rand.Intn(5)
		if rand.Intn(5) >= 3 {
			tempPrepaySign = -1
		}
		change := &PriceChangeEvent{
			stakeDelta:          sdk.NewInt(int64(stakeChangePerm[i]) * stos2ustos).Mul(sdk.NewInt(int64(tempStakeSign))),
			unissuedPrepayDelta: sdk.NewInt(int64(prepayChangePerm[i]) * stos2ustos).Mul(sdk.NewInt(int64(tempPrepaySign))),
		}
		fmt.Printf("\nstakeDelta: %v, unissuedPrepayDelta: %v\n", change.stakeDelta.String(), change.unissuedPrepayDelta.String())
		initFactorsBefore, _, _ = simulatePriceChange(change, initFactorsBefore)
	}
}

func TestOzPriceChange(t *testing.T) {
	/********************* initialize mock app *********************/
	//mApp, k, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper := getMockApp(t)
	accs, balances := setupAccounts()
	//stApp := app.SetupWithGenesisAccounts(accs, chainID, balances...)
	validators := make([]*tmtypes.Validator, 0)
	valSet := tmtypes.NewValidatorSet(validators)
	metaNodes := setupAllMetaNodes()
	//resourceNodes := setupAllResourceNodes()
	resourceNodes := make([]registertypes.ResourceNode, 0)

	stApp := app.SetupWithGenesisNodeSet(t, true, valSet, metaNodes, resourceNodes, accs, totalUnissuedPrepay, chainID, balances...)

	accountKeeper := stApp.GetAccountKeeper()
	bankKeeper := stApp.GetBankKeeper()
	registerKeeper := stApp.GetRegisterKeeper()
	potKeeper := stApp.GetPotKeeper()

	/********************* foundation account deposit *********************/
	header := tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := stApp.BaseApp.NewContext(true, header)

	foundationDepositMsg := pottypes.NewMsgFoundationDeposit(foundationDeposit, foundationDepositorAccAddr)
	txGen := app.MakeTestEncodingConfig().TxConfig

	foundationDepositorAcc := accountKeeper.GetAccount(ctx, foundationDepositorAccAddr)
	accNum := foundationDepositorAcc.GetAccountNumber()
	accSeq := foundationDepositorAcc.GetSequence()
	_, _, err := app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{foundationDepositMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, foundationDepositorPrivKey)
	require.NoError(t, err)
	foundationAccountAddr := accountKeeper.GetModuleAddress(pottypes.FoundationAccount)
	app.CheckBalance(t, stApp, foundationAccountAddr, foundationDeposit)

	/********************* create validator with 50% commission *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	description := stakingtypes.NewDescription("foo_moniker", chainID, "", "", "")
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(valOpValAddr1, valConsPubk1, sdk.NewCoin("ustos", valInitialStake), description, commission, sdk.OneInt())

	valOpAcc1 := accountKeeper.GetAccount(ctx, valOpAccAddr1)
	accNum = valOpAcc1.GetAccountNumber()
	accSeq = valOpAcc1.GetSequence()
	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createValidatorMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, valOpPrivKey1)
	require.NoError(t, err)
	app.CheckBalance(t, stApp, valOpAccAddr1, nil)

	/********************** commit **********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	validator := checkValidator(t, stApp, valOpValAddr1, true)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, valInitialStake, validator.BondedTokens()))

	uozPriceFactorsSeq0, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, UozPriceFactors{
		UOzonePrice:        registerKeeper.CurrUozPrice(ctx),
		InitialTotalStakes: registerKeeper.GetInitialGenesisStakeTotal(ctx),
		//EffectiveTotalStakes: registerKeeper.GetEffectiveGenesisStakeTotal(ctx),
		TotalUnissuedPrepay: registerKeeper.GetTotalUnissuedPrepay(ctx).Amount,
		StakeAndPrepay:      registerKeeper.GetInitialGenesisStakeTotal(ctx).Add(registerKeeper.GetTotalUnissuedPrepay(ctx).Amount),
		OzoneLimit:          registerKeeper.GetRemainingOzoneLimit(ctx),
	})

	// start testing
	println("\n********************************* Deliver Prepay Tx START ********************************************")
	prepayMsg := setupPrepayMsg()
	/********************* deliver tx *********************/

	resOwnerAcc := accountKeeper.GetAccount(ctx, resOwner1)
	ownerAccNum := resOwnerAcc.GetAccountNumber()
	ownerAccSeq := resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{prepayMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey1)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	uozPriceFactorsSeq1, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq0)
	require.True(t, uozPricePercentage.GT(sdk.ZeroDec()), "Uoz price should increase after PREPAY")
	require.True(t, ozoneLimitPercentage.LT(sdk.ZeroDec()), "OzLimit should decrease after PREPAY")
	println("********************************* Deliver Prepay Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver CreateResourceNode Tx START ********************************************")
	createResourceNodeMsg := setupMsgCreateResourceNode1()
	/********************* deliver tx *********************/

	resOwnerAcc = accountKeeper.GetAccount(ctx, resOwner1)
	ownerAccNum = resOwnerAcc.GetAccountNumber()
	ownerAccSeq = resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey1)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	uozPriceFactorsSeq2, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq1)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
	require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
	println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver UnsuspendResourceNode Tx (Slashing) START ********************************************")
	unsuspendMsg := setupUnsuspendMsg()
	/********************* deliver tx *********************/

	idxOwnerAcc := accountKeeper.GetAccount(ctx, idxOwner1)
	ownerAccNum = idxOwnerAcc.GetAccountNumber()
	ownerAccSeq = idxOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{unsuspendMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, idxOwnerPrivKey1)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	slashingAmtSetup := registerKeeper.GetSlashing(ctx, resOwner1)

	totalConsumedUoz := sdk.ZeroInt().ToDec()

	slashingAmtCheck := potKeeper.GetTrafficReward(ctx, totalConsumedUoz)
	println("slashingAmtSetup=" + slashingAmtSetup.String())
	require.Equal(t, slashingAmtSetup, slashingAmtCheck.TruncateInt())

	uozPriceFactorsSeq3, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq2)
	require.True(t, uozPricePercentage.Equal(sdk.ZeroDec()), "Uoz price shouldn't change after UnsuspendResourceNode")
	require.True(t, ozoneLimitPercentage.Equal(sdk.ZeroDec()), "OzLimit shouldn't change after UnsuspendResourceNode")
	println("********************************* Deliver UnsuspendResourceNode Tx (Slashing) END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver Slashing Tx (not suspending/unsuspending) START ********************************************")
	slashingMsg := setupSlashingMsg()
	/********************* deliver tx *********************/

	idxOwnerAcc = accountKeeper.GetAccount(ctx, idxOwner1)
	ownerAccNum = idxOwnerAcc.GetAccountNumber()
	ownerAccSeq = idxOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{slashingMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, idxOwnerPrivKey1)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	slashingAmtSetup = registerKeeper.GetSlashing(ctx, resOwner1)

	totalConsumedUoz = resNodeSlashingUOZAmt1.ToDec()

	slashingAmtCheck = potKeeper.GetTrafficReward(ctx, totalConsumedUoz)
	println("slashingAmtSetup=" + slashingAmtSetup.String())
	require.Equal(t, slashingAmtSetup, slashingAmtCheck.TruncateInt())

	uozPriceFactorsSeq4, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq3)
	require.True(t, uozPricePercentage.Equal(sdk.ZeroDec()), "Uoz price shouldn't change after SlashResourceNode")
	require.True(t, ozoneLimitPercentage.Equal(sdk.ZeroDec()), "OzLimit shouldn't change after SlashResourceNode")
	println("********************************* Deliver Slashing Tx (not suspending/unsuspending) END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver VolumeReport Tx START ********************************************")
	/********************* prepare tx data *********************/
	volumeReportMsg := setupMsgVolumeReport(1)

	lastTotalMinedToken := potKeeper.GetTotalMinedTokens(ctx)
	println("last committed TotalMinedTokens = " + lastTotalMinedToken.String())
	epoch, ok := sdk.NewIntFromString(volumeReportMsg.Epoch.String())
	require.Equal(t, ok, true)

	totalConsumedUoz = potKeeper.GetTotalConsumedUoz(volumeReportMsg.WalletVolumes).ToDec()

	/********************* print info *********************/
	println("epoch " + volumeReportMsg.Epoch.String())
	S := registerKeeper.GetInitialGenesisStakeTotal(ctx).ToDec()
	Pt := registerKeeper.GetTotalUnissuedPrepay(ctx).Amount.ToDec()
	Y := totalConsumedUoz
	Lt := registerKeeper.GetRemainingOzoneLimit(ctx).ToDec()
	R := S.Add(Pt).Mul(Y).Quo(Lt.Add(Y))
	//println("R = (S + Pt) * Y / (Lt + Y)")
	println("S=" + S.String() + "\nPt=" + Pt.String() + "\nY=" + Y.String() + "\nLt=" + Lt.String() + "\nR=" + R.String() + "\n")

	println("---------------------------")
	distributeGoal := pottypes.InitDistributeGoal()
	distributeGoal, err = potKeeper.CalcTrafficRewardInTotal(ctx, distributeGoal, totalConsumedUoz)
	require.NoError(t, err)

	distributeGoal, err = potKeeper.CalcMiningRewardInTotal(ctx, distributeGoal) //for main net
	require.NoError(t, err)
	println(distributeGoal.String())

	println("---------------------------")
	println("distribute detail:")
	distributeGoalBalance := distributeGoal
	rewardDetailMap := make(map[string]pottypes.Reward)
	rewardDetailMap = potKeeper.CalcRewardForResourceNode(ctx, totalConsumedUoz, volumeReportMsg.WalletVolumes, distributeGoalBalance, rewardDetailMap)
	rewardDetailMap = potKeeper.CalcRewardForMetaNode(ctx, distributeGoalBalance, rewardDetailMap)

	println("resource_wallet1:  address = " + resOwner1.String())
	println("              miningReward = " + rewardDetailMap[resOwner1.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[resOwner1.String()].RewardFromTrafficPool.String())

	println("resource_wallet2:  address = " + resOwner2.String())
	println("              miningReward = " + rewardDetailMap[resOwner2.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[resOwner2.String()].RewardFromTrafficPool.String())

	println("resource_wallet3:  address = " + resOwner3.String())
	println("              miningReward = " + rewardDetailMap[resOwner3.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[resOwner3.String()].RewardFromTrafficPool.String())

	println("resource_wallet4:  address = " + resOwner4.String())
	println("              miningReward = " + rewardDetailMap[resOwner4.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[resOwner4.String()].RewardFromTrafficPool.String())

	println("resource_wallet5:  address = " + resOwner5.String())
	println("              miningReward = " + rewardDetailMap[resOwner5.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[resOwner5.String()].RewardFromTrafficPool.String())

	println("indexing_wallet1:  address = " + idxOwner1.String())
	println("              miningReward = " + rewardDetailMap[idxOwner1.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[idxOwner1.String()].RewardFromTrafficPool.String())

	println("indexing_wallet2:  address = " + idxOwner2.String())
	println("              miningReward = " + rewardDetailMap[idxOwner2.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[idxOwner2.String()].RewardFromTrafficPool.String())

	println("indexing_wallet3:  address = " + idxOwner3.String())
	println("              miningReward = " + rewardDetailMap[idxOwner3.String()].RewardFromMiningPool.String())
	println("             trafficReward = " + rewardDetailMap[idxOwner3.String()].RewardFromTrafficPool.String())
	println("---------------------------")

	/********************* record data before delivering tx  *********************/
	lastFoundationAccBalance := bankKeeper.GetAllBalances(ctx, foundationAccountAddr)
	lastUnissuedPrepay := registerKeeper.GetTotalUnissuedPrepay(ctx)
	lastMatureTotalOfResNode1 := potKeeper.GetMatureTotalReward(ctx, resOwner1)

	/********************* deliver tx *********************/
	idxOwnerAcc = accountKeeper.GetAccount(ctx, idxOwner1)
	ownerAccNum = idxOwnerAcc.GetAccountNumber()
	ownerAccSeq = idxOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{volumeReportMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, idxOwnerPrivKey1)
	require.NoError(t, err)

	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	epoch, ok = sdk.NewIntFromString(volumeReportMsg.Epoch.String())
	require.Equal(t, ok, true)

	checkResult(t, ctx, potKeeper,
		accountKeeper,
		bankKeeper,
		registerKeeper,
		epoch,
		lastFoundationAccBalance,
		lastUnissuedPrepay,
		lastMatureTotalOfResNode1,
		slashingAmtSetup,
	)

	uozPriceFactorsSeq5, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq4)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should increase after VolumeReport")
	require.True(t, ozoneLimitPercentage.Equal(sdk.ZeroDec()), "OzLimit shouldn't change after VolumeReport")
	println("********************************* Deliver VolumeReport Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver CreateResourceNode Tx START ********************************************")
	createResourceNodeMsg2 := setupMsgCreateResourceNode2()
	/********************* deliver tx *********************/

	resOwnerAcc = accountKeeper.GetAccount(ctx, resOwner2)
	ownerAccNum = resOwnerAcc.GetAccountNumber()
	ownerAccSeq = resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg2}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey2)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	uozPriceFactorsSeq6, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq5)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
	require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
	println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver CreateResourceNode Tx START ********************************************")
	createResourceNodeMsg3 := setupMsgCreateResourceNode3()
	/********************* deliver tx *********************/

	resOwnerAcc = accountKeeper.GetAccount(ctx, resOwner3)
	ownerAccNum = resOwnerAcc.GetAccountNumber()
	ownerAccSeq = resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg3}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey3)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	uozPriceFactorsSeq7, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq6)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
	require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
	println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver CreateResourceNode Tx START ********************************************")
	createResourceNodeMsg4 := setupMsgCreateResourceNode4()
	/********************* deliver tx *********************/

	resOwnerAcc = accountKeeper.GetAccount(ctx, resOwner4)
	ownerAccNum = resOwnerAcc.GetAccountNumber()
	ownerAccSeq = resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg4}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey4)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	uozPriceFactorsSeq8, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq7)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
	require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
	println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

	println("********************************* Deliver CreateResourceNode Tx START ********************************************")
	createResourceNodeMsg5 := setupMsgCreateResourceNode5()
	/********************* deliver tx *********************/

	resOwnerAcc = accountKeeper.GetAccount(ctx, resOwner5)
	ownerAccNum = resOwnerAcc.GetAccountNumber()
	ownerAccSeq = resOwnerAcc.GetSequence()

	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg5}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey5)
	require.NoError(t, err)
	/********************* commit & check result *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	_, uozPricePercentage, ozoneLimitPercentage = printCurrUozPrice(ctx, registerKeeper, uozPriceFactorsSeq8)
	require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
	require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
	println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

}

// initialize data of volume report
func setupMsgVolumeReport(newEpoch int64) *pottypes.MsgVolumeReport {
	volume1 := pottypes.NewSingleWalletVolume(resOwner1, resourceNodeVolume1)
	volume2 := pottypes.NewSingleWalletVolume(resOwner2, resourceNodeVolume2)
	volume3 := pottypes.NewSingleWalletVolume(resOwner3, resourceNodeVolume3)

	nodesVolume := []*pottypes.SingleWalletVolume{volume1, volume2, volume3}
	reporter := idxNodeNetworkId1
	epoch := sdk.NewInt(newEpoch)
	reportReference := "report for epoch " + epoch.String()
	reporterOwner := idxOwner1

	pubKeys := make([][]byte, 1)
	for i := range pubKeys {
		pubKeys[i] = make([]byte, 1)
	}

	signature := pottypes.NewBLSSignatureInfo(pubKeys, []byte("signature"), []byte("txData"))

	volumeReportMsg := pottypes.NewMsgVolumeReport(nodesVolume, reporter, epoch, reportReference, reporterOwner, signature)

	return volumeReportMsg
}

func setupSlashingMsg() *pottypes.MsgSlashingResourceNode {
	reporters := make([]stratos.SdsAddress, 0)
	reporters = append(reporters, idxNodeNetworkId1)
	reportOwner := make([]sdk.AccAddress, 0)
	reportOwner = append(reportOwner, idxOwner1)
	slashingMsg := pottypes.NewMsgSlashingResourceNode(reporters, reportOwner, resNodeNetworkId1, resOwner1, resNodeSlashingUOZAmt1, true)
	return slashingMsg
}

func setupUnsuspendMsg() *pottypes.MsgSlashingResourceNode {
	reporters := make([]stratos.SdsAddress, 0)
	reporters = append(reporters, idxNodeNetworkId1)
	reportOwner := make([]sdk.AccAddress, 0)
	reportOwner = append(reportOwner, idxOwner1)
	slashingMsg := pottypes.NewMsgSlashingResourceNode(reporters, reportOwner, resNodeNetworkId1, resOwner1, sdk.ZeroInt(), false)
	return slashingMsg
}
func setupPrepayMsg() *sdstypes.MsgPrepay {
	sender := resOwner1
	prepayMsg := sdstypes.NewMsgPrepay(sender.String(), sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(75000000000))))
	return prepayMsg
}

func setupMsgRemoveResourceNode(i int, resNodeNetworkId stratos.SdsAddress, resOwner sdk.AccAddress) *registertypes.MsgRemoveResourceNode {
	removeResourceNodeMsg := registertypes.NewMsgRemoveResourceNode(resNodeNetworkId, resOwner)
	return removeResourceNodeMsg
}
func setupMsgCreateResourceNode(i int, resNodeNetworkId stratos.SdsAddress, resNodePubKey cryptotypes.PubKey, resOwner sdk.AccAddress) *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId, resNodePubKey, sdk.NewCoin(stratos.USTOS, resNodeInitialStakeForMultipleNodes), resOwner, registertypes.NewDescription("sds://resourceNode"+strconv.Itoa(i+1), "", "", "", ""), nodeType)
	return createResourceNodeMsg
}

func setupMsgCreateResourceNode1() *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId1, resNodePubKey1, sdk.NewCoin(stratos.USTOS, resNodeInitialStake1), resOwner1, registertypes.NewDescription("sds://resourceNode1", "", "", "", ""), nodeType)
	return createResourceNodeMsg
}
func setupMsgCreateResourceNode2() *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId2, resNodePubKey2, sdk.NewCoin(stratos.USTOS, resNodeInitialStake2), resOwner2, registertypes.NewDescription("sds://resourceNode2", "", "", "", ""), nodeType)
	return createResourceNodeMsg
}
func setupMsgCreateResourceNode3() *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId3, resNodePubKey3, sdk.NewCoin(stratos.USTOS, resNodeInitialStake3), resOwner3, registertypes.NewDescription("sds://resourceNode3", "", "", "", ""), nodeType)
	return createResourceNodeMsg
}

func setupMsgCreateResourceNode4() *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId4, resNodePubKey4, sdk.NewCoin(stratos.USTOS, resNodeInitialStake4), resOwner4, registertypes.NewDescription("sds://resourceNode4", "", "", "", ""), nodeType)
	return createResourceNodeMsg
}

func setupMsgCreateResourceNode5() *registertypes.MsgCreateResourceNode {
	nodeType := uint32(registertypes.STORAGE)
	createResourceNodeMsg, _ := registertypes.NewMsgCreateResourceNode(resNodeNetworkId5, resNodePubKey5, sdk.NewCoin(stratos.USTOS, resNodeInitialStake5), resOwner5, registertypes.NewDescription("sds://resourceNode5", "", "", "", ""), nodeType)
	return createResourceNodeMsg
}

func printCurrUozPrice(ctx sdk.Context, registerKeeper registerKeeper.Keeper, uozPriceFactorsBefore UozPriceFactors) (UozPriceFactors, sdk.Dec, sdk.Dec) {
	uozPriceFactorsAfter := UozPriceFactors{}
	uozPriceFactorsAfter.InitialTotalStakes = registerKeeper.GetInitialGenesisStakeTotal(ctx)
	//uozPriceFactorsAfter.EffectiveTotalStakes = registerKeeper.GetEffectiveGenesisStakeTotal(ctx)
	uozPriceFactorsAfter.TotalUnissuedPrepay = registerKeeper.GetTotalUnissuedPrepay(ctx).Amount
	uozPriceFactorsAfter.StakeAndPrepay = uozPriceFactorsAfter.InitialTotalStakes.Add(uozPriceFactorsAfter.TotalUnissuedPrepay)
	uozPriceFactorsAfter.OzoneLimit = registerKeeper.GetRemainingOzoneLimit(ctx)
	uozPriceFactorsAfter.UOzonePrice = registerKeeper.CurrUozPrice(ctx)
	_, uozPriceFactorsAfter.UozSupply = registerKeeper.UozSupply(ctx)

	uozPriceDelta := uozPriceFactorsAfter.UOzonePrice.Sub(uozPriceFactorsBefore.UOzonePrice)
	initialTotalStakesDelta := uozPriceFactorsAfter.InitialTotalStakes.Sub(uozPriceFactorsBefore.InitialTotalStakes)
	//effectiveTotalStakesDelta := uozPriceFactorsAfter.EffectiveTotalStakes.Sub(uozPriceFactorsBefore.EffectiveTotalStakes)
	totalUnissuedPrepayDelta := uozPriceFactorsAfter.TotalUnissuedPrepay.Sub(uozPriceFactorsBefore.TotalUnissuedPrepay)
	stakeAndPrepayDelta := uozPriceFactorsAfter.StakeAndPrepay.Sub(uozPriceFactorsBefore.StakeAndPrepay)
	ozoneLimitDelta := uozPriceFactorsAfter.OzoneLimit.Sub(uozPriceFactorsBefore.OzoneLimit)
	uozSupplyDelta := uozPriceFactorsAfter.UozSupply.Sub(uozPriceFactorsBefore.UozSupply)

	uozPricePercentage := uozPriceDelta.Quo(uozPriceFactorsBefore.UOzonePrice).MulInt(sdk.NewInt(100))
	//initialTotalStakesPercentage := initialTotalStakesDelta.Quo(uozPriceFactorsBefore.InitialTotalStakes)
	//effectiveTotalStakesPercentage := effectiveTotalStakesDelta.Quo(uozPriceFactorsBefore.EffectiveTotalStakes)
	//totalUnissuedPrepayPercentage := totalUnissuedPrepayDelta.Quo(uozPriceFactorsBefore.TotalUnissuedPrepay)
	//stakeAndPrepayPercentage := stakeAndPrepayDelta.Quo(uozPriceFactorsBefore.StakeAndPrepay)
	ozoneLimitPercentage := ozoneLimitDelta.ToDec().Quo(uozPriceFactorsBefore.OzoneLimit.ToDec()).MulInt(sdk.NewInt(100))

	println("===>>>>>>>>>>>>>>     Current Uoz Price    ===>>>>>>>>>>>>>>")
	println("UOzonePrice: 									" + uozPriceFactorsAfter.UOzonePrice.String() + "(delta: " + uozPriceDelta.String() + ", " + uozPricePercentage.String()[:5] + "%)")
	println("InitialTotalStakes: 							" + uozPriceFactorsAfter.InitialTotalStakes.String() + "(delta: " + initialTotalStakesDelta.String() + ")")
	//println("EffectiveTotalStakes: 							" + uozPriceFactorsAfter.EffectiveTotalStakes.String() + "(delta: " + effectiveTotalStakesDelta.String() + ")")
	println("TotalUnissuedPrepay: 							" + uozPriceFactorsAfter.TotalUnissuedPrepay.String() + "(delta: " + totalUnissuedPrepayDelta.String() + ")")
	println("InitialTotalStakes+TotalUnissuedPrepay:			" + uozPriceFactorsAfter.StakeAndPrepay.String() + "(delta: " + stakeAndPrepayDelta.String() + ")")
	println("OzoneLimit: 									" + uozPriceFactorsAfter.OzoneLimit.String() + "(delta: " + ozoneLimitDelta.String() + ", " + ozoneLimitPercentage.String()[:5] + "%)")
	println("UozSupply: 									    " + uozPriceFactorsAfter.UozSupply.String() + "(delta: " + uozSupplyDelta.String() + ")")

	return uozPriceFactorsAfter, uozPricePercentage, ozoneLimitPercentage
}

// return : coins - slashing
func deductSlashingAmt(ctx sdk.Context, coins sdk.Coins, slashing sdk.Int) sdk.Coins {
	ret := sdk.Coins{}
	for _, coin := range coins {
		if coin.Amount.GTE(slashing) {
			coin = coin.Sub(sdk.NewCoin(coin.Denom, slashing))
			ret = ret.Add(coin)
			slashing = sdk.ZeroInt()
		} else {
			slashing = slashing.Sub(coin.Amount)
			coin = sdk.NewCoin(coin.Denom, sdk.ZeroInt())
			ret = ret.Add(coin)
		}
	}
	return ret
}

//for main net
func checkResult(t *testing.T, ctx sdk.Context,
	k potKeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankKeeper.Keeper,
	registerKeeper registerKeeper.Keeper,
	currentEpoch sdk.Int,
	lastFoundationAccBalance sdk.Coins,
	lastUnissuedPrepay sdk.Coin,
	lastMatureTotalOfResNode1 sdk.Coins,
	slashingAmtSetup sdk.Int) {

	currentSlashing := registerKeeper.GetSlashing(ctx, resNodeAddr2)
	println("currentSlashing					= " + currentSlashing.String())

	individualRewardTotal := sdk.Coins{}
	newMatureEpoch := currentEpoch.Add(sdk.NewInt(k.MatureEpoch(ctx)))

	k.IteratorIndividualReward(ctx, newMatureEpoch, func(walletAddress sdk.AccAddress, individualReward pottypes.Reward) (stop bool) {
		individualRewardTotal = individualRewardTotal.Add(individualReward.RewardFromTrafficPool...).Add(individualReward.RewardFromMiningPool...)
		println("individualReward of [" + walletAddress.String() + "] = " + individualReward.String())
		return false
	})

	feePoolAccAddr := accountKeeper.GetModuleAddress(authtypes.FeeCollectorName)
	require.NotNil(t, feePoolAccAddr)
	foundationAccountAddr := accountKeeper.GetModuleAddress(pottypes.FoundationAccount)
	newFoundationAccBalance := bankKeeper.GetAllBalances(ctx, foundationAccountAddr)
	newUnissuedPrepay := sdk.NewCoins(registerKeeper.GetTotalUnissuedPrepay(ctx))

	slashingChange := slashingAmtSetup.Sub(registerKeeper.GetSlashing(ctx, resOwner1))
	println("resource node 1 slashing change		= " + slashingChange.String())
	matureTotal := k.GetMatureTotalReward(ctx, resOwner1)
	immatureTotal := k.GetImmatureTotalReward(ctx, resOwner1)
	println("resource node 1 matureTotal		= " + matureTotal.String())
	println("resource node 1 immatureTotal		= " + immatureTotal.String())

	rewardSrcChange := lastFoundationAccBalance.
		Sub(newFoundationAccBalance).
		Add(lastUnissuedPrepay).
		Sub(newUnissuedPrepay)
	println("rewardSrcChange				= " + rewardSrcChange.String())

	// distribution module will send all tokens from "fee_collector" to "distribution" account in the BeginBlocker() method
	feePoolValChange := bankKeeper.GetAllBalances(ctx, feePoolAccAddr)
	println("reward send to validator fee pool	= " + feePoolValChange.String())

	rewardDestChange := feePoolValChange.Add(individualRewardTotal...)
	println("rewardDestChange			= " + rewardDestChange.String())

	require.Equal(t, rewardSrcChange, rewardDestChange)

	println("************************ slashing test***********************************")
	println("slashing change				= " + slashingChange.String())

	upcomingMaturedIndividual := sdk.Coins{}
	individualReward, found := k.GetIndividualReward(ctx, resOwner1, currentEpoch)
	if found {
		tmp := individualReward.RewardFromTrafficPool.Add(individualReward.RewardFromMiningPool...)
		upcomingMaturedIndividual = deductSlashingAmt(ctx, tmp, slashingChange)
	}
	println("upcomingMaturedIndividual		= " + upcomingMaturedIndividual.String())

	// get mature total changes
	newMatureTotalOfResNode1 := k.GetMatureTotalReward(ctx, resOwner1)
	matureTotalOfResNode1Change, _ := newMatureTotalOfResNode1.SafeSub(lastMatureTotalOfResNode1)
	if matureTotalOfResNode1Change == nil || matureTotalOfResNode1Change.IsAnyNegative() {
		matureTotalOfResNode1Change = sdk.Coins{}
	}
	println("matureTotalOfResNode1Change		= " + matureTotalOfResNode1Change.String())
	require.Equal(t, matureTotalOfResNode1Change.String(), upcomingMaturedIndividual.String())
}

func checkValidator(t *testing.T, app *app.NewApp, addr sdk.ValAddress, expFound bool) stakingtypes.Validator {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	validator, found := app.GetStakingKeeper().GetValidator(ctxCheck, addr)

	require.Equal(t, expFound, found)
	return validator
}

func TestMain(m *testing.M) {
	config := stratos.GetConfig()
	config.Seal()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func setupAccounts() ([]authtypes.GenesisAccount, []banktypes.Balance) {

	//************************** setup resource nodes owners' accounts **************************
	resOwnerAcc1 := &authtypes.BaseAccount{Address: resOwner1.String()}
	resOwnerAcc2 := &authtypes.BaseAccount{Address: resOwner2.String()}
	resOwnerAcc3 := &authtypes.BaseAccount{Address: resOwner3.String()}
	resOwnerAcc4 := &authtypes.BaseAccount{Address: resOwner4.String()}
	resOwnerAcc5 := &authtypes.BaseAccount{Address: resOwner5.String()}
	//************************** setup indexing nodes owners' accounts **************************
	idxOwnerAcc1 := &authtypes.BaseAccount{Address: idxOwner1.String()}
	idxOwnerAcc2 := &authtypes.BaseAccount{Address: idxOwner2.String()}
	idxOwnerAcc3 := &authtypes.BaseAccount{Address: idxOwner3.String()}
	//************************** setup validator delegators' accounts **************************
	valOwnerAcc1 := &authtypes.BaseAccount{Address: valOpAccAddr1.String()}
	//************************** setup indexing nodes' accounts **************************
	idxNodeAcc1 := &authtypes.BaseAccount{Address: idxNodeAddr1.String()}
	foundationDepositorAcc := &authtypes.BaseAccount{Address: foundationDepositorAccAddr.String()}

	accs := []authtypes.GenesisAccount{
		resOwnerAcc1, resOwnerAcc2, resOwnerAcc3, resOwnerAcc4, resOwnerAcc5,
		idxOwnerAcc1, idxOwnerAcc2, idxOwnerAcc3,
		valOwnerAcc1,
		foundationDepositorAcc,
		idxNodeAcc1,
	}

	balances := []banktypes.Balance{
		{
			Address: resOwner1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake1.Add(depositForSendingTx))},
		},
		{
			Address: resOwner2.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake2)},
		},
		{
			Address: resOwner3.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake3)},
		},
		{
			Address: resOwner4.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake4)},
		},
		{
			Address: resOwner5.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake5)},
		},
		{
			Address: idxOwner1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake1)},
		},
		{
			Address: idxOwner2.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake2)},
		},
		{
			Address: idxOwner3.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake3)},
		},
		{
			Address: valOpAccAddr1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", valInitialStake)},
		},
		{
			Address: idxNodeAddr1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", sdk.ZeroInt())},
		},
		{
			Address: foundationDepositorAccAddr.String(),
			Coins:   foundationDeposit,
		},
	}
	return accs, balances
}

func setupAccountsMultipleResNodes(resOwners []sdk.AccAddress) ([]authtypes.GenesisAccount, []banktypes.Balance) {

	resOwnerAccs := make([]*authtypes.BaseAccount, 0, len(resOwners))
	//************************** setup resource nodes owners' accounts **************************
	for _, resOwner := range resOwners {
		resOwnerAccs = append(resOwnerAccs, &authtypes.BaseAccount{Address: resOwner.String()})
	}
	//resOwnerAcc1 := &authtypes.BaseAccount{Address: resOwner1.String()}
	//resOwnerAcc2 := &authtypes.BaseAccount{Address: resOwner2.String()}
	//resOwnerAcc3 := &authtypes.BaseAccount{Address: resOwner3.String()}
	//resOwnerAcc4 := &authtypes.BaseAccount{Address: resOwner4.String()}
	//resOwnerAcc5 := &authtypes.BaseAccount{Address: resOwner5.String()}
	//************************** setup indexing nodes owners' accounts **************************
	idxOwnerAcc1 := &authtypes.BaseAccount{Address: idxOwner1.String()}
	idxOwnerAcc2 := &authtypes.BaseAccount{Address: idxOwner2.String()}
	idxOwnerAcc3 := &authtypes.BaseAccount{Address: idxOwner3.String()}
	//************************** setup validator delegators' accounts **************************
	valOwnerAcc1 := &authtypes.BaseAccount{Address: valOpAccAddr1.String()}
	//************************** setup indexing nodes' accounts **************************
	idxNodeAcc1 := &authtypes.BaseAccount{Address: idxNodeAddr1.String()}
	foundationDepositorAcc := &authtypes.BaseAccount{Address: foundationDepositorAccAddr.String()}

	accs := []authtypes.GenesisAccount{
		//resOwnerAcc1, resOwnerAcc2, resOwnerAcc3, resOwnerAcc4, resOwnerAcc5,
		idxOwnerAcc1, idxOwnerAcc2, idxOwnerAcc3,
		valOwnerAcc1,
		foundationDepositorAcc,
		idxNodeAcc1,
	}

	balances := []banktypes.Balance{
		//{
		//	Address: resOwner1.String(),
		//	Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake1.Add(depositForSendingTx))},
		//},
		//{
		//	Address: resOwner2.String(),
		//	Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake2)},
		//},
		//{
		//	Address: resOwner3.String(),
		//	Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake3)},
		//},
		//{
		//	Address: resOwner4.String(),
		//	Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake4)},
		//},
		//{
		//	Address: resOwner5.String(),
		//	Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake5)},
		//},
		{
			Address: idxOwner1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake1)},
		},
		{
			Address: idxOwner2.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake2)},
		},
		{
			Address: idxOwner3.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake3)},
		},
		{
			Address: valOpAccAddr1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", valInitialStake)},
		},
		{
			Address: idxNodeAddr1.String(),
			Coins:   sdk.Coins{sdk.NewCoin("ustos", sdk.ZeroInt())},
		},
		{
			Address: foundationDepositorAccAddr.String(),
			Coins:   foundationDeposit,
		},
	}

	for _, resOwnerAcc := range resOwnerAccs {
		accs = append(accs, resOwnerAcc)
		balances = append(balances, banktypes.Balance{
			Address: resOwnerAcc.Address,
			Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake1.Add(depositForSendingTx))},
		})
	}

	return accs, balances
}

func setupAllResourceNodes() []registertypes.ResourceNode {

	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")
	nodeType := registertypes.STORAGE
	resourceNode1, _ := registertypes.NewResourceNode(resNodeNetworkId1, resNodePubKey1, resOwner1, registertypes.NewDescription("sds://resourceNode1", "", "", "", ""), nodeType, time)
	resourceNode2, _ := registertypes.NewResourceNode(resNodeNetworkId2, resNodePubKey2, resOwner2, registertypes.NewDescription("sds://resourceNode2", "", "", "", ""), nodeType, time)
	resourceNode3, _ := registertypes.NewResourceNode(resNodeNetworkId3, resNodePubKey3, resOwner3, registertypes.NewDescription("sds://resourceNode3", "", "", "", ""), nodeType, time)
	resourceNode4, _ := registertypes.NewResourceNode(resNodeNetworkId4, resNodePubKey4, resOwner4, registertypes.NewDescription("sds://resourceNode4", "", "", "", ""), nodeType, time)
	resourceNode5, _ := registertypes.NewResourceNode(resNodeNetworkId5, resNodePubKey5, resOwner5, registertypes.NewDescription("sds://resourceNode5", "", "", "", ""), nodeType, time)

	resourceNode1 = resourceNode1.AddToken(resNodeInitialStake1)
	resourceNode2 = resourceNode2.AddToken(resNodeInitialStake2)
	resourceNode3 = resourceNode3.AddToken(resNodeInitialStake3)
	resourceNode4 = resourceNode4.AddToken(resNodeInitialStake4)
	resourceNode5 = resourceNode5.AddToken(resNodeInitialStake5)

	resourceNode1.Status = stakingtypes.Bonded
	resourceNode2.Status = stakingtypes.Bonded
	resourceNode3.Status = stakingtypes.Bonded
	resourceNode4.Status = stakingtypes.Bonded
	resourceNode5.Status = stakingtypes.Bonded

	var resourceNodes []registertypes.ResourceNode
	resourceNodes = append(resourceNodes, resourceNode1)
	resourceNodes = append(resourceNodes, resourceNode2)
	resourceNodes = append(resourceNodes, resourceNode3)
	resourceNodes = append(resourceNodes, resourceNode4)
	resourceNodes = append(resourceNodes, resourceNode5)
	return resourceNodes
}

func setupMultipleResourceNodes(resOwnerPrivKeys []*secp256k1.PrivKey, resNodePubKeys []cryptotypes.PubKey, resOwners []sdk.AccAddress, resNodeNetworkIds []stratos.SdsAddress) []registertypes.ResourceNode {
	if len(resOwnerPrivKeys) != len(resNodePubKeys) ||
		len(resNodePubKeys) != len(resOwners) ||
		len(resOwners) != len(resNodeNetworkIds) {
		return nil
	}

	numOfNodes := len(resOwnerPrivKeys)
	resourceNodes := make([]registertypes.ResourceNode, 0, numOfNodes)

	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")
	nodeType := registertypes.STORAGE

	for i, _ := range resOwnerPrivKeys {
		resourceNodeTmp, _ := registertypes.NewResourceNode(resNodeNetworkIds[i], resNodePubKeys[i], resOwners[i], registertypes.NewDescription("sds://resourceNode"+strconv.Itoa(i+1), "", "", "", ""), nodeType, time)
		resourceNodeTmp = resourceNodeTmp.AddToken(resNodeInitialStakeForMultipleNodes)
		resourceNodeTmp.Status = stakingtypes.Bonded
		resourceNodes = append(resourceNodes, resourceNodeTmp)
	}

	return resourceNodes
}

func setupAllMetaNodes() []registertypes.MetaNode {
	var indexingNodes []registertypes.MetaNode

	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")
	indexingNode1, _ := registertypes.NewMetaNode(stratos.SdsAddress(idxNodeAddr1), idxNodePubKey1, idxOwner1, registertypes.NewDescription("sds://indexingNode1", "", "", "", ""), time)
	indexingNode2, _ := registertypes.NewMetaNode(stratos.SdsAddress(idxNodeAddr2), idxNodePubKey2, idxOwner2, registertypes.NewDescription("sds://indexingNode2", "", "", "", ""), time)
	indexingNode3, _ := registertypes.NewMetaNode(stratos.SdsAddress(idxNodeAddr3), idxNodePubKey3, idxOwner3, registertypes.NewDescription("sds://indexingNode3", "", "", "", ""), time)

	indexingNode1.Suspend = false
	indexingNode2.Suspend = false
	indexingNode3.Suspend = false

	indexingNode1 = indexingNode1.AddToken(idxNodeInitialStake1)
	indexingNode2 = indexingNode2.AddToken(idxNodeInitialStake2)
	indexingNode3 = indexingNode3.AddToken(idxNodeInitialStake3)

	indexingNode1.Status = stakingtypes.Bonded
	indexingNode2.Status = stakingtypes.Bonded
	indexingNode3.Status = stakingtypes.Bonded

	indexingNodes = append(indexingNodes, indexingNode1)
	indexingNodes = append(indexingNodes, indexingNode2)
	indexingNodes = append(indexingNodes, indexingNode3)

	return indexingNodes
}

var (
	initialTotalStakesStore   = sdk.NewInt(1500000000000)
	effectiveTotalStakesStore = sdk.NewInt(1500000000000)
	remainOzoneLimitStore     = sdk.NewInt(1500000000000000)
	totalUnissuedPrepayStore  = sdk.ZeroInt()
	uozPrice                  = sdk.NewDecWithPrec(1000000, 9)
	//priceChangeChan = make(chan PriceChangeEvent, 0)
)

type PriceChangeEvent struct {
	stakeDelta          sdk.Int
	unissuedPrepayDelta sdk.Int
}

func simulatePriceChange(priceChangeEvent *PriceChangeEvent, uozPriceFactorsBefore *UozPriceFactors) (*UozPriceFactors, sdk.Dec, sdk.Dec) {
	uozPriceFactorsAfter := &UozPriceFactors{}
	uozPriceFactorsAfter.InitialTotalStakes = uozPriceFactorsBefore.InitialTotalStakes
	//uozPriceFactorsAfter.EffectiveTotalStakes = registerKeeper.GetEffectiveGenesisStakeTotal(ctx)
	uozPriceFactorsAfter.TotalUnissuedPrepay = uozPriceFactorsBefore.TotalUnissuedPrepay.Add(priceChangeEvent.unissuedPrepayDelta)
	uozPriceFactorsAfter.StakeAndPrepay = uozPriceFactorsAfter.InitialTotalStakes.Add(uozPriceFactorsAfter.TotalUnissuedPrepay)
	deltaUozLimit := sdk.ZeroInt()
	if !priceChangeEvent.stakeDelta.Equal(sdk.ZeroInt()) {
		ozoneLimitChangeByStake := uozPriceFactorsBefore.OzoneLimit.ToDec().Quo(uozPriceFactorsBefore.InitialTotalStakes.ToDec()).Mul(priceChangeEvent.stakeDelta.ToDec()).TruncateInt()
		deltaUozLimit = deltaUozLimit.Add(ozoneLimitChangeByStake)
	}
	if !priceChangeEvent.unissuedPrepayDelta.Equal(sdk.ZeroInt()) {
		ozoneLimitChangeByPrepay := uozPriceFactorsBefore.OzoneLimit.ToDec().
			Mul(uozPriceFactorsBefore.InitialTotalStakes.Add(uozPriceFactorsBefore.TotalUnissuedPrepay).ToDec()).
			Quo(uozPriceFactorsBefore.InitialTotalStakes.Add(uozPriceFactorsBefore.TotalUnissuedPrepay).Add(priceChangeEvent.unissuedPrepayDelta).ToDec()).
			TruncateInt().
			Sub(uozPriceFactorsBefore.OzoneLimit)
		deltaUozLimit = deltaUozLimit.Add(ozoneLimitChangeByPrepay)
	}

	uozPriceFactorsAfter.OzoneLimit = uozPriceFactorsBefore.OzoneLimit.Add(deltaUozLimit)
	uozPriceFactorsAfter.UOzonePrice = uozPriceFactorsAfter.StakeAndPrepay.ToDec().Quo(uozPriceFactorsAfter.OzoneLimit.ToDec())

	uozPriceDelta := uozPriceFactorsAfter.UOzonePrice.Sub(uozPriceFactorsBefore.UOzonePrice)
	initialTotalStakesDelta := uozPriceFactorsAfter.InitialTotalStakes.Sub(uozPriceFactorsBefore.InitialTotalStakes)
	//effectiveTotalStakesDelta := uozPriceFactorsAfter.EffectiveTotalStakes.Sub(uozPriceFactorsBefore.EffectiveTotalStakes)
	totalUnissuedPrepayDelta := uozPriceFactorsAfter.TotalUnissuedPrepay.Sub(uozPriceFactorsBefore.TotalUnissuedPrepay)
	stakeAndPrepayDelta := uozPriceFactorsAfter.StakeAndPrepay.Sub(uozPriceFactorsBefore.StakeAndPrepay)
	ozoneLimitDelta := uozPriceFactorsAfter.OzoneLimit.Sub(uozPriceFactorsBefore.OzoneLimit)

	uozPricePercentage := uozPriceDelta.Quo(uozPriceFactorsBefore.UOzonePrice).MulInt(sdk.NewInt(100))
	//initialTotalStakesPercentage := initialTotalStakesDelta.Quo(uozPriceFactorsBefore.InitialTotalStakes)
	//effectiveTotalStakesPercentage := effectiveTotalStakesDelta.Quo(uozPriceFactorsBefore.EffectiveTotalStakes)
	//totalUnissuedPrepayPercentage := totalUnissuedPrepayDelta.Quo(uozPriceFactorsBefore.TotalUnissuedPrepay)
	//stakeAndPrepayPercentage := stakeAndPrepayDelta.Quo(uozPriceFactorsBefore.StakeAndPrepay)
	ozoneLimitPercentage := ozoneLimitDelta.ToDec().Quo(uozPriceFactorsBefore.OzoneLimit.ToDec()).MulInt(sdk.NewInt(100))

	println("===>>>>>>>>>>>>>>     Current Uoz Price    ===>>>>>>>>>>>>>>")
	println("UOzonePrice: 									" + uozPriceFactorsAfter.UOzonePrice.String() + "(delta: " + uozPriceDelta.String() + ", " + uozPricePercentage.String()[:5] + "%)")
	println("InitialTotalStakes: 							" + uozPriceFactorsAfter.InitialTotalStakes.String() + "(delta: " + initialTotalStakesDelta.String() + ")")
	//println("EffectiveTotalStakes: 							" + uozPriceFactorsAfter.EffectiveTotalStakes.String() + "(delta: " + effectiveTotalStakesDelta.String() + ")")
	println("TotalUnissuedPrepay: 							" + uozPriceFactorsAfter.TotalUnissuedPrepay.String() + "(delta: " + totalUnissuedPrepayDelta.String() + ")")
	println("InitialTotalStakes+TotalUnissuedPrepay:			" + uozPriceFactorsAfter.StakeAndPrepay.String() + "(delta: " + stakeAndPrepayDelta.String() + ")")
	println("OzoneLimit: 									" + uozPriceFactorsAfter.OzoneLimit.String() + "(delta: " + ozoneLimitDelta.String() + ", " + ozoneLimitPercentage.String()[:5] + "%)")

	return uozPriceFactorsAfter, uozPricePercentage, ozoneLimitPercentage
}

func TestOzPriceChangePrepay(t *testing.T) {
	NUM_OF_SAMPLE := 100
	dataToExcel := make([]UozPriceFactors, 0, NUM_OF_SAMPLE)
	/********************* initialize mock app *********************/
	//mApp, k, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper := getMockApp(t)
	accs, balances := setupAccounts()
	//stApp := app.SetupWithGenesisAccounts(accs, chainID, balances...)
	validators := make([]*tmtypes.Validator, 0)
	valSet := tmtypes.NewValidatorSet(validators)
	metaNodes := setupAllMetaNodes()
	//resourceNodes := setupAllResourceNodes()
	resourceNodes := make([]registertypes.ResourceNode, 0)

	stApp := app.SetupWithGenesisNodeSet(t, true, valSet, metaNodes, resourceNodes, accs, totalUnissuedPrepay, chainID, balances...)

	accountKeeper := stApp.GetAccountKeeper()
	//bankKeeper := stApp.GetBankKeeper()
	registerKeeper := stApp.GetRegisterKeeper()
	//potKeeper := stApp.GetPotKeeper()

	/********************* foundation account deposit *********************/
	header := tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := stApp.BaseApp.NewContext(true, header)

	foundationDepositMsg := pottypes.NewMsgFoundationDeposit(foundationDeposit, foundationDepositorAccAddr)
	txGen := app.MakeTestEncodingConfig().TxConfig

	foundationDepositorAcc := accountKeeper.GetAccount(ctx, foundationDepositorAccAddr)
	accNum := foundationDepositorAcc.GetAccountNumber()
	accSeq := foundationDepositorAcc.GetSequence()
	_, _, err := app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{foundationDepositMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, foundationDepositorPrivKey)
	require.NoError(t, err)
	foundationAccountAddr := accountKeeper.GetModuleAddress(pottypes.FoundationAccount)
	app.CheckBalance(t, stApp, foundationAccountAddr, foundationDeposit)

	/********************* create validator with 50% commission *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	description := stakingtypes.NewDescription("foo_moniker", chainID, "", "", "")
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(valOpValAddr1, valConsPubk1, sdk.NewCoin("ustos", valInitialStake), description, commission, sdk.OneInt())

	valOpAcc1 := accountKeeper.GetAccount(ctx, valOpAccAddr1)
	accNum = valOpAcc1.GetAccountNumber()
	accSeq = valOpAcc1.GetSequence()
	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createValidatorMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, valOpPrivKey1)
	require.NoError(t, err)
	app.CheckBalance(t, stApp, valOpAccAddr1, nil)

	/********************** commit **********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	validator := checkValidator(t, stApp, valOpValAddr1, true)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, valInitialStake, validator.BondedTokens()))
	_, uozSupply := registerKeeper.UozSupply(ctx)
	uozPriceFactorsSeq0, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, UozPriceFactors{
		UOzonePrice:        registerKeeper.CurrUozPrice(ctx),
		InitialTotalStakes: registerKeeper.GetInitialGenesisStakeTotal(ctx),
		//EffectiveTotalStakes: registerKeeper.GetEffectiveGenesisStakeTotal(ctx),
		TotalUnissuedPrepay: registerKeeper.GetTotalUnissuedPrepay(ctx).Amount,
		StakeAndPrepay:      registerKeeper.GetInitialGenesisStakeTotal(ctx).Add(registerKeeper.GetTotalUnissuedPrepay(ctx).Amount),
		OzoneLimit:          registerKeeper.GetRemainingOzoneLimit(ctx),
		UozSupply:           uozSupply,
	})

	// start testing
	println("\n********************************* Deliver Prepay Tx START ********************************************")

	priceBefore := uozPriceFactorsSeq0
	priceAfter := uozPriceFactorsSeq0
	dataToExcel = append(dataToExcel, priceBefore)

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		prepayMsg := setupPrepayMsg()
		/********************* deliver tx *********************/

		resOwnerAcc := accountKeeper.GetAccount(ctx, resOwner1)
		ownerAccNum := resOwnerAcc.GetAccountNumber()
		ownerAccSeq := resOwnerAcc.GetSequence()

		_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{prepayMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKey1)
		require.NoError(t, err)
		/********************* commit & check result *********************/
		header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
		stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
		ctx = stApp.BaseApp.NewContext(true, header)
		priceAfter, uozPricePercentage, ozoneLimitPercentage = printCurrUozPrice(ctx, registerKeeper, priceBefore)
		dataToExcel = append(dataToExcel, priceAfter)
		require.True(t, uozPricePercentage.GT(sdk.ZeroDec()), "Uoz price should increase after PREPAY")
		require.True(t, ozoneLimitPercentage.LT(sdk.ZeroDec()), "OzLimit should decrease after PREPAY")
		println("********************************* Deliver Prepay Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

		priceBefore = priceAfter
	}
	exportToCSV(dataToExcel)
}

func TestOzPriceChangeAddMultipleResourceNodeAndThenRemove(t *testing.T) {
	NUM_OF_SAMPLE := 100
	dataToExcel := make([]UozPriceFactors, 0, NUM_OF_SAMPLE)

	resOwners := make([]sdk.AccAddress, 0, NUM_OF_SAMPLE)
	resOwnerPrivKeys := make([]*secp256k1.PrivKey, 0, NUM_OF_SAMPLE)
	resOwnerPubkeys := make([]cryptotypes.PubKey, 0, NUM_OF_SAMPLE)
	resNodeNetworkIds := make([]stratos.SdsAddress, 0, NUM_OF_SAMPLE)

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		resOwnerPrivKeyTmp := secp256k1.GenPrivKey()
		resOwnerPrivKeys = append(resOwnerPrivKeys, resOwnerPrivKeyTmp)
		resOwnerPubKeyTmp := resOwnerPrivKeyTmp.PubKey()
		resOwnerPubkeys = append(resOwnerPubkeys, resOwnerPrivKeyTmp.PubKey())
		resNodeAddrTmp := sdk.AccAddress(resOwnerPubKeyTmp.Address())
		resOwners = append(resOwners, resNodeAddrTmp)
		resNodeNetworkIds = append(resNodeNetworkIds, stratos.SdsAddress(resNodeAddrTmp))
	}

	/********************* initialize mock app *********************/
	//mApp, k, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper := getMockApp(t)
	accs, balances := setupAccountsMultipleResNodes(resOwners)
	//stApp := app.SetupWithGenesisAccounts(accs, chainID, balances...)
	validators := make([]*tmtypes.Validator, 0)
	valSet := tmtypes.NewValidatorSet(validators)
	metaNodes := setupAllMetaNodes()
	//resourceNodes := setupAllResourceNodes()
	resourceNodes := make([]registertypes.ResourceNode, 0)

	stApp := app.SetupWithGenesisNodeSet(t, true, valSet, metaNodes, resourceNodes, accs, totalUnissuedPrepay, chainID, balances...)

	accountKeeper := stApp.GetAccountKeeper()
	//bankKeeper := stApp.GetBankKeeper()
	registerKeeper := stApp.GetRegisterKeeper()
	//potKeeper := stApp.GetPotKeeper()

	/********************* foundation account deposit *********************/
	header := tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := stApp.BaseApp.NewContext(true, header)

	foundationDepositMsg := pottypes.NewMsgFoundationDeposit(foundationDeposit, foundationDepositorAccAddr)
	txGen := app.MakeTestEncodingConfig().TxConfig

	foundationDepositorAcc := accountKeeper.GetAccount(ctx, foundationDepositorAccAddr)
	accNum := foundationDepositorAcc.GetAccountNumber()
	accSeq := foundationDepositorAcc.GetSequence()
	_, _, err := app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{foundationDepositMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, foundationDepositorPrivKey)
	require.NoError(t, err)
	foundationAccountAddr := accountKeeper.GetModuleAddress(pottypes.FoundationAccount)
	app.CheckBalance(t, stApp, foundationAccountAddr, foundationDeposit)

	/********************* create validator with 50% commission *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	description := stakingtypes.NewDescription("foo_moniker", chainID, "", "", "")
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(valOpValAddr1, valConsPubk1, sdk.NewCoin("ustos", valInitialStake), description, commission, sdk.OneInt())

	valOpAcc1 := accountKeeper.GetAccount(ctx, valOpAccAddr1)
	accNum = valOpAcc1.GetAccountNumber()
	accSeq = valOpAcc1.GetSequence()
	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createValidatorMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, valOpPrivKey1)
	require.NoError(t, err)
	app.CheckBalance(t, stApp, valOpAccAddr1, nil)

	/********************** commit **********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	validator := checkValidator(t, stApp, valOpValAddr1, true)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, valInitialStake, validator.BondedTokens()))
	_, uozSupply := registerKeeper.UozSupply(ctx)
	uozPriceFactorsSeq0, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, UozPriceFactors{
		UOzonePrice:        registerKeeper.CurrUozPrice(ctx),
		InitialTotalStakes: registerKeeper.GetInitialGenesisStakeTotal(ctx),
		//EffectiveTotalStakes: registerKeeper.GetEffectiveGenesisStakeTotal(ctx),
		TotalUnissuedPrepay: registerKeeper.GetTotalUnissuedPrepay(ctx).Amount,
		StakeAndPrepay:      registerKeeper.GetInitialGenesisStakeTotal(ctx).Add(registerKeeper.GetTotalUnissuedPrepay(ctx).Amount),
		OzoneLimit:          registerKeeper.GetRemainingOzoneLimit(ctx),
		UozSupply:           uozSupply,
	})

	// start testing
	println("********************************* Deliver CreateResourceNode Tx START ********************************************")

	priceBefore := uozPriceFactorsSeq0
	priceAfter := uozPriceFactorsSeq0
	dataToExcel = append(dataToExcel, priceBefore)

	//resOwnerAcc := accountKeeper.GetAccount(ctx, resOwner1)
	//ownerAccNum := resOwnerAcc.GetAccountNumber()
	//ownerAccSeq := resOwnerAcc.GetSequence()

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		createResourceNodeMsg := setupMsgCreateResourceNode(i, resNodeNetworkIds[i], resOwnerPubkeys[i], resOwners[i])
		/********************* deliver tx *********************/

		resOwnerAcc := accountKeeper.GetAccount(ctx, resOwners[i])
		ownerAccNum := resOwnerAcc.GetAccountNumber()
		ownerAccSeq := resOwnerAcc.GetSequence()

		_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKeys[i])
		require.NoError(t, err)
		/********************* commit & check result *********************/
		header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
		stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
		ctx = stApp.BaseApp.NewContext(true, header)

		priceAfter, uozPricePercentage, ozoneLimitPercentage = printCurrUozPrice(ctx, registerKeeper, priceBefore)
		require.True(t, uozPricePercentage.LT(sdk.ZeroDec()), "Uoz price should decrease after CreateResourceNode")
		require.True(t, ozoneLimitPercentage.GT(sdk.ZeroDec()), "OzLimit should increase after CreateResourceNode")
		println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

		dataToExcel = append(dataToExcel, priceAfter)
		priceBefore = priceAfter
	}

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		createResourceNodeMsg := setupMsgRemoveResourceNode(i, resNodeNetworkIds[i], resOwners[i])
		/********************* deliver tx *********************/

		resOwnerAcc := accountKeeper.GetAccount(ctx, resOwners[i])
		ownerAccNum := resOwnerAcc.GetAccountNumber()
		ownerAccSeq := resOwnerAcc.GetSequence()

		_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKeys[i])
		require.NoError(t, err)
		/********************* commit & check result *********************/
		header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
		stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
		ctx = stApp.BaseApp.NewContext(true, header)

		priceAfter, uozPricePercentage, ozoneLimitPercentage = printCurrUozPrice(ctx, registerKeeper, priceBefore)
		require.True(t, uozPricePercentage.GT(sdk.ZeroDec()), "Uoz price should increase after RemoveResourceNode")
		require.True(t, ozoneLimitPercentage.LT(sdk.ZeroDec()), "OzLimit should decrease after RemoveResourceNode")
		println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

		dataToExcel = append(dataToExcel, priceAfter)
		priceBefore = priceAfter
	}

	exportToCSV(dataToExcel)

}

func TestOzPriceChangeRemoveMultipleResourceNodeAfterGenesis(t *testing.T) {
	NUM_OF_SAMPLE := 100
	dataToExcel := make([]UozPriceFactors, 0, NUM_OF_SAMPLE)

	resOwners := make([]sdk.AccAddress, 0, NUM_OF_SAMPLE)
	resOwnerPrivKeys := make([]*secp256k1.PrivKey, 0, NUM_OF_SAMPLE)
	resOwnerPubkeys := make([]cryptotypes.PubKey, 0, NUM_OF_SAMPLE)
	resNodeNetworkIds := make([]stratos.SdsAddress, 0, NUM_OF_SAMPLE)

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		resOwnerPrivKeyTmp := secp256k1.GenPrivKey()
		resOwnerPrivKeys = append(resOwnerPrivKeys, resOwnerPrivKeyTmp)
		resOwnerPubKeyTmp := resOwnerPrivKeyTmp.PubKey()
		resOwnerPubkeys = append(resOwnerPubkeys, resOwnerPrivKeyTmp.PubKey())
		resNodeAddrTmp := sdk.AccAddress(resOwnerPubKeyTmp.Address())
		resOwners = append(resOwners, resNodeAddrTmp)
		resNodeNetworkIds = append(resNodeNetworkIds, stratos.SdsAddress(resNodeAddrTmp))
	}

	/********************* initialize mock app *********************/
	//mApp, k, stakingKeeper, bankKeeper, supplyKeeper, registerKeeper := getMockApp(t)
	accs, balances := setupAccountsMultipleResNodes(resOwners)
	//stApp := app.SetupWithGenesisAccounts(accs, chainID, balances...)
	validators := make([]*tmtypes.Validator, 0)
	valSet := tmtypes.NewValidatorSet(validators)
	metaNodes := setupAllMetaNodes()
	resourceNodes := setupMultipleResourceNodes(resOwnerPrivKeys, resOwnerPubkeys, resOwners, resNodeNetworkIds)

	stApp := app.SetupWithGenesisNodeSet(t, true, valSet, metaNodes, resourceNodes, accs, totalUnissuedPrepay, chainID, balances...)

	accountKeeper := stApp.GetAccountKeeper()
	//bankKeeper := stApp.GetBankKeeper()
	registerKeeper := stApp.GetRegisterKeeper()
	//potKeeper := stApp.GetPotKeeper()

	/********************* foundation account deposit *********************/
	header := tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := stApp.BaseApp.NewContext(true, header)

	foundationDepositMsg := pottypes.NewMsgFoundationDeposit(foundationDeposit, foundationDepositorAccAddr)
	txGen := app.MakeTestEncodingConfig().TxConfig

	foundationDepositorAcc := accountKeeper.GetAccount(ctx, foundationDepositorAccAddr)
	accNum := foundationDepositorAcc.GetAccountNumber()
	accSeq := foundationDepositorAcc.GetSequence()
	_, _, err := app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{foundationDepositMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, foundationDepositorPrivKey)
	require.NoError(t, err)
	foundationAccountAddr := accountKeeper.GetModuleAddress(pottypes.FoundationAccount)
	app.CheckBalance(t, stApp, foundationAccountAddr, foundationDeposit)

	/********************* create validator with 50% commission *********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	commission := stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	description := stakingtypes.NewDescription("foo_moniker", chainID, "", "", "")
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(valOpValAddr1, valConsPubk1, sdk.NewCoin("ustos", valInitialStake), description, commission, sdk.OneInt())

	valOpAcc1 := accountKeeper.GetAccount(ctx, valOpAccAddr1)
	accNum = valOpAcc1.GetAccountNumber()
	accSeq = valOpAcc1.GetSequence()
	_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createValidatorMsg}, chainID, []uint64{accNum}, []uint64{accSeq}, true, true, valOpPrivKey1)
	require.NoError(t, err)
	app.CheckBalance(t, stApp, valOpAccAddr1, nil)

	/********************** commit **********************/
	header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
	stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx = stApp.BaseApp.NewContext(true, header)

	validator := checkValidator(t, stApp, valOpValAddr1, true)
	require.Equal(t, stakingtypes.Bonded, validator.Status)
	require.True(sdk.IntEq(t, valInitialStake, validator.BondedTokens()))
	_, uozSupply := registerKeeper.UozSupply(ctx)
	uozPriceFactorsSeq0, uozPricePercentage, ozoneLimitPercentage := printCurrUozPrice(ctx, registerKeeper, UozPriceFactors{
		UOzonePrice:        registerKeeper.CurrUozPrice(ctx),
		InitialTotalStakes: registerKeeper.GetInitialGenesisStakeTotal(ctx),
		//EffectiveTotalStakes: registerKeeper.GetEffectiveGenesisStakeTotal(ctx),
		TotalUnissuedPrepay: registerKeeper.GetTotalUnissuedPrepay(ctx).Amount,
		StakeAndPrepay:      registerKeeper.GetInitialGenesisStakeTotal(ctx).Add(registerKeeper.GetTotalUnissuedPrepay(ctx).Amount),
		OzoneLimit:          registerKeeper.GetRemainingOzoneLimit(ctx),
		UozSupply:           uozSupply,
	})

	// start testing
	println("********************************* Deliver RemoveResourceNode Tx START ********************************************")

	priceBefore := uozPriceFactorsSeq0
	priceAfter := uozPriceFactorsSeq0
	dataToExcel = append(dataToExcel, priceBefore)

	for i := 0; i < NUM_OF_SAMPLE; i++ {
		createResourceNodeMsg := setupMsgRemoveResourceNode(i, resNodeNetworkIds[i], resOwners[i])
		/********************* deliver tx *********************/

		resOwnerAcc := accountKeeper.GetAccount(ctx, resOwners[i])
		ownerAccNum := resOwnerAcc.GetAccountNumber()
		ownerAccSeq := resOwnerAcc.GetSequence()

		_, _, err = app.SignCheckDeliver(t, txGen, stApp.BaseApp, header, []sdk.Msg{createResourceNodeMsg}, chainID, []uint64{ownerAccNum}, []uint64{ownerAccSeq}, true, true, resOwnerPrivKeys[i])
		require.NoError(t, err)
		/********************* commit & check result *********************/
		header = tmproto.Header{Height: stApp.LastBlockHeight() + 1, ChainID: chainID}
		stApp.BeginBlock(abci.RequestBeginBlock{Header: header})
		ctx = stApp.BaseApp.NewContext(true, header)

		priceAfter, uozPricePercentage, ozoneLimitPercentage = printCurrUozPrice(ctx, registerKeeper, priceBefore)
		require.True(t, uozPricePercentage.GT(sdk.ZeroDec()), "Uoz price should increase after RemoveResourceNode")
		require.True(t, ozoneLimitPercentage.LT(sdk.ZeroDec()), "OzLimit should decrease after RemoveResourceNode")
		println("********************************* Deliver CreateResourceNode Tx END ********************************************\n\n...\n[NEXT TEST CASE]")

		dataToExcel = append(dataToExcel, priceAfter)
		priceBefore = priceAfter
	}

	exportToCSV(dataToExcel)

}

func exportToCSV(factors []UozPriceFactors) {
	fmt.Printf("\n%v, %v, %v, %v, %v, %v, %v", "Index", "InitialTotalStake", "TotalUnissuedPrepay", "StakeAndPrepay",
		"UOzonePrice", "RemainingOzoneLimit", "TotalUozSupply")
	for i, factor := range factors {
		fmt.Printf("\n%v, %v, %v, %v, %v, %v, %v", i+1, factor.InitialTotalStakes.String(), factor.TotalUnissuedPrepay.String(), factor.StakeAndPrepay.String(),
			factor.UOzonePrice.String(), factor.OzoneLimit.String(), factor.UozSupply.String())
	}
	println("\n")
}
