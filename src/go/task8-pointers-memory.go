package main

import "fmt"

type counter struct {
    value int
}

func increment(ptr *counter) {
    ptr.value++
}

func main() {
    c := counter{value: 5}
    fmt.Println("Initial:", c.value)
    increment(&c)
    fmt.Println("After increment:", c.value)

    ref := &c
    fmt.Println("Pointer address:", ref)

    // Go's garbage collector reclaims heap data once no references remain.
}
