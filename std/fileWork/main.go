package main

import (
	"fmt"
	"os"
)

func main() {
	_, er := os.Create("text.txt");

	if er != nil {
		fmt.Println("error: ", er);
		os.Exit(-1);
	}

	file, err := os.OpenFile("text.txt", os.O_WRONLY, 0666);
	defer file.Close();

	if err != nil {
		fmt.Println("error: ", err);
		os.Exit(1);
	}

	fmt.Println(file.Name());
}