package main

import "fmt"

//冒泡排序
func selectSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		fmt.Println(arr)
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5}
	fmt.Println(selectSort(arr))
}
