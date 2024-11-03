package types

import (
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

// TxData implements the Ethereum transaction tx structure. It is used
// solely as intended in Ethereum abiding by the protocol.
type TxData interface {
	// TODO: embed ethtypes.TxData. See https://github.com/ethereum/go-ethereum/issues/23154

	TxType() byte
	Copy() TxData
	GetChainID() *big.Int
	GetAccessList() ethtypes.AccessList
	GetData() []byte
	GetNonce() uint64
	GetGas() uint64
	GetGasPrice() *big.Int
	GetGasTipCap() *big.Int
	GetGasFeeCap() *big.Int
	GetValue() *big.Int
	GetTo() *common.Address

	GetRawSignatureValues() (v, r, s *big.Int)
	SetSignatureValues(chainID, v, r, s *big.Int)

	AsEthereumData() ethtypes.TxData
	Validate() error

	// static fee
	Fee() *big.Int
	Cost() *big.Int

	// effective gasPrice/fee/cost according to current base fee
	EffectiveGasPrice(baseFee *big.Int) *big.Int
	EffectiveFee(baseFee *big.Int) *big.Int
	EffectiveCost(baseFee *big.Int) *big.Int
}
