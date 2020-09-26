package hw2

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	inputArray := []int{4,2,5,1,2,6,2}
	MergeSort(inputArray, 0, len(inputArray)-1)
	expected := []int{1, 2, 2, 2, 4, 5, 6}

	if !reflect.DeepEqual(inputArray, expected) {
		t.Errorf("expected '%q' but got '%q'", expected, inputArray)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	inputArray := []int{1,2,321423513241,1234,23,423,1,53,2,12521,42442,213,2,23,412224,4326,2,41,32,1543,632,24,21,12645134,1523323,23,14,21,34,412,324,21,1,22,314,2351,23,122}
	for i := 0; i<b.N; i++ {
		MergeSort(inputArray, 0, len(inputArray)-1)
	}
}

func ExampleMergeSort_0() {
	inputArray := []int{}
	MergeSort(inputArray, 0, len(inputArray)-1)
	fmt.Printf("%v", inputArray)
	// Output: []
}

func ExampleMergeSort_10() {
	inputArray := []int{15,2,3,55,24,21,34,11,3415,34}
	MergeSort(inputArray, 0, len(inputArray)-1)
	fmt.Printf("%v", inputArray)
	// Output: [2 3 11 15 21 24 34 34 55 3415]
}

