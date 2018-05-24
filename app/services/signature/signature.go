package signature

import (
	"crypto/sha256"
	"encoding/hex"
)

// SignatureService is a service generates a random hash
type SignatureService struct{}

// New creates an instance of the service
func New() *SignatureService {
	return &SignatureService{}
}

// Create generates a hash of the string
func (ss *SignatureService) Create(str string) (hash string, err error) {
	hash = ss.hasher(str)

	return
}

// Verify checks if the signature of the string matches a previouslt generated hash
func (ss *SignatureService) Verify(str string, hash string) (isValid bool, err error) {
	if ss.hasher(str) == hash {
		isValid = true
	}

	return
}

// hasher hashes a string
func (ss *SignatureService) hasher(str string) string {
	h := sha256.New()

	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
