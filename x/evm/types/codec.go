package types

import (
	errorsmod "cosmossdk.io/errors"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEthereumTx{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// UnpackTxData unpacks an Any into a TxData. It returns an error if the
// client state can't be unpacked into a TxData.
func UnpackTxData(any *codectypes.Any) (TxData, error) {
	if any == nil {
		return nil, errorsmod.Wrap(errortypes.ErrUnpackAny, "protobuf Any message cannot be nil")
	}

	txData, ok := any.GetCachedValue().(TxData)
	if !ok {
		return nil, errorsmod.Wrapf(errortypes.ErrUnpackAny, "cannot unpack Any into TxData %T", any)
	}

	return txData, nil
}
