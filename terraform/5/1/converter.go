package main

import "fmt"

func main() {

    var m float64 = 1

    for m != 0 {
        fmt.Print("Enter meters to convert (or 0 for exit): ")
        fmt.Scanf("%f", &m)
        foots := Converter(m)
        fmt.Println(m, "m = ", foots, "foots")
    }
   fmt.Println("Bye-bye!")
}

func Converter(m float64) float64 {

    return  m / 0.3048
}

