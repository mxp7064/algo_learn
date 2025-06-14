package main

import "fmt"

func main() {
	fmt.Println(binarySearchRecursive([]int{5, 6, 33, 57}, 33))
}

func binarySearchRecursive(arr []int, target int) int {
	return bsHelper(arr, target, 0, len(arr)-1)
}

func bsHelper(arr []int, target, low, high int) int {
	if low >= high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == target {
		return mid
	}
	if target < arr[mid] {
		return bsHelper(arr, target, low, mid-1)
	} else {
		return bsHelper(arr, target, mid+1, high)
	}
}
