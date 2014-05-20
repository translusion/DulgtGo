package main

// example usage: dulgt -password=foobar -len=12 amazon

import (
    "encoding/base64"
    "code.google.com/p/go.crypto/scrypt"
    "flag"
    "fmt"
    "strings"
    "strconv"
)

var username    = flag.String("username", "", "what you login with")
var masterpassword = flag.String("password", "", "easy to remember master password")
var pepper      = flag.String("secret", "", "long written down secret")
var keyLen      = flag.Int("len", 8, "length of generated password")
var series      = flag.Int("series", 0, "increment when site requires new password")

// recommended cost parameters for interactive login in 2009
var N           = flag.Int("N", 16384, "General cost.")
var r           = flag.Int("r", 8, "Memory cost relative to total.")
var p           = flag.Int("p", 1, "CPU cost relative to total.")

func b64_decode_len(n int) int { return n / 4 * 3 + 2 }

func main() {    
    flag.Parse()
    salt := []byte(flag.Arg(0))         // use this as domain, or really hard password

    seeds := []string{*username, *masterpassword, *pepper, strconv.Itoa(*series)}
    finalpassword := []byte(strings.Join(seeds, " "))
    
    // we could use keylen directly, but this shortens slightly the length
    // of the string we need to create.
    dklen := b64_decode_len(*keyLen) + 1
    dk, _ := scrypt.Key(finalpassword, salt, *N, *r, *p, dklen)
    derivedPassword := base64.StdEncoding.EncodeToString(dk)[:*keyLen]
    
    fmt.Println(derivedPassword)
}