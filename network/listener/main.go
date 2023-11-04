package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)
func main() { 
    // message := "Hello, I am a server";
	var message string;
    listener, err := net.Listen("tcp", ":4545"); 
	read := bufio.NewReader(os.Stdin);
     
    if err != nil {
        fmt.Println(err); 
        return;
    } 
    defer listener.Close();
    fmt.Println("Server is listening...");
    for {
		// fmt.Fscan(os.Stdin, &message);
		message, _ = read.ReadString('\n');
        conn, err := listener.Accept(); 
        if err != nil { 
            fmt.Println(err); 
            return;
        } 
        conn.Write([]byte(message));
        conn.Close();
    } 
}