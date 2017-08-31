package main

import (
    "sort"
    "fmt"
)

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}