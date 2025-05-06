package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CourseTokenKeyPrefix is the prefix to retrieve all CourseToken
	CourseTokenKeyPrefix = "CourseToken/value/"
)

// CourseTokenKey returns the store key to retrieve a CourseToken from the index fields
func CourseTokenKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
