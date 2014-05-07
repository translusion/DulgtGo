package main

import (
	"encoding/base64"
    "code.google.com/p/go.crypto/scrypt"
    "fmt"
)

// recommended cost parameters for interactive login in 2009
const N = 16384
const r = 8
const p = 1 

func main() {
    keyLen := 8
    password := []byte("some password")
    salt := []byte("some salt")
    dk, _ := scrypt.Key(password, salt, N, r, p, keyLen)
	str := base64.StdEncoding.EncodeToString(dk)
    
    fmt.Println(str)
}