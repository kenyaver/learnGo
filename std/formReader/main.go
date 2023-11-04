package main

import (
	"fmt"
	"io"
)

type phoneNumder string

func (ph phoneNumder) Read(phone []byte) (int, error) {
	count := 0;
	for i := 0; i < len(ph); i++ {
		if ph[i] >= '0' && ph[i] <= '9' {
			phone[count] = ph[i];
			count++;
		}
	}

	if phone[0] == '7' {
		phone[0] = '8';
	}

	return count, io.EOF;
}

func main(){
	var phone1 phoneNumder = "+7 (903)290-95-66";
	phone2 := phoneNumder("8-991-270-84-34");

	buffer := make([]byte, len(phone1));
	phone1.Read(buffer);
	fmt.Println(string(buffer));

	buffer = make([]byte, len(phone2));
	phone2.Read(buffer);
	fmt.Println(string(buffer));

}