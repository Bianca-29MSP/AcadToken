package types

// CourseNftKeyPrefix is the prefix for storing CourseNft
var CourseNftKeyPrefix = []byte{0x00}

// CourseNftKey returns the store key for a CourseNft
func CourseNftKey(
    nftId string,
) []byte {
    var key []byte
    
    key = append(key, CourseNftKeyPrefix...)
    key = append(key, []byte(nftId)...)
    
    return key
}