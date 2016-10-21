package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    // Get the numbers passed to the program
    input := os.Args[1:]
    var numbers []int64
    for _, e := range input {
        temp_value, _ := strconv.ParseInt(e, 10, 0)
        // Convert them to numbers
        numbers = append(numbers, temp_value)
    }
    // Quicksort this number slice
    // We will sort from the first element (index 0) to the last element (index len(numbers)-1)
    quicksort(numbers, 0, len(numbers)-1)
    // Print the sorted result
    fmt.Println(numbers)
}

// We will sort part of a slice, from start_index to end_index, both inclusive
func quicksort(numbers []int64, start_index, end_index int) {

    // It's an error if start_index > end_index
    // If start_index == end_index then we don't need to do anything
    // So we just return in both cases
    if start_index >= end_index {
        return
    }

    pivot_index := end_index // Choose last element as the pivot
    pivot := numbers[pivot_index]
    left := start_index      // First element as left
    right := pivot_index - 1 // Second to last element as right

    for {
        // Continuously move left pointer right
        // Left pointer can cross the right pointer and equals pivot_index
        for numbers[left] < pivot && left < pivot_index {
            left++
        }

        // Continuously move right pointer left
        // Right pointer can equal left pointer but not cross it
        for numbers[right] > pivot && right > left {
            right--
        }

        // If left pointer and right pointer are not equal, then exchange the elements
        if left < right {
            numbers[left], numbers[right] = numbers[right], numbers[left]
        } else {
            break
        }
    }

    // If left pointer and right pointer point to the same element, exchange it with the pivot
    // and sort left part and right part recursively
    if left == right {
        numbers[left], numbers[pivot_index] = numbers[pivot_index], numbers[left]
        quicksort(numbers, start_index, left-1)
        quicksort(numbers, left+1, end_index)
    } else if left > right {
        // If left pointer crossed right pointer, then just sort the left part of the slice recursively
        quicksort(numbers, start_index, left-1)
    }

}
