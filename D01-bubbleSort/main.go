package main

import "fmt"

//冒泡排序
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		fmt.Println(arr)
	}
	return arr
}

func main() {
	arr := []int{1, 9, 2, 8, 3, 7, 6, 4, 5}
	fmt.Println("初始数组：", arr)
	fmt.Println(bubbleSort(arr))
}
