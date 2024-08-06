package main

import "fmt"

func arraysSlicesLoops(){
	return
	// var intArr [3]int32 = [3]int32{1,2,3}
	// intArr := [3]int32{1,2,3}
	// ... infers size by looking at the initial value
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)

	// This is a slice (essentially a dynamic non-fixed value array. It is an array)
	// You initialize a slice by omitting the length of the array
	// when you add an element it will check the arrays CAPACITY and if it won't fit it adds more
	// the LENGTH of the array and the CAPACITY of the array can differ
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)

	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice = append(intSlice, 7)

	fmt.Printf("The length is %v with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// this makes an int32 array with a length of 3 and capacity of 8
	// var intSlice3 []int32 = make([]int32, 3, 8)

	// MAPS
	// key value
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])

	var age, ok = myMap2["Jason"]

	if (ok) {
		fmt.Sprintln(age)
	} else {
		fmt.Println("Jason not found")
	}

	// delete a value from a map. Deletes by reference so it doesn't return anything
	delete(myMap2, "Adam")

	// Name and age of each key value in the myMap2 map
	for name,age:= range myMap2{
		fmt.Printf("Name: %v\nAge: %v\n", name, age)
	}

	// index and value of the intArr
	for i, v := range intArr{
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	// standard for loop syntax
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// alternative for loop syntax
	i := 0
	for{
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}

	
}