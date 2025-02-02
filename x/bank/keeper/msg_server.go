package keeper

import (
	"context"
	// "fmt"
	// "log"
	// "math/big"

	"github.com/armon/go-metrics"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/ethclient"

	"github.com/cosmos/cosmos-sdk/telemetry"
	// "github.com/cosmos/cosmos-sdk/tools"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

// func IsChainOpen() bool {

// 	log.Println("INSIDE THE CHAIN OPEN FUNCTION\n")
// 	// Connect to the Ethereum node
// 	client, err := ethclient.Dial(tools.NodeURL)
// 	if err != nil {
// 		log.Fatal("Failed to connect to Ethereum node:", err)
// 	}
// 	defer client.Close()
// 	privateKey, err := crypto.HexToECDSA(tools.PrivateKeyHex)
// 	if err != nil {
// 		log.Fatal("Failed to load private key:", err)
// 	}

// 	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(tools.ChainID))
// 	if err != nil {
// 		log.Fatal("Failed to create transactor:", err)
// 	}

// 	// Load the contract
// 	contract, err := tools.NewOnlineServerMonitor(common.HexToAddress(tools.ContractAddress), client)
// 	if err != nil {
// 		log.Fatal("Failed to load contract:", err)
// 	}

// 	// Get the current online server count
// 	count, err := contract.GetOnlineServerCount(&bind.CallOpts{})
// 	if err != nil {
// 		log.Fatal("Failed to get online server count:", err)
// 	}
// 	log.Println("Current Online Server Count:", count)

// 	// Get the state variable that tracks if 1000 servers were ever reached
// 	hasReached1000, err := contract.Reached1000ServerCountValue(&bind.CallOpts{})
// 	if err != nil {
// 		log.Fatal("Failed to check if 1000 server count was reached:", err)
// 	}
// 	log.Println("Has the chain ever reached 1000 servers?:", hasReached1000)

// 	// If server count is below 1000, check if it has ever reached 1000 before
// 	if count.Cmp(big.NewInt(1000)) < 0 {
// 		if hasReached1000 {
// 			return false
// 		}
// 	}

// 	// If server count is 1000 or more and hasReached1000 is false, update the contract state
// 	if count.Cmp(big.NewInt(1000)) >= 0 && !hasReached1000 {

// 		tx, err := contract.Reached1000ServerCount(auth)
// 		if err != nil {
// 			log.Fatal("Failed to update Reached1000ServerCountValue:", err)
// 		}

// 		fmt.Println("Updated Reached1000ServerCountValue, transaction hash:", tx.Hash().Hex())
// 	}

// 	return false
// }


func IsAddressAllowed(address string, amounts sdk.Coins) bool {
	// TODO: Get the boolean value from the smart contract by sending the address and value and return it
	return true
}

func (k msgServer) Send(goCtx context.Context, msg *types.MsgSend) (*types.MsgSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// // Check if chain is open or not
	// if !IsChainOpen() {
	// 	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "chain is closed")
	// }

	// Check if sender address is allowed to send the amount
	if !IsAddressAllowed(msg.FromAddress, msg.Amount) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "sender address is not allowed to send the amount")
	}

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

	return &types.MsgSendResponse{}, nil
}

func (k msgServer) MultiSend(goCtx context.Context, msg *types.MsgMultiSend) (*types.MsgMultiSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// // Check if chain is open or not
	// if !IsChainOpen() {
	// 	return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "chain is closed")
	// }

	// NOTE: totalIn == totalOut should already have been checked
	for _, in := range msg.Inputs {
		if err := k.IsSendEnabledCoins(ctx, in.Coins...); err != nil {
			return nil, err
		}
	}

	for _, out := range msg.Outputs {
		accAddr := sdk.MustAccAddressFromBech32(out.Address)

		if k.BlockedAddr(accAddr) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not allowed to receive funds", out.Address)
		}
	}

	err := k.InputOutputCoins(ctx, msg.Inputs, msg.Outputs)
	if err != nil {
		return nil, err
	}

	return &types.MsgMultiSendResponse{}, nil
}

func (k msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if k.GetAuthority() != req.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), req.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.SetParams(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

func (k msgServer) SetSendEnabled(goCtx context.Context, msg *types.MsgSetSendEnabled) (*types.MsgSetSendEnabledResponse, error) {
	if k.GetAuthority() != msg.Authority {
		return nil, sdkerrors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.GetAuthority(), msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if len(msg.SendEnabled) > 0 {
		k.SetAllSendEnabled(ctx, msg.SendEnabled)
	}
	if len(msg.UseDefaultFor) > 0 {
		k.DeleteSendEnabled(ctx, msg.UseDefaultFor...)
	}

	return &types.MsgSetSendEnabledResponse{}, nil
}
