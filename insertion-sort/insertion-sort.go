package main

import "fmt"

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = key
	}
	fmt.Println(array)
}

func main() {
	insertionSort([]int{5, 2, 4, 6, 1, 3})
}
