package main

import (
	"fmt"
)

type transport interface{
	move(x int);
	signal(s string);
	getSpeed() int;
}

type Car struct{
	speed int;
}

type Airplane struct{
	speed int;
}

func (c *Car) move(x int){
	c.speed = x;
	fmt.Printf("car move with speed %v\n", x);
}

func (c Car) signal(s string){
	fmt.Println(s);
}

func (c Car) getSpeed() int{
	return c.speed;
}

func (a *Airplane) move(x int){
	a.speed = x
	fmt.Printf("airplane move with speed %v\n", x);
}

func (a Airplane) signal(s string){
	fmt.Println(s);
}

func (a Airplane) getSpeed() int{
	return a.speed;
}

func getMove(t transport){
	fmt.Println(t.getSpeed());
}

func (c Car) stop(){
	fmt.Println("stop car");
}

func main(){
	car := &Car{};
	airplane :=  &Airplane{};
	car.move(64);
	airplane.move(1024);
	car.signal("bip-bip");
	airplane.signal("uuuuuuuuuu");
	getMove(car);
	getMove(airplane);
	car.stop();
}