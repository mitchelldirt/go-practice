package main

import "fmt"

func pointers(){
	return
	var p *int32 = new(int32)
	var i int32

	fmt.Printf("The value p points to is: %v \n", *p)
	fmt.Printf("The value of i is: %v \n", i)

	p = &i
	*p = 1

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	var k int32 = 2
	i = k

	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)

	// IMPORTANT SLICES ARE BY REFERENCE NOT VALUE
	// The below example will print the same exact array

	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 5
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// If you add a new element to one of the arrays it only affects that array
	sliceCopy = append(sliceCopy, 9)
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// Pointers work great to save memory. Especially when working with large arrays
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("The memory location of the thing1 array is: %p\n", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("The result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("The memory location of the thing2 array is: %p\n", thing2)
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}