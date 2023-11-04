package main

import (
	"fmt"
	"math" 
)

type Vertex struct{
	x, y float64 // типичный сишный вариант созданиия структуры
}

func (v Vertex) abs() float64{
	return math.Sqrt(v.x * v.x + v.y * v.y) // это метод для структуры.. странно, но привыкнуть возможно
}

func swap(value1, value2 int) (int, int){
	return value2, value1 // поддерживается множественныйвозврат
}

func main(){
	v := Vertex{3, 4}
	fmt.Println(v.abs())

	value1, value2 := swap(3, 4)
	fmt.Println("value1: ", value1)
	fmt.Println("value2: ", value2)

}