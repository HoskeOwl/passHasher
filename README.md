# passhasher
Password hasher library for go

Library for hashing passwords with different modules (like md5, sha1, sha2, e.t.c).

Provide signle point for hashing password with different algorythm. Two salt can be used.

Simple replacing to another algorythm.

Hash can be returned as pure bytes or as hex string.

Method of concatinationg password and salt can be redefined.

---

**Do not use default salt concatination in production!**

---

### Simple example:
```
package main

import (
	"fmt"

	"github.com/HoskeOwl/passhasher"
	"golang.org/x/crypto/sha3"
)

func main() {
	hash := sha3.New256()
	passwordHasher := passhasher.NewPasswordHasher("s", hash)
	fmt.Printf("password: %v\n", passwordHasher.StringHash("password"))
}
```

### Redefine salt concatination:
```
package main

import (
	"fmt"

	"github.com/HoskeOwl/passhasher"
	"golang.org/x/crypto/sha3"
)

func SaltConcatination(password, globalSalt, salt string) string{
    return fmt.Sprintf("%v:%v:%v", password, globalSalt, salt)
}

func main() {
	hash := sha3.New256()
	passwordHasher := passhasher.NewPasswordHasherWithSalter("s", hash, SaltConcatination)
	fmt.Printf("password: %v\n", passwordHasher.StringHash("password"))
}
```
