package cred

import (
	"crypto/sha256"
	"encoding/base64"

	"golang.org/x/crypto/pbkdf2"
)

type Pbkdf2SaltCredManager struct{}

func NewPbkdf2SaltCredManager() *Pbkdf2SaltCredManager {
	cm := &Pbkdf2SaltCredManager{}
	return cm
}

func (cm *Pbkdf2SaltCredManager) GetHashedPassword(password string, userSalt string, organizationSalt string) string {
	// https://www.keycloak.org/docs/latest/server_admin/index.html#password-database-compromised
	decodedSalt, _ := base64.StdEncoding.DecodeString(userSalt)
	res := pbkdf2.Key([]byte(password), decodedSalt, 27500, 64, sha256.New)
	return base64.StdEncoding.EncodeToString(res)
}

func (cm *Pbkdf2SaltCredManager) IsPasswordCorrect(plainPwd string, hashedPwd string, userSalt string, organizationSalt string) bool {
	return hashedPwd == cm.GetHashedPassword(plainPwd, userSalt, organizationSalt)
}
