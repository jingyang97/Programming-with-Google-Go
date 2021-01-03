/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s =½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := GenParameter("acceleration")
	v0 := GenParameter("initial velocity")
	s0 := GenParameter("initial displacement")

	fn := GenDisplaceFn(a, v0, s0)

	var t float64
	for true {
		t = GenParameter("time")
		fmt.Println("\nThe displacement is: ")
		fmt.Println(fn(t))
	}

}

func GenParameter(param string) float64 {
	var userInput string

	for true {

		fmt.Println("\nPlease enter " + param + ": ")
		fmt.Scanln(&userInput)

		f, err := strconv.ParseFloat(userInput, 64)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			return f
		}
	}
	return 0

}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return 0.5*a*t*t + v0*t + s0
	}

	return fn
}