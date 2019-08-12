package main

import (
	"fmt"
)


func quick_sort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	middle := arr[0]
	var left []int
	var right []int

	// put the number to left or right
	for i:=1; i<len(arr); i++ {
		if middle < arr[i] {
			right = append(right,arr[i])
		} else {
			left = append(left,arr[i])
		}
	}


	left = quick_sort(left)
	right = quick_sort(right)

	left = append(left,middle)
	left = append(left,right...)

	return 	left


}



func main(){

	// slice
	arr := []int{25,133,452,364,5889,293,607,365,8633,555,18,66}
	result := quick_sort(arr)
	fmt.Print(result)
	fmt.Println("\ndone.\n")


}
