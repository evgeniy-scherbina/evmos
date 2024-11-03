package keeper

import (
	"context"

	"evmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EthereumTx(goCtx context.Context, msg *types.MsgEthereumTx) (*types.MsgEthereumTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	sender := msg.From
	tx := msg.AsTransaction()

	_, _ = sender, tx

	return &types.MsgEthereumTxResponse{}, nil
}
