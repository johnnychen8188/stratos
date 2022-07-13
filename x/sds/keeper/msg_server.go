package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stratos "github.com/stratosnet/stratos-chain/types"
	"github.com/stratosnet/stratos-chain/x/sds/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// HandleMsgFileUpload Handles MsgFileUpload.
func (k msgServer) HandleMsgFileUpload(c context.Context, msg *types.MsgFileUpload) (*types.MsgFileUploadResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	reporter, err := stratos.SdsAddressFromBech32(msg.GetReporter())
	if err != nil {
		return &types.MsgFileUploadResponse{}, err
	}

	if _, found := k.RegisterKeeper.GetMetaNode(ctx, reporter); found == false {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "Reporter %s isn't an SP node", msg.GetReporter())
	}
	height := sdk.NewInt(ctx.BlockHeight())
	heightByteArr, _ := height.MarshalJSON()
	var heightReEncoded sdk.Int
	err = heightReEncoded.UnmarshalJSON(heightByteArr)
	if err != nil {
		return &types.MsgFileUploadResponse{}, err
	}

	fileInfo := types.NewFileInfo(&heightReEncoded, msg.Reporter, msg.Uploader)
	fileHashByte := []byte(msg.FileHash)
	k.SetFileHash(ctx, fileHashByte, fileInfo)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFileUpload,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From),
			sdk.NewAttribute(types.AttributeKeyReporter, msg.GetReporter()),
			sdk.NewAttribute(types.AttributeKeyUploader, msg.GetUploader()),
			sdk.NewAttribute(types.AttributeKeyFileHash, msg.GetFileHash()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetFrom()),
		),
	})

	return &types.MsgFileUploadResponse{}, nil
}

// HandleMsgPrepay Handles MsgPrepay.
func (k msgServer) HandleMsgPrepay(c context.Context, msg *types.MsgPrepay) (*types.MsgPrepayResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	//if k.BankKeeper.GetSendEnabled(ctx) == false {
	//	return nil, nil
	//}

	if k.bankKeeper.IsSendEnabledCoin(ctx, sdk.NewCoin(types.DefaultBondDenom, sdk.OneInt())) == false {
		return &types.MsgPrepayResponse{}, nil
	}

	sender, err := sdk.AccAddressFromBech32(msg.GetSender())
	if err != nil {
		return &types.MsgPrepayResponse{}, nil
	}

	purchased, err := k.Prepay(ctx, sender, msg.Coins)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypePrepay,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSender()),
			sdk.NewAttribute(types.AttributeKeyCoins, msg.Coins.String()),
			sdk.NewAttribute(types.AttributeKeyPurchasedUoz, purchased.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.GetSender()),
		),
	})

	return &types.MsgPrepayResponse{}, nil
}