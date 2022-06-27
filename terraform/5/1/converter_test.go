package main

import "testing"

func TestMain(t *testing.T) {

    var result float64

    result = Converter(0.3048)
    if result != 1 {
        t.Error("Expected 1, got ", result)
    }
}
