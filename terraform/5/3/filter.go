package main

import (
    "fmt"
    "bytes"
    "strconv"
)

func main() {

    result := FilterBy3(1, 100)
    fmt.Println(result)
}


func FilterBy3(start int, end int) string {

    result := bytes.Buffer{}
    for i := start; i <= end; i++ {
	    if i % 3 == 0 {
            result.WriteString(strconv.Itoa(i) + " ")
	    }
    }
    return result.String()
}
