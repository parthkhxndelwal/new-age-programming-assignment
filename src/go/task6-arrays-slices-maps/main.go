package main

import "fmt"

func main() {
    var fixed [3]string
    fixed[0] = "Go"
    fixed[1] = "F#"
    fixed[2] = "Kotlin"

    slice := []string{"Syntax", "Control"}
    slice = append(slice, "Memory")

    copied := make([]string, len(slice))
    copy(copied, slice)
    // Carve a view of the tracked slice to show slicing semantics.
    window := copied[1:]

    stats := map[string]int{"Go": 2009, "Kotlin": 2011}
    stats["Clojure"] = 2007

    fmt.Println("Array:", fixed)
    fmt.Println("Slice:", slice)
    fmt.Println("Copied slice:", copied)
    fmt.Println("Window:", window)
    fmt.Println("Map:", stats)
}
