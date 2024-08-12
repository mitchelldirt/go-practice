package main

import "fmt"

type gasEngine2 struct {
	gallons float32
	mpg     float32
}

type electricEngine2 struct {
	kwh   float32
	mpkwh float32
}

type car[T gasEngine2 | electricEngine2] struct {
	carMake  string
	carModel string
	engine   T
}

func generics() {
	var intSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(sumSlice[float32](float32Slice))

	var float64Slice = []float64{1.0, 2.0, 3.0, 4.0, 5.0}
	fmt.Println(sumSlice[float64](float64Slice))

	var gasCar = car[gasEngine2]{
		carMake:  "Honda",
		carModel: "Accord",
		engine: gasEngine2{
			gallons: 1.0,
			mpg:     2.0,
		},
	}

	fmt.Println(gasCar)

	var electricCar = car[electricEngine2]{
		carMake:  "Honda",
		carModel: "Accord E",
		engine: electricEngine2{
			kwh:   1.0,
			mpkwh: 2.0,
		},
	}

	fmt.Println(electricCar)

	return
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
