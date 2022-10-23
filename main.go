package main

import "fmt"

// min implements a three way sort, assuming there are three possible values in
// the input, it sorts them in order by iterating over the array with three cursors:
// low, mid, and high. As we iterate, if the item is smaller than the pivot, it is
// swapped with the item at the low cursor and both the low and mid cursors are advanced.
// If the item is equal to the pivot, we simply advance the mid cursor. Finally, if the
// item is greater than the pivot, the item is swapped with the item at the high cursor and
// the high cursor is decreased. Once the mid cursor equals the high cursor, we stop.
func main() {
	input := []int{1, 4, 2, 4, 2, 4, 1, 2, 4, 1, 2, 2, 2, 2, 4, 1, 4, 4, 4}
	low := 0
	mid := 0
	high := len(input) - 1
	pivot := 2
	var swap int
	for mid < high {
		if input[mid] < pivot {
			// swap to beginning
			swap = input[low]
			input[low] = input[mid]
			input[mid] = swap
			low++
			mid++
		} else if input[mid] > pivot {
			// swap to end
			swap = input[high]
			input[high] = input[mid]
			input[mid] = swap
			high--
		} else {
			mid++
		}
		fmt.Printf("%v\n", input)
	}

	fmt.Printf("%v\n", input)
}
