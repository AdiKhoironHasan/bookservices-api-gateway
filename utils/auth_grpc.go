package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMACToken(secretKey string) string {
	key := []byte(secretKey)

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(secretKey))
	hmac := mac.Sum(nil)

	return hex.EncodeToString(hmac)
}
