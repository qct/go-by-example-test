package main

import (
    "os"
    "fmt"
)

func createFile(p string) *os.File {
    fmt.Println("creating...")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing...")
    fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
    fmt.Println("closing...")
    f.Close()
}

func main() {
    f := createFile("D:\\my-go\\src\\test-go\\defer.txt")
    defer closeFile(f)
    writeFile(f)
}
