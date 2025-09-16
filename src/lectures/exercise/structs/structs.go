//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width int
}


func calculateArea(r Rectangle) int {
	return r.length * r.width
}

func calculatePerimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

func increaseRectangleTimes(r *Rectangle, timesLength int, timesWidth int) (int, int){
	r.length *= timesLength
	r.width *= timesWidth
	return r.length, r.width
}


func main() {
	r := Rectangle {length: 10,width: 20}

	fmt.Println(calculateArea(r))
	fmt.Println(calculatePerimeter(r))
	x,y := increaseRectangleTimes(&r, 2, 2)
	fmt.Println(x, " , ", y)

	fmt.Println(calculateArea(r))
	fmt.Println(calculatePerimeter(r))
}



