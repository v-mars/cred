package cred

type CredManager interface {
	GetHashedPassword(password string, userSalt string, organizationSalt string) string
	IsPasswordCorrect(password string, passwordHash string, userSalt string, organizationSalt string) bool
}

func GetCredManager(passwordType string) CredManager {
	if passwordType == "plain" {
		return NewPlainCredManager()
	} else if passwordType == "salt" {
		return NewSha256SaltCredManager()
	} else if passwordType == "md5-salt" {
		return NewMd5UserSaltCredManager()
	} else if passwordType == "bcrypt" {
		return NewBcryptCredManager()
	} else if passwordType == "pbkdf2-salt" {
		return NewPbkdf2SaltCredManager()
	} else if passwordType == "argon2id" {
		return NewArgon2idCredManager()
	}
	return nil
}
