package main

import (
	"fmt"
	"strings"
	"time"
)

func stringRunesBytes(){
	var myString = []rune("Résumé")
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}

	// Since we've case myString as an array of runes we get the correct length
	// Had we not done that we would get the length of unicode segments... I think
	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	// You can make a rune type by using single quotes
	var myRune = 'a'
	fmt.Printf("\nMy rune = %v\n", myRune)

	// The below is very inefficient because we create a new string every time we add to catStr
	start := time.Now()
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("%v\n", catStr)
	fmt.Printf("Time to build string without string builder: %v\n", time.Since(start))

	// This is much more efficient
	start2 := time.Now()
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	var catStrBetter = strBuilder.String()
	
	fmt.Printf("%v\n", catStrBetter)
	fmt.Printf("With string builder: %v\n", time.Since(start2))
}