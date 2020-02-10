package utils

import "sort"

// BubbleSort - basic bubble sort function
// []int {9,8,7,6,5} => []int{5,6,7,8,9}
func BubbleSort(elements []int){
	keepRunning := true
	for keepRunning {
		keepRunning = false

		for i:= 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1]{
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunning = true
			}
		}
	}
}

// Sort uses bubble sort for up to 1000 elements and native go sort above
// results provided by the benchmark test
func Sort(elements []int) {
	if len(elements) < 1000 {
		BubbleSort(elements)
		return
	}

	sort.Ints(elements)
}