package leetcode

import (
	"testing"

	"github.com/luo/leetcode-kit/util"
)

// an example about table-driven testing
/*
func TestArea(t *testing.T) {

    areaTests := []struct {
        shape Shape
        want  float64
    }{
        {Rectangle{12, 6}, 72.0},
        {Circle{10}, 314.1592653589793},
    }

    for _, tt := range areaTests {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %.2f want %.2f", got, tt.want)
        }
    }

}
*/

func TestTwoSum(t *testing.T) {
    // use your solution 
    got := twoSum([]int{2, 7, 11, 15}, 9)
    // the answer you want
    want := []int{0, 1}

    if  util.IntSliceIsEqual(got, want) {
        t.Errorf("want '%v' but got '%v'", want, got)
    }
}
