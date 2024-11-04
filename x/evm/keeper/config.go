package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"evmos/x/evm/statedb"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/big"
)

// EVMConfig creates the EVMConfig based on current state
func (k *Keeper) EVMConfig(ctx sdk.Context, proposerAddress sdk.ConsAddress, chainID *big.Int) (*statedb.EVMConfig, error) {
	params := k.GetParams(ctx)
	ethCfg := params.ChainConfig.EthereumConfig(chainID)

	// get the coinbase address from the block proposer
	coinbase, err := k.GetCoinbaseAddress(ctx, proposerAddress)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to obtain coinbase address")
	}

	baseFee := k.GetBaseFee(ctx, ethCfg)
	return &statedb.EVMConfig{
		Params:      params,
		ChainConfig: ethCfg,
		CoinBase:    coinbase,
		BaseFee:     baseFee,
	}, nil
}
