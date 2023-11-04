package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	text := "text from Golang programm";
	file, err := os.Create("text.txt");

	if err != nil {
		fmt.Println("error: ", err);
		os.Exit(1);
	}
	// defer file.Close();

	file.WriteString(text);

	file.Close();

	ffile, _ := os.OpenFile("text.txt", os.O_RDONLY, 0666);

	data := make([]byte, 64);

	// file.ReadString(data);

	for {
        n, er := ffile.Read(data);
        if er == io.EOF {   // если конец файла
            break;           // выходим из цикла
        }
        fmt.Print(string(data[:n]));
    }

	fmt.Printf("\ndone\n");
}