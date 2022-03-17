/*
Let us assume the following formula for
displacement s as a function of time t, 
acceleration a, initial velocity vo,
and initial displacement so.
s = ½ a t2 + vot + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity,
and initial displacement.Then the program should 
prompt the user to enter a value for time and the
program should compute the displacement after the 
entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and
initial displacement so. GenDisplaceFn() should return
a function which computes displacement as a function of 
time,assuming the given values acceleration, initial 
velocity, and initial displacement. The function returned
by GenDisplaceFn() should take one float64 argument t, 
representing time, and return one float64 argument which
is the displacement travelled after time t.
*/


package main

import "fmt"

func main() {
	var a, u, di, t float64

	fmt.Println("Enter Acceleration: ")
	fmt.Scan(&a)
	fmt.Println("Enter Velocity: ")
	fmt.Scan(&u)
	fmt.Println("Enter Initial Displacement: ")
	fmt.Scan(&di)
	fmt.Println("Enter Time: ")
	fmt.Scan(&t)

	fn := GenDisplayFn(a, u, di)

	fmt.Printf("The calculated displacement is: %v", fn(t))

}

func GenDisplayFn(a, u, di float64) func(float64) float64 {
	return func(t float64) float64 {
		return u*t + 0.5*a*t*t + di
	}
}
