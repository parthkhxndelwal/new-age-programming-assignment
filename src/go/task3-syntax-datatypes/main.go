package main

import (
    "fmt"
    "math"
)

const course = "ENSP415"

func main() {
    var enrolled int = 42
    var passing float64 = 38.5
    language := "Go"
    inferred := 3.14

    var unsigned uint8 = 200
    var converted = int(unsigned)

    // Round the passing ratio to two decimals for report readability.
    ratio := passing / float64(enrolled)
    rounded := math.Round(ratio*100) / 100

    fmt.Println("Course:", course)
    fmt.Println("Language:", language)
    fmt.Println("Inferred literal:", inferred)
    fmt.Println("Unsigned to int:", converted)
    fmt.Println("Passing ratio:", rounded)
}
