package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// AcademicTreeKeyPrefix is the prefix to retrieve all AcademicTree
	AcademicTreeKeyPrefix = "AcademicTree/value/"
)

// AcademicTreeKey returns the store key to retrieve an AcademicTree from the index fields
func AcademicTreeKey(
	student string,
) []byte {
	var key []byte

	studentBytes := []byte(student)
	key = append(key, studentBytes...)
	key = append(key, []byte("/")...)

	return key
}
