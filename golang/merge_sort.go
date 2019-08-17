package main

import (
	"fmt"
)



func merge(array1 []int , array2 []int) []int {


	len1 := len(array1)
	len2 := len(array2)

	if len1 == 0 && len2 == 0 {
		return []int{}
	}

	if len1 == 0 {
		return array2
	}

	if len2 == 0 {
		return array1
	}

	var k = 0
	array3 := make([]int,0)
	for i := 0 ; i < len1 ; i++ {
		for j := k ; j < len2; j++ {
			if array1[i] > array2[j] {
				array3 = append(array3,array2[j])
				k++
			} else {
				array3 = append(array3,array1[i])
				// reach the end
				if i == len1 - 1 {
					array3 = append(array3, array2[j:]...)
				}
				break
			}
		}
		if k == len2 {
			array3 = append(array3,array1[i])
		}
	}
	return array3

}

func main(){

	/* merge sort */
	array := []int{25,133,452,364,5889,293,607,365,8633,555,18,66,78,15,43,67,96,155,387,972,1002,74,38,633,801}

	var (
		array1 []int	// section of array
		array2 []int	// store the tmp combine result
		arrayLen int
	)

	size := 2	// each array section size
	arrayLen = len(array)

	for i := 0 ; i < arrayLen ; i++ {
		array1 = append(array1,array[i])
		if len(array1) == size {
			// swap element make it sequential
			if array1[0] > array1[1] {
				tmp := array1[0]
				array1[0] = array1[1]
				array1[1] = tmp
			}
			array2 = merge(array1,array2)
			array1 = []int{}
		}

	}

	// maybe some elements left, merge it too
	array2 = merge(array1,array2)
	fmt.Println(array2)

}
