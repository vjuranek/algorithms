package main

import (
	"fmt"
)

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func quicksort(toSort []int, left int, right int) {
	if left < right {
		m := left
		for i := left + 1; i <= right; i++ {
			if toSort[i] < toSort[left] {
				m++
				swap(toSort, m, i)
			}
		}
		swap(toSort, left, m)
		quicksort(toSort, left, m-1)
		quicksort(toSort, m+1, right)
	}
}

func main() {
	arr := []int{2, 5, 1, 3, 4}
	fmt.Println(arr)
	quicksort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}
