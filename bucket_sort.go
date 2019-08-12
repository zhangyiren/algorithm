package main

import (
	"fmt"
	"math"
)


/* insert sort */
func insert_sort(array []int) []int {

	var (
		key int
		tmp int
	)

	for i := 1; i < len(array); i++ {
		tmp = array[i]
		key = i-1

		for {
			if key >= 0 && tmp < array[key] {
				array[key+1] = array[key]
				key--
			} else {
				break
			}
		}

		if (key + 1) != i {
			array[key+1] = tmp
		}
	}

	return array

}


func main(){

	/* Bucket sort */

	array := []int{25,133,452,364,5889,293,607,365,8633,555,18,66,78,15,43,67,96,155,387,972,1002,74,38,633}

	var (
		max int
		min int
		arrayLength int
	)

	min = array[0]
	arrayLength = len(array)

	// find max and min number in one traversal
	for i := 0 ; i < arrayLength ; i++ {

		if array[i] > max {
			max = array[i]
		}

		if array[i] < min {
			min = array[i]
		}

	}


	// bucket number
	bucketNumber := 24

	var result [][]int = make([][]int, bucketNumber)
	var output []int = make([]int, arrayLength)

	// init big slice
	for i := 0 ; i < bucketNumber ; i++ {
		result[i] = make([]int,0)
	}


	stepSize := math.Floor(float64((max - min + 1) / bucketNumber))

	// traverse array again and allot positions for array numbers
	for i :=0 ; i < arrayLength ; i++ {

		pos := int(math.Floor(float64((array[i] - min + 1)) / stepSize))
		// prevent overflow
		if pos >= bucketNumber {
			pos = bucketNumber - 1
		}
		result[pos] = append(result[pos],array[i])
		tmp := insert_sort(result[pos])
		result[pos] = tmp

	}


	var k int

	for i:=0 ; i < bucketNumber ; i++ {

		eachPiceLen := len(result[i])
		if eachPiceLen > 0 {
			for j:=0 ; j < eachPiceLen ; j++ {
				output[k] = result[i][j]
				k++
			}

		}

	}

	fmt.Println(result)
	fmt.Println(output)


}
