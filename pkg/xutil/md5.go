package xutil

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5HexBytes(v []byte) string {
	hash := md5.Sum(v)
	return hex.EncodeToString(hash[:])
}

func MD5Hex(v string) string {
	return MD5HexBytes([]byte(v))
}
