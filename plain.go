package cred

type PlainCredManager struct{}

func NewPlainCredManager() *PlainCredManager {
	cm := &PlainCredManager{}
	return cm
}

func (cm *PlainCredManager) GetHashedPassword(password string, userSalt string, organizationSalt string) string {
	return password
}

func (cm *PlainCredManager) IsPasswordCorrect(plainPwd string, hashedPwd string, userSalt string, organizationSalt string) bool {
	return hashedPwd == plainPwd
}
