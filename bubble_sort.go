package main

import (
	"fmt"
)



func main(){

	/* bubble sort */
	var array [10] int = [...]int{10,87,13,9,26,374,198,55,63,70}
	for key,value := range(array){
		fmt.Println(key,value)
	}

	fmt.Println()

	var k = len(array)
	for i:= 0; i<k; i++ {
		for j:=0; j<k-1; j++ {
			if array[i] > array[j] {
				array[i],array[j] = array[j],array[i]
			}
		}
	}

	for key,value := range(array){
		fmt.Println(key,value)
	}

}
