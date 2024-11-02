package simulation

import (
	"math/rand"

	"evmos/x/evm/keeper"
	"evmos/x/evm/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgEthereumTx(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgEthereumTx{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the EthereumTx simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "EthereumTx simulation not implemented"), nil, nil
	}
}
