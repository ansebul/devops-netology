package main
import "testing"

func TestMain(t *testing.T) {

    x := []int{31,69,68,82,63,110,20,11,13,27,19,997,17,}
    var result int

    result = GetMinimal(x)
    if result != 11 {
        t.Error("Expected 11, got ", result)
    }

}

