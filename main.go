package main

// example usage: dulgt -password=foobar -len=12 amazon

import (
	"encoding/base64"
    "code.google.com/p/go.crypto/scrypt"
    "flag"
    "fmt"
)

var passwordStr = flag.String("password", "", "master password")
var keyLen      = flag.Int("len", 8, "length of generated password")

// recommended cost parameters for interactive login in 2009
var N           = flag.Int("N", 16384, "General cost.")
var r           = flag.Int("r", 8, "Memory cost relative to total.")
var p           = flag.Int("p", 1, "CPU cost relative to total.")

func main() {    
    flag.Parse()
    password := []byte(*passwordStr)
    salt := []byte(flag.Arg(0))         // use this as domain, or really hard password
    
    //to debug that we actually got a value for password
    //fmt.Printf("password:%s, salt:%s, keylen: %d\n\n", *passwordStr, flag.Arg(0), *keyLen)
    
    dk, _ := scrypt.Key(password, salt, *N, *r, *p, *keyLen)
	str := base64.StdEncoding.EncodeToString(dk)
    
    fmt.Println(str)
}