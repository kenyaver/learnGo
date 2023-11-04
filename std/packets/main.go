package main

import (
	"./fun"
	"fmt"
)

func main(){
	var x, y int;
	fmt.Scan(&x, &y);
	fmt.Println(fun.Sum(x, y));
	fmt.Println(fun.Div(x, y));
}