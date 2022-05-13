package pot

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stratosnet/stratos-chain/x/pot/keeper"
	"github.com/stratosnet/stratos-chain/x/pot/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case *types.MsgVolumeReport:
			res, err := msgServer.HandleMsgVolumeReport(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgVolumeReport(ctx, k, msg)
		case *types.MsgWithdraw:
			res, err := msgServer.HandleMsgWithdraw(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgWithdraw(ctx, k, msg)
		case *types.MsgFoundationDeposit:
			res, err := msgServer.HandleMsgFoundationDeposit(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgFoundationDeposit(ctx, k, msg)
		case *types.MsgSlashingResourceNode:
			res, err := msgServer.HandleMsgSlashingResourceNode(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
			//return handleMsgSlashingResourceNode(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

//// Handle handleMsgVolumeReport.
//func handleMsgVolumeReport(ctx sdk.Context, k keeper.Keeper, msg types.MsgVolumeReport) (*sdk.Result, error) {
//	if !(k.IsSPNode(ctx, msg.Reporter)) {
//		errMsg := fmt.Sprint("Volume report is not sent by a superior peer")
//		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
//	}
//
//	// ensure epoch increment
//	lastEpoch := k.GetLastReportedEpoch(ctx)
//	if msg.Epoch.LTE(lastEpoch) {
//		e := sdkerrors.Wrapf(types.ErrMatureEpoch, "expected epoch should be greater than %s, got %s",
//			lastEpoch.String(), msg.Epoch.String())
//		return nil, e
//	}
//
//	// TODO: verify BLS signature
//
//	txBytes := ctx.TxBytes()
//	txhash := fmt.Sprintf("%X", tmhash.Sum(txBytes))
//
//	totalConsumedOzone, err := k.VolumeReport(ctx, msg.WalletVolumes, msg.Reporter, msg.Epoch, msg.ReportReference, txhash)
//	if err != nil {
//		return nil, err
//	}
//
//	ctx.EventManager().EmitEvents(sdk.Events{
//		sdk.NewEvent(
//			types.EventTypeVolumeReport,
//			sdk.NewAttribute(types.AttributeKeyTotalConsumedOzone, totalConsumedOzone.String()),
//			sdk.NewAttribute(types.AttributeKeyReportReference, hex.EncodeToString([]byte(msg.ReportReference))),
//			sdk.NewAttribute(types.AttributeKeyEpoch, msg.Epoch.String()),
//		),
//		sdk.NewEvent(
//			sdk.EventTypeMessage,
//			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
//			sdk.NewAttribute(sdk.AttributeKeySender, msg.ReporterOwner.String()),
//		),
//	})
//
//	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
//}

func handleMsgWithdraw(ctx sdk.Context, k keeper.Keeper, msg types.MsgWithdraw) (*sdk.Result, error) {
	err := k.Withdraw(ctx, msg.Amount, msg.WalletAddress, msg.TargetAddress)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeWithdraw,
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
			sdk.NewAttribute(types.AttributeKeyWalletAddress, msg.WalletAddress.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.WalletAddress.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgFoundationDeposit(ctx sdk.Context, k keeper.Keeper, msg types.MsgFoundationDeposit) (*sdk.Result, error) {
	err := k.FoundationDeposit(ctx, msg.Amount, msg.From)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFoundationDeposit,
			sdk.NewAttribute(types.AttributeKeyAmount, msg.Amount.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.From.String()),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSlashingResourceNode(ctx sdk.Context, k keeper.Keeper, msg types.MsgSlashingResourceNode) (*sdk.Result, error) {
	for _, reporter := range msg.Reporters {
		if !(k.IsSPNode(ctx, reporter)) {
			errMsg := fmt.Sprint("Slashing msg is not sent by a meta node")
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, errMsg)
		}
	}

	amt, nodeType, err := k.SlashingResourceNode(ctx, msg.NetworkAddress, msg.WalletAddress, msg.Slashing, msg.Suspend)
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSlashing,
			sdk.NewAttribute(types.AttributeKeyWalletAddress, msg.WalletAddress.String()),
			sdk.NewAttribute(types.AttributeKeyNodeP2PAddress, msg.NetworkAddress.String()),
			sdk.NewAttribute(types.AttributeKeyAmount, amt.String()),
			sdk.NewAttribute(types.AttributeKeySlashingNodeType, nodeType.String()),
			sdk.NewAttribute(types.AttributeKeyNodeSuspended, strconv.FormatBool(msg.Suspend)),
		),
	})
	return &sdk.Result{Events: ctx.EventManager().Events()}, err
}
