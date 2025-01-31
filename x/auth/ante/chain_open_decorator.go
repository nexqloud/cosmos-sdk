package ante

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/cosmos/cosmos-sdk/tools"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ChainOpenDecorator checks if the chain is open before processing a transaction.
type ChainOpenDecorator struct {
	// Add any dependencies you need (e.g., a keeper)
}

// NewChainOpenDecorator creates a new ChainOpenDecorator.
func NewChainOpenDecorator() ChainOpenDecorator {
	return ChainOpenDecorator{}
}

// IsChainOpen checks if the chain is open by querying an Ethereum smart contract.
func IsChainOpen() bool {
	log.Println("INSIDE THE CHAIN OPEN FUNCTION\n")
	// Connect to the Ethereum node
	client, err := ethclient.Dial(tools.NodeURL)
	if err != nil {
		// log.Fatal("Failed to connect to Ethereum node:", err)
		log.Println("Failed to connect to Ethereum node:", err)
		return false // Return false if the node is unavailable
	}
	defer client.Close()
	privateKey, err := crypto.HexToECDSA(tools.PrivateKeyHex)
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(tools.ChainID))
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	// Load the contract
	contract, err := tools.NewOnlineServerMonitor(common.HexToAddress(tools.ContractAddress), client)
	if err != nil {
		log.Fatal("Failed to load contract:", err)
	}

	// Get the current online server count
	count, err := contract.GetOnlineServerCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to get online server count:", err)
	}
	log.Println("Current Online Server Count:", count)

	// Get the state variable that tracks if 1000 servers were ever reached
	hasReached1000, err := contract.Reached1000ServerCountValue(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to check if 1000 server count was reached:", err)
	}
	log.Println("Has the chain ever reached 1000 servers?:", hasReached1000)

	// If server count is below 1000, check if it has ever reached 1000 before
	if count.Cmp(big.NewInt(1000)) < 0 {
		if hasReached1000 {
			return false
		}
	}

	// If server count is 1000 or more and hasReached1000 is false, update the contract state
	if count.Cmp(big.NewInt(1000)) >= 0 && !hasReached1000 {
		tx, err := contract.Reached1000ServerCount(auth)
		if err != nil {
			log.Fatal("Failed to update Reached1000ServerCountValue:", err)
		}

		log.Println("Updated Reached1000ServerCountValue, transaction hash:", tx.Hash().Hex())
	}

	return false
}

// AnteHandle implements the AnteDecorator interface.
func (cod ChainOpenDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	if !IsChainOpen() {
		return ctx, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "chain is closed")
	}

	// Call the next ante handler
	return next(ctx, tx, simulate)
}