# NEW AGE PROGRAMMING - Assignment

## Overview of Modern Languages
- **Go:** Compiled, statically typed, built-in concurrency with goroutines, simple syntax, strong standard library.
- **F#:** Functional-first on .NET, type inference, pattern matching, immutable defaults, interoperates with C#.
- **Clojure:** Lisp dialect on the JVM, emphasizes immutability, powerful macro system, seamless Java interop.
- **Kotlin:** Pragmatic JVM/Android language, concise syntax, null safety, coroutines, interoperable with Java.

## Comparison with Traditional Languages
Traditional languages like C, C++, and Java prioritize manual memory management or verbose syntax, while the modern set focuses on safer concurrency, succinct expression, and interoperability. Detailed comparisons live in `docs/comparison-table.md`.

## Environment Setup & Hello World Outputs
### Go
1. Download Go from https://go.dev/dl
2. Install using the system installer; accept defaults.
3. Verify with `go version`.
4. Run the sample with `go run ./src/go/task3-syntax-datatypes`.

Screenshots: `screenshots/go-setup.png`, `screenshots/go-hello-world.png`

### Kotlin
1. Install OpenJDK 21 from https://adoptium.net/
2. Install Kotlin via the Kotlin compiler bundle or IntelliJ IDEA.
3. Verify with `kotlinc -version`.
4. Compile and run `src/other-language/hello-world.kt` using `kotlinc` and `kotlin`.

Screenshots: `screenshots/other-language-setup.png`, `screenshots/other-language-hello-world.png`

## GO Basics Implemented
- Syntax, constants, and type conversion examples (`src/go/task3-syntax-datatypes`, `outputs/task3-output.txt`).
- Control structures covering `if`, `switch`, `for`, and `range` (`src/go/task4-control-structures`, `outputs/task4-output.txt`).
- Modular functions with custom package usage (`src/go/task5-functions`, `outputs/task5-output.txt`).
- Arrays, slices, maps, and slice operations (`src/go/task6-arrays-slices-maps`, `outputs/task6-output.txt`).
- Structs with methods illustrating encapsulation (`src/go/task7-structs`, `outputs/task7-output.txt`).
- Pointer usage and memory notes (`src/go/task8-pointers-memory`, `outputs/task8-output.txt`).

## Running Task Programs
- Task 3: `go run ./src/go/task3-syntax-datatypes`
- Task 4: `go run ./src/go/task4-control-structures`
- Task 5: `go run ./src/go/task5-functions`
- Task 6: `go run ./src/go/task6-arrays-slices-maps`
- Task 7: `go run ./src/go/task7-structs`
- Task 8: `go run ./src/go/task8-pointers-memory`

## Repository Layout
```
README.md
docs/
  concept-notebook.md
  comparison-table.md
screenshots/
  go-setup.png
  go-hello-world.png
  other-language-setup.png
  other-language-hello-world.png
src/
  go/
    task3-syntax-datatypes/
      main.go
    task4-control-structures/
      main.go
    task5-functions/
      main.go
      mypackage/utils.go
    task6-arrays-slices-maps/
      main.go
    task7-structs/
      main.go
    task8-pointers-memory/
      main.go
  other-language/
    hello-world.kt
outputs/
  task3-output.txt
  task4-output.txt
  task5-output.txt
  task6-output.txt
  task7-output.txt
  task8-output.txt
```

## Task Checklist
- [x] Task 1: Concept notebook with code snippets
- [x] Task 2: Installation documentation
- [x] Task 3: GO syntax and data types program
- [x] Task 4: Control structures implementation
- [x] Task 5: Functions and packages
- [x] Task 6: Arrays, slices, and maps
- [x] Task 7: Structs and methods
- [x] Task 8: Pointers and memory management
- [x] Task 9: GitHub-ready documentation
- [x] README.md completed with required sections
- [x] Screenshots placeholders added
- [x] Repository ready for submission
