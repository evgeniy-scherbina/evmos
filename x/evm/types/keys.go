package types

const (
	// ModuleName defines the module name
	ModuleName = "evm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_evm"
)

var (
	ParamsKey = []byte("p_evm")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
