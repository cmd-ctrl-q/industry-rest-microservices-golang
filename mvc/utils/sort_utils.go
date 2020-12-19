package utils

import "sort"

// BubbleSort ...
func bubbleSort(elements []int) []int {
	running := true
	for running {
		running = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				running = true
			}
		}
	}
	return elements
}

func Sort(els []int) []int {
	if els == nil {
		return nil
	}
	if len(els) <= 23500 {
		// call custom bubble sort
		bubbleSort(els)
		return els
	}
	// else call built-in bubble sort
	sort.Ints(els)
	return els
}
