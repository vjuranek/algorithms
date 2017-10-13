package main

import (
	"fmt"
)

func merge(left []int, right []int) []int {
	merged := make([]int, 0, len(left) + len(right)) //(Type, lenght, capacity)
	//merged := []int{}
	for i,j := 0, 0; i < len(left) && j < len(right);  {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
		if i == len(left) {
			return append(merged, right[j:]...)
		}
		if j == len(right) {
			return append(merged, left[i:]...)
		}
	}
	return merged
}

func mergesort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr)/2
	arr1 := mergesort(arr[:middle])
	arr2 := mergesort(arr[middle:])
	return merge(arr1, arr2)
}

func main() {
	arr := []int{2, 5, 1, 3, 4}
	fmt.Println(arr)
	arr = mergesort(arr)
	fmt.Println(arr)
}