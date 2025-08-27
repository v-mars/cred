package cred

import (
	"fmt"
	"testing"
)

func TestGetSaltedPassword(t *testing.T) {
	password := "123456"
	salt := "123"
	cm := NewSha256SaltCredManager()
	fmt.Printf("%s -> %s\n", password, cm.GetHashedPassword(password, "", salt))
}

func TestGetPassword(t *testing.T) {
	password := "123456"
	cm := NewSha256SaltCredManager()
	// https://passwordsgenerator.net/sha256-hash-generator/
	fmt.Printf("%s -> %s\n", "8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92", cm.GetHashedPassword(password, "", ""))
}
