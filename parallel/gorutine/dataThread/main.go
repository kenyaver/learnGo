package main

import "fmt"

func factorial(max int, chanel chan int){
	defer close(chanel);
	result := 1;
	for i := 1; i < max; i++{
		result *= i;
		chanel <- result;
	}
	// fmt.Println("completed");
}

func main(){
	ch := make(chan int);

	go factorial(7, ch);

	for{
		v, op := <-ch;

		if !op {
			break;
		}

		fmt.Println(v);
	}

	// fmt.Println("completed");
}