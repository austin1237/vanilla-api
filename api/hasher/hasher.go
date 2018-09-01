package hasher

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"io"
)

func HashString(pword string) string {
	hash := sha512.New()
	io.WriteString(hash, pword)
	sum := hash.Sum(nil)
	base64 := b64.StdEncoding.EncodeToString(sum)
	return base64
}
