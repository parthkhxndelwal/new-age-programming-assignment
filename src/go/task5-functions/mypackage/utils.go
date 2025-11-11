package mypackage

import "sort"

func Normalize(scores []int) []int {
    copied := append([]int(nil), scores...)
    sort.Ints(copied)
    return copied
}

func Average(scores []int) float64 {
    if len(scores) == 0 {
        return 0
    }
    total := 0
    for _, score := range scores {
        total += score
    }
    return float64(total) / float64(len(scores))
}
