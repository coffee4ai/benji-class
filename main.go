package main

import "math"

//function to print hello world
func main() {
	println("Hello World")
	//print Binyameen Lanka
	println("Binyameen Lanka")
	//print average of 20345 and 1357
	println(average(20345, 1357))
	//print area of a triangle with width 10 and height 5
	println(area(10, 5))
	//print perimeter of a rectangle with width 10 and height 5
	println(perimeter(10, 5))
	//print area of a circle with radius 13
	println(areaCircle(13))
}

//function to get average of two numbers
func average(a, b int) int {
	return (a + b) / 2
}

//function to get area of a triangle
func area(width, height int) int {
	return width * height / 2
}

//function to get perimeter of a rectangle
func perimeter(width, height int) int {
	return 2 * (width + height)
}

//function to get area of a circle
func areaCircle(radius int) float64 {
	return math.Pi * math.Pow(float64(radius), 2)
}
