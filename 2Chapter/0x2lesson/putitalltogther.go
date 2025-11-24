package main 

import (
	"fmt"
)
type MOVES interface{
	move()
}

type car struct{
	model string 
	make string
} 

func (c car) move(){
	fmt.Printf("model [%s] of the Car and Make is [%s]\n",c.model,c.make)
}

type bike struct{
	model string
}

func (b bike) move(){
	fmt.Printf("Bike Model is [%s] \n",b.model)
}

type battery_car struct{
	LifeOfBattery int 
}

type Electric struct{
	car
	battery_car
}
func (ex Electric) move(){
	fmt.Printf("Hey you Buy a new Genaration Car [%s] Best Model [%s] Charge [%d]",ex.model,ex.make,ex.LifeOfBattery)
}
func main(){
	vehical := []MOVES{
		car{model:"X89OPR",make:"BMW X"},
		bike{model:"H2R Y 9O342xC3"},
		Electric{
			car:car{model:"EVCAR9034",make:"Tesla 50T3X"},
			battery_car:battery_car{LifeOfBattery:100},
		},
	}
	for _ , v := range vehical{
		v.move()
	}
}

