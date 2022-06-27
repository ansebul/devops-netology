package main

import  "testing"

func TestMain(t *testing.T) {

    var result string

    result = FilterBy3(3, 3)
    if result != "3 " {
        t.Error("Expected '3 ', got '", result, "'")
    }
}
