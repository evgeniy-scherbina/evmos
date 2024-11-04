package keeper

import (
	"context"
	errorsmod "cosmossdk.io/errors"

	"evmos/x/evm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) EthereumTx(goCtx context.Context, msg *types.MsgEthereumTx) (*types.MsgEthereumTxResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	sender := msg.From
	tx := msg.AsTransaction()

	response, err := k.ApplyTransaction(ctx, tx)
	if err != nil {
		return nil, errorsmod.Wrap(err, "failed to apply transaction")
	}

	_, _ = sender, tx
	_ = response

	return &types.MsgEthereumTxResponse{}, nil
}
