package main

func squaredRange(num int) int {
	var result int

	for i := 1; i <= num; i++ {
		result += (i * i)
	}

	return result
}