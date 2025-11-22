# Concept Notebook: Modern Programming Languages

## Go Overview
Go delivers lightweight concurrency, simple syntax, and a batteries-included tooling suite.
```go
package main
import "fmt"
func main() {
    message := make(chan string)
    go func() { message <- "Hello from a goroutine" }()
    fmt.Println(<-message)
}
```

## F# Overview
F# favors functional pipelines on .NET with strong type inference.
```fsharp
let squares = [1..5] |> List.map (fun n -> n * n)
printfn "%A" squares
```

## Clojure Overview
Clojure embraces immutable data and macros on the JVM.
```clojure
(def numbers [1 2 3 4])
(println (reduce + numbers))
```

## Kotlin Overview
Kotlin blends object-oriented and functional constructs for JVM and Android.
```kotlin
fun main() {
    val items = listOf("Go", "F#", "Clojure", "Kotlin")
    items.filter { it.length > 2 }.forEach(::println)
}
```

## Key Takeaways
- Each language targets safer concurrency and composability compared to legacy ecosystems.
- Toolchains focus on developer experience with modern package managers and REPLs.
- Interop layers (Go C-ABI, Kotlin JVM, F#/.NET, Clojure/Java) ease gradual adoption.

## Task Map
- Task 1 – Concept exploration: `docs/concept-notebook.md`, `docs/comparison-table.md`
- Task 2 – Environment setup: `screenshots/` directory
- Task 3 – Syntax and data types: `src/go/task3-syntax-datatypes`
- Task 4 – Control structures: `src/go/task4-control-structures`
- Task 5 – Functions and packages: `src/go/task5-functions`
- Task 6 – Arrays, slices, maps: `src/go/task6-arrays-slices-maps`
- Task 7 – Structs and methods: `src/go/task7-structs`
- Task 8 – Pointers and memory: `src/go/task8-pointers-memory`
