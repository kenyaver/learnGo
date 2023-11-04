package main

import (
	"fmt"
	"sync"
)

func counter(x int, ch chan bool, mut* sync.Mutex){
	// defer close(ch);
	mut.Lock();
	count := 0;
	for i := 1; i < 6; i++ {
		count++;
		fmt.Println(x, " - ", count);
	}
	mut.Unlock();
	ch <- true;
}

func main(){
	var mut sync.Mutex;
	ch := make(chan bool);
	for i := 1; i < 5; i++ {
		go counter(i, ch, &mut);
	}

	for i := 1; i < 5; i++ {
		<-ch;
	}
}