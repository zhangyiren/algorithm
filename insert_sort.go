package main

import (
	"fmt"
	"reflect"
)


func main(){

	/* insert sort */

	array := []int{25,133,452,364,5889,293,607,365,8633,555,18,66,78,15,43,67,96,155,387,972,1002,74,38,633}
	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))

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

	fmt.Println(array)

}
