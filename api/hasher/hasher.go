package hasher

import (
	"crypto/sha512"
	b64 "encoding/base64"
	"io"
	"net/http"
)

func hashString(pword string) string {
	hash := sha512.New()
	io.WriteString(hash, pword)
	sum := hash.Sum(nil)
	base64 := b64.StdEncoding.EncodeToString(sum)
	return base64
}

func GenerateHash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// Validate form here
		userPassword := r.Form["password"][0]
		hashStr := hashString(userPassword)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(hashStr))
		next.ServeHTTP(w, r)
	})
}
