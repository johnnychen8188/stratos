package pot

//
//import (
//	"os"
//	"testing"
//	"time"
//
//	"github.com/cosmos/cosmos-sdk/baseapp"
//	"github.com/cosmos/cosmos-sdk/codec"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
//	"github.com/cosmos/cosmos-sdk/x/mock"
//	stratos "github.com/stratosnet/stratos-chain/types"
//	"github.com/stratosnet/stratos-chain/x/register"
//	"github.com/stretchr/testify/require"
//	abci "github.com/tendermint/tendermint/abci/types"
//	"github.com/tendermint/tendermint/crypto"
//	"github.com/tendermint/tendermint/crypto/secp256k1"
//)
//
//const (
//	chainID             = ""
//	StratosBech32Prefix = "st"
//
//	stos2ustos = 1000000000
//)
//
//var (
//	AccountPubKeyPrefix    = StratosBech32Prefix + "pub"
//	ValidatorAddressPrefix = StratosBech32Prefix + "valoper"
//	ValidatorPubKeyPrefix  = StratosBech32Prefix + "valoperpub"
//	ConsNodeAddressPrefix  = StratosBech32Prefix + "valcons"
//	ConsNodePubKeyPrefix   = StratosBech32Prefix + "valconspub"
//	SdsNodeP2PKeyPrefix    = StratosBech32Prefix + "sdsp2p"
//
//	resNodeSlashingUOZAmt1 = sdk.NewInt(1000000000000000000)
//
//	resourceNodeVolume1 = sdk.NewInt(500000)
//	resourceNodeVolume2 = sdk.NewInt(300000)
//	resourceNodeVolume3 = sdk.NewInt(200000)
//
//	depositForSendingTx, _    = sdk.NewIntFromString("100000000000000000000000000000")
//	totalUnissuedPrepayVal, _ = sdk.NewIntFromString("1000000000000")
//	totalUnissuedPrepay       = sdk.NewCoin("ustos", totalUnissuedPrepayVal)
//	initialUOzonePrice        = sdk.NewDecWithPrec(10000000, 9) // 0.001 ustos -> 1 uoz
//
//	foundationDepositorPrivKey = secp256k1.GenPrivKey()
//	foundationDepositorAccAddr = sdk.AccAddress(foundationDepositorPrivKey.PubKey().Address())
//	foundationDeposit          = sdk.NewCoins(sdk.NewCoin("ustos", sdk.NewInt(40000000000000000)), sdk.NewCoin("utros", sdk.NewInt(40000000000000000)))
//
//	resOwnerPrivKey1 = secp256k1.GenPrivKey()
//	resOwnerPrivKey2 = secp256k1.GenPrivKey()
//	resOwnerPrivKey3 = secp256k1.GenPrivKey()
//	resOwnerPrivKey4 = secp256k1.GenPrivKey()
//	resOwnerPrivKey5 = secp256k1.GenPrivKey()
//	idxOwnerPrivKey1 = secp256k1.GenPrivKey()
//	idxOwnerPrivKey2 = secp256k1.GenPrivKey()
//	idxOwnerPrivKey3 = secp256k1.GenPrivKey()
//
//	resOwner1 = sdk.AccAddress(resOwnerPrivKey1.PubKey().Address())
//	resOwner2 = sdk.AccAddress(resOwnerPrivKey2.PubKey().Address())
//	resOwner3 = sdk.AccAddress(resOwnerPrivKey3.PubKey().Address())
//	resOwner4 = sdk.AccAddress(resOwnerPrivKey4.PubKey().Address())
//	resOwner5 = sdk.AccAddress(resOwnerPrivKey5.PubKey().Address())
//	idxOwner1 = sdk.AccAddress(idxOwnerPrivKey1.PubKey().Address())
//	idxOwner2 = sdk.AccAddress(idxOwnerPrivKey2.PubKey().Address())
//	idxOwner3 = sdk.AccAddress(idxOwnerPrivKey3.PubKey().Address())
//
//	resNodePubKey1       = secp256k1.GenPrivKey().PubKey()
//	resNodeAddr1         = sdk.AccAddress(resNodePubKey1.Address())
//	resNodeNetworkId1    = stratos.SdsAddress(resNodePubKey1.Address())
//	resNodeInitialStake1 = sdk.NewInt(3 * stos2ustos)
//
//	resNodePubKey2       = secp256k1.GenPrivKey().PubKey()
//	resNodeAddr2         = sdk.AccAddress(resNodePubKey2.Address())
//	resNodeNetworkId2    = stratos.SdsAddress(resNodePubKey2.Address())
//	resNodeInitialStake2 = sdk.NewInt(3 * stos2ustos)
//
//	resNodePubKey3       = secp256k1.GenPrivKey().PubKey()
//	resNodeAddr3         = sdk.AccAddress(resNodePubKey3.Address())
//	resNodeNetworkId3    = stratos.SdsAddress(resNodePubKey3.Address())
//	resNodeInitialStake3 = sdk.NewInt(3 * stos2ustos)
//
//	resNodePubKey4       = secp256k1.GenPrivKey().PubKey()
//	resNodeAddr4         = sdk.AccAddress(resNodePubKey4.Address())
//	resNodeNetworkId4    = stratos.SdsAddress(resNodePubKey4.Address())
//	resNodeInitialStake4 = sdk.NewInt(3 * stos2ustos)
//
//	resNodePubKey5       = secp256k1.GenPrivKey().PubKey()
//	resNodeAddr5         = sdk.AccAddress(resNodePubKey5.Address())
//	resNodeNetworkId5    = stratos.SdsAddress(resNodePubKey5.Address())
//	resNodeInitialStake5 = sdk.NewInt(3 * stos2ustos)
//
//	idxNodePrivKey1      = secp256k1.GenPrivKey()
//	idxNodePubKey1       = idxNodePrivKey1.PubKey()
//	idxNodeAddr1         = sdk.AccAddress(idxNodePubKey1.Address())
//	idxNodeNetworkId1    = stratos.SdsAddress(idxNodePubKey1.Address())
//	idxNodeInitialStake1 = sdk.NewInt(5 * stos2ustos)
//
//	idxNodePubKey2       = secp256k1.GenPrivKey().PubKey()
//	idxNodeAddr2         = sdk.AccAddress(idxNodePubKey2.Address())
//	idxNodeNetworkId2    = stratos.SdsAddress(idxNodePubKey2.Address())
//	idxNodeInitialStake2 = sdk.NewInt(5 * stos2ustos)
//
//	idxNodePubKey3       = secp256k1.GenPrivKey().PubKey()
//	idxNodeAddr3         = sdk.AccAddress(idxNodePubKey3.Address())
//	idxNodeNetworkId3    = stratos.SdsAddress(idxNodePubKey3.Address())
//	idxNodeInitialStake3 = sdk.NewInt(5 * stos2ustos)
//
//	valOpPrivKey1 = secp256k1.GenPrivKey()
//	valOpPubKey1  = valOpPrivKey1.PubKey()
//	valOpValAddr1 = sdk.ValAddress(valOpPubKey1.Address())
//	valOpAccAddr1 = sdk.AccAddress(valOpPubKey1.Address())
//
//	valConsPrivKey1 = secp256k1.GenPrivKey()
//	valConsPubk1    = valConsPrivKey1.PubKey()
//	valInitialStake = sdk.NewInt(15 * stos2ustos)
//)
//
//func TestMain(m *testing.M) {
//	config := stratos.GetConfig()
//	config.Seal()
//	exitVal := m.Run()
//	os.Exit(exitVal)
//}
//
//func setupAccounts(mApp *mock.App) []authexported.Account {
//
//	//************************** setup resource nodes owners' accounts **************************
//	resOwnerAcc1 := &auth.BaseAccount{
//		Address: resOwner1,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake1.Add(depositForSendingTx))},
//	}
//	resOwnerAcc2 := &auth.BaseAccount{
//		Address: resOwner2,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake2)},
//	}
//	resOwnerAcc3 := &auth.BaseAccount{
//		Address: resOwner3,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake3)},
//	}
//	resOwnerAcc4 := &auth.BaseAccount{
//		Address: resOwner4,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake4)},
//	}
//	resOwnerAcc5 := &auth.BaseAccount{
//		Address: resOwner5,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", resNodeInitialStake5)},
//	}
//
//	//************************** setup indexing nodes owners' accounts **************************
//	idxOwnerAcc1 := &auth.BaseAccount{
//		Address: idxOwner1,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake1)},
//	}
//	idxOwnerAcc2 := &auth.BaseAccount{
//		Address: idxOwner2,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake2)},
//	}
//	idxOwnerAcc3 := &auth.BaseAccount{
//		Address: idxOwner3,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", idxNodeInitialStake3)},
//	}
//
//	//************************** setup validator delegators' accounts **************************
//	valOwnerAcc1 := &auth.BaseAccount{
//		Address: valOpAccAddr1,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", valInitialStake)},
//	}
//
//	//************************** setup indexing nodes' accounts **************************
//	idxNodeAcc1 := &auth.BaseAccount{
//		Address: idxNodeAddr1,
//		Coins:   sdk.Coins{sdk.NewCoin("ustos", sdk.ZeroInt())},
//	}
//
//	foundationDepositorAcc := &auth.BaseAccount{
//		Address: foundationDepositorAccAddr,
//		Coins:   foundationDeposit,
//	}
//
//	accs := []authexported.Account{
//		resOwnerAcc1, resOwnerAcc2, resOwnerAcc3, resOwnerAcc4, resOwnerAcc5,
//		idxOwnerAcc1, idxOwnerAcc2, idxOwnerAcc3,
//		valOwnerAcc1,
//		foundationDepositorAcc,
//		idxNodeAcc1,
//	}
//
//	ctx1 := mApp.BaseApp.NewContext(true, abci.Header{})
//	ctx1.Logger().Info("idxNodeAcc1 -> " + idxNodeAcc1.String())
//
//	return accs
//}
//
//func setupAllResourceNodes() []register.ResourceNode {
//
//	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")
//	resourceNode1 := register.NewResourceNode(resNodeNetworkId1, resNodePubKey1, resOwner1, register.NewDescription("sds://resourceNode1", "", "", "", ""), 4, time)
//	resourceNode2 := register.NewResourceNode(resNodeNetworkId2, resNodePubKey2, resOwner2, register.NewDescription("sds://resourceNode2", "", "", "", ""), 4, time)
//	resourceNode3 := register.NewResourceNode(resNodeNetworkId3, resNodePubKey3, resOwner3, register.NewDescription("sds://resourceNode3", "", "", "", ""), 4, time)
//	resourceNode4 := register.NewResourceNode(resNodeNetworkId4, resNodePubKey4, resOwner4, register.NewDescription("sds://resourceNode4", "", "", "", ""), 4, time)
//	resourceNode5 := register.NewResourceNode(resNodeNetworkId5, resNodePubKey5, resOwner5, register.NewDescription("sds://resourceNode5", "", "", "", ""), 4, time)
//
//	resourceNode1 = resourceNode1.AddToken(resNodeInitialStake1)
//	resourceNode2 = resourceNode2.AddToken(resNodeInitialStake2)
//	resourceNode3 = resourceNode3.AddToken(resNodeInitialStake3)
//	resourceNode4 = resourceNode4.AddToken(resNodeInitialStake4)
//	resourceNode5 = resourceNode5.AddToken(resNodeInitialStake5)
//
//	resourceNode1.Status = sdk.Bonded
//	resourceNode2.Status = sdk.Bonded
//	resourceNode3.Status = sdk.Bonded
//	resourceNode4.Status = sdk.Bonded
//	resourceNode5.Status = sdk.Bonded
//
//	var resourceNodes []register.ResourceNode
//	resourceNodes = append(resourceNodes, resourceNode1)
//	resourceNodes = append(resourceNodes, resourceNode2)
//	resourceNodes = append(resourceNodes, resourceNode3)
//	resourceNodes = append(resourceNodes, resourceNode4)
//	resourceNodes = append(resourceNodes, resourceNode5)
//	return resourceNodes
//}
//
//func setupAllIndexingNodes() []register.IndexingNode {
//	var indexingNodes []register.IndexingNode
//
//	time, _ := time.Parse(time.RubyDate, "Fri Sep 24 10:37:13 -0400 2021")
//	indexingNode1 := register.NewIndexingNode(stratos.SdsAddress(idxNodeAddr1), idxNodePubKey1, idxOwner1, register.NewDescription("sds://indexingNode1", "", "", "", ""), time)
//	indexingNode2 := register.NewIndexingNode(stratos.SdsAddress(idxNodeAddr2), idxNodePubKey2, idxOwner2, register.NewDescription("sds://indexingNode2", "", "", "", ""), time)
//	indexingNode3 := register.NewIndexingNode(stratos.SdsAddress(idxNodeAddr3), idxNodePubKey3, idxOwner3, register.NewDescription("sds://indexingNode3", "", "", "", ""), time)
//
//	indexingNode1 = indexingNode1.AddToken(idxNodeInitialStake1)
//	indexingNode2 = indexingNode2.AddToken(idxNodeInitialStake2)
//	indexingNode3 = indexingNode3.AddToken(idxNodeInitialStake3)
//
//	indexingNode1.Status = sdk.Bonded
//	indexingNode2.Status = sdk.Bonded
//	indexingNode3.Status = sdk.Bonded
//
//	indexingNodes = append(indexingNodes, indexingNode1)
//	indexingNodes = append(indexingNodes, indexingNode2)
//	indexingNodes = append(indexingNodes, indexingNode3)
//
//	return indexingNodes
//
//}
//
//// SignCheckDeliver checks a generated signed transaction and simulates a
//// block commitment with the given transaction. A test assertion is made using
//// the parameter 'expPass' against the result. A corresponding result is
//// returned.
//func SignCheckDeliver(
//	t *testing.T, cdc *codec.Codec, app *baseapp.BaseApp, header abci.Header, msgs []sdk.Msg,
//	accNums, seq []uint64, expSimPass, expPass bool, priv ...crypto.PrivKey,
//) (sdk.GasInfo, *sdk.Result, error) {
//
//	tx := GenTx(msgs, accNums, seq, priv...)
//
//	txBytes, err := cdc.MarshalBinaryLengthPrefixed(tx)
//	require.Nil(t, err)
//
//	// Must simulate now as CheckTx doesn't run Msgs anymore
//	_, res, err := app.Simulate(txBytes, tx)
//
//	if expSimPass {
//		require.NoError(t, err)
//		require.NotNil(t, res)
//	} else {
//		require.Error(t, err)
//		require.Nil(t, res)
//	}
//
//	// Simulate a sending a transaction and committing a block
//	app.BeginBlock(abci.RequestBeginBlock{Header: header})
//	gInfo, res, err := app.Deliver(tx)
//
//	if expPass {
//		require.NoError(t, err)
//		require.NotNil(t, res)
//	} else {
//		require.Error(t, err)
//		require.Nil(t, res)
//	}
//
//	app.EndBlock(abci.RequestEndBlock{})
//	app.Commit()
//
//	return gInfo, res, err
//}
//
//// GenTx generates a signed mock transaction.
//func GenTx(msgs []sdk.Msg, accnums []uint64, seq []uint64, priv ...crypto.PrivKey) auth.StdTx {
//	// Make the transaction free
//	fee := auth.StdFee{
//		Amount: sdk.NewCoins(sdk.NewInt64Coin("foocoin", 0)),
//		Gas:    5000000,
//	}
//
//	sigs := make([]auth.StdSignature, len(priv))
//	memo := "testmemotestmemo"
//
//	for i, p := range priv {
//		sig, err := p.Sign(auth.StdSignBytes(chainID, accnums[i], seq[i], fee, msgs, memo))
//		if err != nil {
//			panic(err)
//		}
//
//		sigs[i] = auth.StdSignature{
//			PubKey:    p.PubKey(),
//			Signature: sig,
//		}
//	}
//
//	return auth.NewStdTx(msgs, fee, sigs, memo)
//}
