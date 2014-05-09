package main

// example usage: dulgt -password=foobar -len=12 amazon

import (
	"encoding/base64"
    "code.google.com/p/go.crypto/scrypt"
    "fmt"
)

const keyLen = 8

// recommended cost parameters for interactive login in 2009
const N = 16384
const r = 8
const p = 1

func main() {    
    dk, _ := scrypt.Key([]byte("mypassword"), []byte("mysalt"), N, r, p, keyLen)
	str := base64.StdEncoding.EncodeToString(dk)
    
    fmt.Println(str)
}