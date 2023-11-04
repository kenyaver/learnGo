package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	read := bufio.NewReader(os.Stdin);
	var value, value1 string;
	var sign string;

	value, _ = read.ReadString(' ');
	value = strings.TrimSpace(value);

	sign, _ = read.ReadString(' ');
	sign = strings.TrimSpace(sign);

	value1, _ = read.ReadString('\n');
	value1 = strings.TrimSpace(value1);

	fmt.Printf("%q %q %q\n", value, sign, value1);

	var x, y int;
	x, _ = strconv.Atoi(value);
	y, _ = strconv.Atoi(value1);
	fmt.Printf("%d %q %d\n", x, sign, y);
}