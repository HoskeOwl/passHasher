package passhasher

import (
	"fmt"
)

type Hasher interface {
	Sum([]byte) []byte
}

type Salter = func(password, globalSalt, salt string) string

func SimpleSalter(password, globalSalt, salt string) string {
	return fmt.Sprintf("%v%v%v", password, globalSalt, salt)
}

type PasswordHasher struct {
	globalSalt string
	hasher     Hasher
	salter     Salter
}

func NewPasswordHasher(globalSalt string, hasher Hasher) *PasswordHasher {
	return &PasswordHasher{globalSalt: globalSalt, hasher: hasher, salter: SimpleSalter}
}

func NewPasswordHasherWithSalter(globalSalt string, hasher Hasher, salter Salter) *PasswordHasher {
	return &PasswordHasher{globalSalt: globalSalt, hasher: hasher, salter: salter}
}

func (a *PasswordHasher) Hash(password string) []byte {
	return a.hasher.Sum([]byte(a.salter(password, a.globalSalt, "")))
}

func (a *PasswordHasher) HashWithSalt(password, salt string) []byte {
	return a.hasher.Sum([]byte(a.salter(password, a.globalSalt, salt)))
}

func (a *PasswordHasher) StringHash(password string) string {
	return fmt.Sprintf("%x", a.hasher.Sum([]byte(a.salter(password, a.globalSalt, ""))))

}

func (a *PasswordHasher) StringHashWithSalt(password, salt string) string {
	return fmt.Sprintf("%x", a.hasher.Sum([]byte(a.salter(password, a.globalSalt, salt))))
}
