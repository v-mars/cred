package cred

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5UserSaltCredManager struct{}

func getMd5(data []byte) []byte {
	hash := md5.Sum(data)
	return hash[:]
}

func getMd5HexDigest(s string) string {
	b := getMd5([]byte(s))
	res := hex.EncodeToString(b)
	return res
}

func NewMd5UserSaltCredManager() *Md5UserSaltCredManager {
	cm := &Md5UserSaltCredManager{}
	return cm
}

func (cm *Md5UserSaltCredManager) GetHashedPassword(password string, userSalt string, organizationSalt string) string {
	res := getMd5HexDigest(password)
	if userSalt != "" {
		res = getMd5HexDigest(res + userSalt)
	}
	return res
}

func (cm *Md5UserSaltCredManager) IsPasswordCorrect(plainPwd string, hashedPwd string, userSalt string, organizationSalt string) bool {
	return hashedPwd == cm.GetHashedPassword(plainPwd, userSalt, organizationSalt)
}
