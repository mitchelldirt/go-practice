package main

import "fmt"

type gasEngine struct{
	mpg uint16
	gallons uint16
}

type electricEngine struct{
	mpkwh uint16
	kwh uint16
}

func (e gasEngine) milesLeft() uint16{
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint16{
	return e.mpkwh * e.kwh
}

type engine interface{
	milesLeft() uint16
}

func canMakeIt(e engine, miles uint16){
	if miles <= e.milesLeft(){
		fmt.Println("You can make it!")
	} else{
		fmt.Println("Fuel up pardner!")
	}
}

type owner struct{
	name string
}

func structsInterfaces() {
	return
	var myEngine gasEngine = gasEngine{mpg: 20, gallons: 15}
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{10, 10}
	myEngine.mpg = 21
	fmt.Println(myEngine.mpg, myEngine.gallons)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	var newEngine gasEngine = gasEngine{25, 15}
	var newerEngine electricEngine = electricEngine{30, 20}

	// by using interfaces the canMakeIt function now supports both gas and electric types
	canMakeIt(newEngine, 100)
	canMakeIt(newerEngine, 1000)
}