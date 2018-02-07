package main

import (
	"fmt"
)

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}


func bubblesort(toSort []int) {
	for i := 1; i < len(toSort); i++ {
		for j := 0; j < len(toSort) - i; j++ {
			if toSort[j] > toSort[j+1] {
				swap(toSort, j, j+1)
			}
		}
	}
}



func main() {
	arr := []int{2, 5, 1, 3, 4}
	fmt.Println(arr)
	bubblesort(arr)
	fmt.Println(arr)
}
