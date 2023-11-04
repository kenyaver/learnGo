package main

import (
	"fmt"
	// "io"
)

type phoneWriter struct { }

func (pw phoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil;
	}

	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]));
		}
	}
	fmt.Println();

	return len(bs), nil;
}

func main(){
	byte1 := []byte("+7(903) 290 95 66");
	byte2 := []byte("89912708434");

	writer := phoneWriter{};

	writer.Write(byte1);
	writer.Write(byte2);
}