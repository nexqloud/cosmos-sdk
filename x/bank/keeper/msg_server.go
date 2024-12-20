package keeper

import (
	"context"
	"os"
	"strconv"

	"github.com/armon/go-metrics"
	"github.com/nexqloud/nxqconfig"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

const (
	devicesListApiUrl   = "https://api.nexqloud.net/license/devices_list_with_cloudscore"
	MinCloudDeviceCount = 1000
)

type DeviceList struct {
	Total int `json:"total"`
}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func isInWhiteList(addr string) bool {
	return addr != nxqconfig.MaintenanceWallet && addr != nxqconfig.Vault1Wallet && addr != nxqconfig.Vault2Wallet &&
		addr != nxqconfig.Vault3Wallet && addr != nxqconfig.Vault4Wallet && addr != nxqconfig.Vault5Wallet && addr != nxqconfig.GasCollector
}

func checkOnlineDevices() error {
	// Read the current devices count from $HOME/.nxqd/devices_count file
	file_path := os.Getenv("HOME") + "/.nxqd/devices_count"
	data, err := os.ReadFile(file_path)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Failed to read devices count")
	}

	// Parse the devices count
	devices_count, err := strconv.Atoi(string(data))
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Failed to parse devices count")
	}

	// If the devices count is less than MinCloudDeviceCount, return an error
	if devices_count < MinCloudDeviceCount {
		return sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Not enough devices are online")
	}

	return nil
}

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// if !isInWhiteList(msg.FromAddress) {
	// if err := checkOnlineDevices(); err != nil {
	// 	return nil, err
	// }

	if err := k.IsSendEnabledCoins(ctx, msg.Amount...); err != nil {
		return nil, err
	}

	from, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return nil, err
	}
	to, err := sdk.AccAddressFromBech32(msg.ToAddress)
	if err != nil {
		return nil, err
	}

	if k.BlockedAddr(to) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", msg.ToAddress)
	}

	err = k.SendCoins(ctx, from, to, msg.Amount)
	if err != nil {
		return nil, err
	}

	defer func() {
		for _, a := range msg.Amount {
			if a.Amount.IsInt64() {
				telemetry.SetGaugeWithLabels(
					[]string{"tx", "msg", "send"},
					float32(a.Amount.Int64()),
					[]metrics.Label{telemetry.NewLabel("denom", a.Denom)},
				)
			}
		}
	}()

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgSendResponse{}, nil
}

func (k msgServer) MultiSend(goCtx context.Context, msg *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// allInWhiteList := true
	// NOTE: totalIn == totalOut should already have been checked
	for _, in := range msg.Inputs {
		if err := k.IsSendEnabledCoins(ctx, in.Coins...); err != nil {
			return nil, err
		}
		// if !isInWhiteList(in.Address) {
		// 	allInWhiteList = false
		// }
	}

	// if !allInWhiteList {
	// 	if err := checkOnlineDevices(); err != nil {
	// 		return nil, err
	// 	}
	// }

	for _, out := range msg.Outputs {
		accAddr := sdk.MustAccAddressFromBech32(out.Address)

		if k.BlockedAddr(accAddr) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive transactions", out.Address)
		}
	}

	err := k.InputOutputCoins(ctx, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
		),
	)

	return &types.MsgMultiSendResponse{}, nil
}
