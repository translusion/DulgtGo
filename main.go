package main

// example usage: dulgt -password=foobar -len=12 amazon

import (
    "code.google.com/p/go.crypto/scrypt"
    "fmt"
)

type TestVector struct {
    P, S []byte
    N, r, p, dklen int
}

var testData = []TestVector {
    TestVector {[]byte(""), []byte(""), 16, 1, 1, 64},
    {[]byte("password"), []byte("NaCl"), 1024, 8, 16, 64},
    {[]byte("pleaseletmein"), []byte("SodiumChloride"), 16384, 8, 1, 64},
    {[]byte("pleaseletmein"), []byte("SodiumChloride"), 1048576, 8, 1, 64},
    {[]byte("mypassword"), []byte("mysalt"), 16384, 8, 1, 8},   
}


func main() {
    for _, t := range testData {
        dk, _ := scrypt.Key(t.P, t.S, t.N, t.r, t.p, t.dklen)
        for _, v := range dk {
            fmt.Printf("%x ", v)
        } 
        fmt.Println("\n")     
    }    
}