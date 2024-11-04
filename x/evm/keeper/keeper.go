package keeper

import (
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"math/big"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"evmos/x/evm/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		bankKeeper types.BankKeeper
		// access historical headers for EVM state transition execution
		stakingKeeper types.StakingKeeper

		// chain ID number obtained from the context's chain id
		eip155ChainID *big.Int
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		bankKeeper:    bankKeeper,
		stakingKeeper: stakingKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetBaseFee returns current base fee, return values:
// - `nil`: london hardfork not enabled.
// - `0`: london hardfork enabled but feemarket is not enabled.
// - `n`: both london hardfork and feemarket are enabled.
func (k Keeper) GetBaseFee(ctx sdk.Context, ethCfg *params.ChainConfig) *big.Int {
	return k.getBaseFee(ctx, types.IsLondon(ethCfg, ctx.BlockHeight()))
}

func (k Keeper) getBaseFee(ctx sdk.Context, london bool) *big.Int {
	if !london {
		return nil
	}

	// TODO(yevhenii): implement feeMarket
	// return 0 if feemarket not enabled.
	baseFee := big.NewInt(0)

	return baseFee
}
