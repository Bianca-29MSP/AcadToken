package types

const (
	// ModuleName defines the module name
	ModuleName = "curriculum"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_curriculum"
)

var (
	ParamsKey = []byte("p_curriculum")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
