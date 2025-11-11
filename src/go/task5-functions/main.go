package main

import (
    "fmt"
    "modernlangs/task5-functions/mypackage"
)

func main() {
    scores := []int{88, 92, 73, 99}
    ordered := mypackage.Normalize(scores)
    avg := mypackage.Average(scores)
    fmt.Println("Ordered scores:", ordered)
    fmt.Printf("Average score: %.2f\n", avg)
}
