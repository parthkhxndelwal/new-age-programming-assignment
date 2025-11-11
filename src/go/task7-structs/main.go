package main

import "fmt"

type course struct {
    title string
    credits int
}

type Catalog struct {
    items []course
}

func NewCatalog() Catalog {
    return Catalog{items: []course{}}
}

func (c *Catalog) Add(title string, credits int) {
    c.items = append(c.items, course{title: title, credits: credits})
}

func (c Catalog) TotalCredits() int {
    total := 0
    for _, item := range c.items {
        total += item.credits
    }
    return total
}

func main() {
    catalog := NewCatalog()
    catalog.Add("Go Fundamentals", 3)
    catalog.Add("Kotlin Intro", 2)
    fmt.Println("Total credits:", catalog.TotalCredits())
}
