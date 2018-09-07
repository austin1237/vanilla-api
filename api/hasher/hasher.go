package hasher

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"io"
)

// GenerateHash creates a new sha512 hash based on the string passed in
func GenerateHash(pword string) string {
	hash := sha512.New()
	io.WriteString(hash, pword)
	sum := hash.Sum(nil)
	base64 := b64.StdEncoding.EncodeToString(sum)
	return base64
}
