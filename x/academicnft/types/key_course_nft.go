package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CourseNftKeyPrefix is the prefix to retrieve all CourseNft
	CourseNftKeyPrefix = "CourseNft/value/"
)

// CourseNftKey returns the store key to retrieve a CourseNft from the index fields
func CourseNftKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
