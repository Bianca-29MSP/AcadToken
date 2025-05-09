package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CourseContentKeyPrefix is the prefix to retrieve all CourseContent
	CourseContentKeyPrefix = "CourseContent/value/"
)

// CourseContentKey returns the store key to retrieve a CourseContent from the index fields
func CourseContentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
