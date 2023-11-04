package main

import "fmt"

func main(){
	var slice = []int {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};
	var sumx, sumy int; 
	for  i := range slice{
		if(i % 2 == 0){
			sumx += slice[i];
		} else {
			sumy += slice[i];
		}
	}

	fmt.Printf("sum of even numbers: %v\nsum of not even numbers: %v\n", sumx, sumy);
}