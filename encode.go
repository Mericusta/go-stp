package stp

import (
	"crypto/md5"
	"encoding/hex"
)

func HashCode(b []byte) string {
	hashStr := md5.Sum(b)
	return hex.EncodeToString(hashStr[:])
}
