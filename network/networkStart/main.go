package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpReq := "GET / HTTP/1.1\n" + "Host: golang.org\n\n";

	conn, err := net.Dial("tcp", "golang.org:80");

	if err != nil {
		fmt.Println("error: ", err);
		os.Exit(1);
	}

	defer conn.Close();

	if _, err := conn.Write([]byte(httpReq)); err != nil {
		fmt.Println(err);
		os.Exit(2);
	}

	io.Copy(os.Stdout, conn);
	fmt.Println("done!");
}