package main
import (
    "fmt"
    "os"
    "net"
    "io"
)
func main() {
	for {
		conn, err := net.Dial("tcp", "127.0.0.1:4545");
		if err != nil { 
			fmt.Println(err);
			return;
		} 
    	io.Copy(os.Stdout, conn);
		conn.Close();
	}
}