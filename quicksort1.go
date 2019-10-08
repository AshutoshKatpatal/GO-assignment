package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var str string
	var n int
	loop := true

	for loop {
		fmt.Println("Select the following operations to be done by quicksort")

		fmt.Println(" 1.Dataset Evenly Distributed \n 2.Dataset Evenly distributed \n 3.Mixed Array of int and char\n 4. Exit")

		var a int

		fmt.Scan(&a)

		switch a {
		case 1:
			fmt.Println("\n\nEnter even Size of array")
			fmt.Scan(&n)
			slice := createSlice(n)
			fmt.Println("\n--- Unsorted Evenly distributed Array --- \n\n", slice)
			quicksort(slice)
			fmt.Println("\n--- Sorted Evenly distributed Array ---\n\n", slice)
		case 2:
                        fmt.Println("\n\nEnter odd no for Size of array")
			fmt.Scan(&n)
			slicea := createSlice(n)
			fmt.Println("\n--- Unsorted Unevenly distributed Array --- \n\n", slicea)
			quicksort(slicea)
			fmt.Println("\n--- Sorted Unevelny distributed Array ---\n\n", slicea)
		case 3:
			fmt.Println("\n--- Enter a String of char and int ---\n\n")
			fmt.Scan(&str)
			r := createRSlice(str)
			quicksort(r)
			fmt.Println("\n--- Sorted String in ASCII --- \n\n", r)
		case 4:
			loop = false
		default:
			fmt.Println("INVALID SELECTION")
		}

	}

}

// Generates a slice of size, size filled with random numbers
func createSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	lo, hi := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[hi] = a[hi], a[pivot]

	for i, _ := range a {
		if a[i] < a[hi] {
			a[lo], a[i] = a[i], a[lo]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]

	quicksort(a[:lo])
	quicksort(a[lo+1:])

	return a
}
func createRSlice(a string) []int {
	runes := []rune(a)
	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	fmt.Println("\n--- Unsorted Array of int and Char in ASCII is --- \n\n", result)
	return result
}
