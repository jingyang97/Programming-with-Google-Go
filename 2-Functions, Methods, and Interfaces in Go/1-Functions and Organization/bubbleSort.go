/*
Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	integers := []int{}
	integers = getIntegers(integers)
	fmt.Println("\nYour integers: ")
	fmt.Println(integers)

	bubbleSort(integers)
	fmt.Println("\nIntegers after sorting: ")
	fmt.Println(integers)
}

func getIntegers(integers []int) []int {
	var userInput string

	for len(integers) < 10 {
		fmt.Println("Please enter an integer.")
		fmt.Scanln(&userInput)

		i, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			integers = append(integers, i)
		}
	}

	return integers
}

func bubbleSort(integers []int) {
	for i := len(integers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if integers[j-1] > integers[j] {
				swap(integers, j, j-1)
			}
		}
	}
}

func swap(integers []int, i int, j int) {
	tmp := integers[j]
	integers[j] = integers[i]
	integers[i] = tmp
}
