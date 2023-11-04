package main // определение пакета

import (
	"fmt" // обьявление либ 
)

var b = "Hello"

func main(){
	var a = "World!"
	fmt.Print(b, " ")
	fmt.Print(a) // ура! это тип хелло ворлд!
}