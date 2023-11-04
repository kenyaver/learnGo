package main

import (
	"fmt"
	"sync"
	"time"
)

func counter(id int, wg* sync.WaitGroup){
	defer wg.Done();
	fmt.Println("goroutine ", id, " was started");
	time.Sleep(2 * time.Second);
	fmt.Println("goroutine ", id, " was done!")
}

func main(){
	var wg sync.WaitGroup;
	wg.Add(2);
	for i := 0; i < 2; i++ {
		go counter(i, &wg);
	}

	wg.Wait();

	fmt.Println("the end");
}