package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkRim(value string) bool{
	flag := false;
	rim := map[string]string{"I": "1",
							"II": "2",
							"III": "3",
							"IV": "4",
							"V": "5",
							"VI": "6",
							"VII": "7",
							"VIII": "8",
							"IX": "9",
							"X": "10"};

	if _, ok := rim[value]; ok{
	flag = true;
	}
	return flag;
}

func calculate(x, y int, sign string) int{
	var result int;
	switch sign{
	case "+":
		result = x + y;
		break;
	case "-":
		result = x - y;
		break;
	case "*":
		result = x * y;
		break;
	case "/":
		result = x / y;
		break;
	default:
		result = 101;
		break;
	}
	return result;
}

func convert(x string) int{
	var xN int;

	switch x{
	case "1":
		xN, _ = strconv.Atoi(x);
		break;
	case "2":
		xN, _ = strconv.Atoi(x);
		break;
	case "3":
		xN, _ = strconv.Atoi(x);
		break;
	case "4":
		xN, _ = strconv.Atoi(x);
		break;
	case "5":
		xN, _ = strconv.Atoi(x);
		break;
	case "6":
		xN, _ = strconv.Atoi(x);
		break;
	case "7":
		xN, _ = strconv.Atoi(x);
		break;
	case "8":
		xN, _ = strconv.Atoi(x);
		break;
	case "9":
		xN, _ = strconv.Atoi(x);
		break;
	case "10":
		xN, _ = strconv.Atoi(x);
		break;
	default:
		xN = -1;
		break;
	}

	return xN;
}

func result(x, y, sign string, flag bool) string{
	var answer string; 
	var xN, yN int;
	var xy int;

	rim := map[string]string{
		"I": "1",
		"II": "2",
		"III": "3",
		"IV": "4",
		"V": "5",
		"VI": "6",
		"VII": "7",
		"VIII": "8",
		"IX": "9",
		"X": "10"};

	res := map[int]string{
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
		10: "X",
		11: "XI",
		12: "XII",
		13: "XIII",
		14: "XIV",
		15: "XV",
		16: "XVI",
		17: "XVII",
		18: "XVIII",
		19: "XIX",
		20: "XX",
		21: "XXI",
		24: "XXIV",
		25: "XXV",
		27: "XXVII",
		30: "XXX",
		32: "XXXII",
		35: "XXXV",
		36: "XXXVI",
		40: "XL",
		42: "XLII",
		45: "XLV",
		48: "XLVIII",
		49: "XLIX",
		50: "L",
		54: "LIV",
		56: "LVI",
		60: "LX",
		63: "LXII",
		64: "LXIV",
		70: "LXX",
		72: "LXXII",
		80: "LXXX",
		81: "LXXXI",
		90: "XC",
		100: "C"};


	if flag{
		if count, ok := rim[x]; ok{
			x = count;
		}
		if count, ok := rim[y]; ok{
			y = count;
		}
	}

	xN = convert(x);
	yN = convert(y);

	xy = calculate(xN, yN, sign);


	if xy != 101 {
		if xN != -1 && yN != -1 {
			if flag{
				if count, ok := res[xy]; ok{
					answer = count;
				} else {
					answer = "errora";
				}
			} else {
				answer = strconv.Itoa(xy);
			}
		} else {
			answer = "errorb";
		}
	} else {
		answer = "errorc";
	}

	return answer;
}

func main(){
	read := bufio.NewReader(os.Stdin);
	var sign string;

	fmt.Printf("Enter primer:\n");
		
	var value1, value2 string;
	// input values with Rim`s numbers support:
	
	// input value1:
	value1, _ = read.ReadString(' ');
	value1 = strings.TrimSpace(value1);
	
	// read str {+, -, *, /}
	sign, _ = read.ReadString(' ');
	sign = strings.TrimSpace(sign);
	
	// input value2:
	value2, _ = read.ReadString('\n');
	value2 = strings.TrimSpace(value2);

	var rim1, rim2 bool = false, false;
	rim1 = checkRim(value1);
	rim2 = checkRim(value2);

	if(rim1 == rim2){
		if res := result(value1, value2, sign, rim1); res != "errora" && res != "errorb" && res != "errorc" {
			fmt.Println(res);
		} else if res == "errora" {
			fmt.Printf("not valid answer\n");
		} else if res == "errorb" {
			fmt.Printf("not valid values\n");
		} else if res == "errorc" {
			fmt.Printf("not valid operation\n");
		}
	} else{
		fmt.Printf("not valid values\n");
	}
}