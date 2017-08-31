package main

import (
    "crypto/sha1"
    "fmt"
    "crypto/md5"
)

func main() {
    s := "sha1 this string"

    h := sha1.New()

    h.Write([]byte(s))

    bs := h.Sum(nil)

    fmt.Println(s)
    fmt.Printf("%x\n", bs)

    md5 := md5.New()
    md5.Write([]byte(s))
    sum := md5.Sum(nil)
    fmt.Printf("%x\n", sum)
}
