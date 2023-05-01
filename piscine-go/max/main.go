package main

import "fmt"

func Max(arr []int) int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr[len(arr)-1]
}

func main() {
	arrInt := []int{23, 123, 1, 1111, 55, 122}
	max := Max(arrInt)
	fmt.Println(max)
}
