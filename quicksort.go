package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    input := os.Args[1:]
    var numbers []int64
    for _, e := range input {
        temp_value, _ := strconv.ParseInt(e, 10, 0)
        numbers = append(numbers, temp_value)
    }
    quicksort(numbers, 0, len(numbers)-1)
    fmt.Println(numbers)
}

func quicksort(numbers []int64, start_index, end_index int) {

    if start_index >= end_index {
        return
    }

    pivot_index := end_index
    pivot := numbers[pivot_index]
    left := start_index
    right := pivot_index - 1

    for {
        for numbers[left] < pivot && left < pivot_index {
            left++
        }

        for numbers[right] > pivot && right > left {
            right--
        }

        if left < right {
            numbers[left], numbers[right] = numbers[right], numbers[left]
        } else {
            break
        }
    }

    if left == right {
        numbers[left], numbers[pivot_index] = numbers[pivot_index], numbers[left]
        quicksort(numbers, start_index, left-1)
        quicksort(numbers, left+1, end_index)
    } else if left > right {
        quicksort(numbers, start_index, left-1)
    }

}
