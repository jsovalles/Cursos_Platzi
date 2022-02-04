package main

import (
    "fmt"
    "crypto/sha256"
    "encoding/hex"
)

func main() {

    s := ""
    h := sha256.New()
    h.Write([]byte(s))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

    fmt.Println(s, sha1_hash)
}