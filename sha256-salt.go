package cred

import (
	"crypto/sha256"
	"encoding/hex"
)

type Sha256SaltCredManager struct{}

func getSha256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func getSha256HexDigest(s string) string {
	b := getSha256([]byte(s))
	res := hex.EncodeToString(b)
	return res
}

func NewSha256SaltCredManager() *Sha256SaltCredManager {
	cm := &Sha256SaltCredManager{}
	return cm
}

func (cm *Sha256SaltCredManager) GetHashedPassword(password string, userSalt string, organizationSalt string) string {
	res := getSha256HexDigest(password)
	if organizationSalt != "" {
		res = getSha256HexDigest(res + organizationSalt)
	}
	return res
}

func (cm *Sha256SaltCredManager) IsPasswordCorrect(plainPwd string, hashedPwd string, userSalt string, organizationSalt string) bool {
	return hashedPwd == cm.GetHashedPassword(plainPwd, userSalt, organizationSalt)
}
