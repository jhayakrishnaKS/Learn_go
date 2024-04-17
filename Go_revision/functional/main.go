package main

import "fmt"

func mapSlice(slice []int, fn func(int) int) []int {
    result := make([]int, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func filterSlice(slice []int, fn func(int) bool) []int {
    var result []int
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

func reduceSlice(slice []int, fn func(int, int) int, initialValue int) int {
    result := initialValue
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}

    doubled := mapSlice(nums, func(x int) int {
        return x * 2
    })
    fmt.Println("Doubled:", doubled) 

    evenOnly := filterSlice(nums, func(x int) bool {
        return x%2 == 0
    })
    fmt.Println("Even Only:", evenOnly) 

    sum := reduceSlice(nums, func(acc, x int) int {
        return acc + x
    }, 0)
    fmt.Println("Sum:", sum) 
}