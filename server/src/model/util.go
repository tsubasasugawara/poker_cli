package model

import (
	"crypto/sha256"
	"encoding/hex"
)

// ハッシュ化関数
func Hash(s string) string {
    r := sha256.Sum256([]byte(s))
    return hex.EncodeToString(r[:])
}
