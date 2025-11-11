package main

import "fmt"

func main() {
    score := 87
    if score >= 90 {
        fmt.Println("Tier: A")
    } else if score >= 80 {
        fmt.Println("Tier: B")
    } else {
        fmt.Println("Tier: Review needed")
    }

    switch {
    case score >= 85:
        fmt.Println("Eligible for mentoring")
    case score >= 70:
        fmt.Println("Eligible for retry")
    default:
        fmt.Println("Needs advisor meeting")
    }

    attempts := 0
    for attempts < 3 {
        fmt.Println("Practice attempt", attempts+1)
        attempts++
    }

    topics := []string{"Syntax", "Control", "Packages"}
    for index, topic := range topics {
        fmt.Printf("%d -> %s\n", index, topic)
    }
}
