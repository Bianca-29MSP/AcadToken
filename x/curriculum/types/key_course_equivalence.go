package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CourseEquivalenceKeyPrefix is the prefix to retrieve all CourseEquivalence
	CourseEquivalenceKeyPrefix = "CourseEquivalence/value/"
)

// CourseEquivalenceKey returns the store key to retrieve a CourseEquivalence from the index fields
func CourseEquivalenceKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
