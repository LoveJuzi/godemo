package memory

import (
	"fmt"
	"testing"
)

func printArray(arr [1 << 10]int) {
	fmt.Println(arr)
}

func printArrayByptr(arr *[1 << 10]int) {
	fmt.Println(arr)
}

func Test_array(t *testing.T) {
	// arr := [1 << 31]int{}
	// printArrayByptr(&arr)
	// printArray(arr)
	// arr := make([]int, 1<<31)
	// t.Error(arr)
}

func TestModifyA(t *testing.T) {
	n := Node{a: 10, b: 11}
	fmt.Println(n)
	ModifyA(n)
	fmt.Println(n)

}
