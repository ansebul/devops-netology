package main

import "fmt"

func main() {

    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    min := GetMinimal(x)
    fmt.Println("Minimal value: ", min)
}


func GetMinimal(x []int) int {

    min := x[0]
    for _, next_element := range x[1:] {
        if min > next_element {
            min = next_element
        }
    }
    return min
}
