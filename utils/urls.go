package utils

import (
	"crypto/sha256"
	"encoding/hex"
)


func ShortenURL(longURL string) (string, error) {
	if len(longURL) == 0 {
		return "", &URLError{400, "Invalid long URL - Length is 0"}
	}
	hash := sha256.Sum256([]byte(longURL))
	shortURL := hex.EncodeToString(hash[:])[:8]
	return shortURL, nil
}
