package keeper

import (
	"context"

	"evmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EthereumTx(goCtx context.Context, msg *types.MsgEthereumTx) (*types.MsgEthereumTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgEthereumTxResponse{}, nil
}
