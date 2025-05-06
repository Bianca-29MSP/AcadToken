package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// InstitutionKeyPrefix is the prefix to retrieve all Institution
	InstitutionKeyPrefix = "Institution/value/"
)

// InstitutionKey returns the store key to retrieve a Institution from the index fields
func InstitutionKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
