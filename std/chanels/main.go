package main

import "fmt"

func factorial(x int, ch chan struct{}, res map[int]int){
	defer close(ch);
	result := 1;
	for i := 1; i <= x; i++{
		result *= i;
		res[i] = result;
	}
}

func main(){
	result := make(map[int]int);
	cha := make(chan struct{});

	go factorial(7, cha, result);
	<-cha;

	for i, v := range result{
		fmt.Printf("%v - %v\n", i, v);
	}
}