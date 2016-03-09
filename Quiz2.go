//Package MyPackage is CDC-Quiz    Developed by Zohaib Ahmed     Section-C   i12-0182
package MyPackage

//MyPublicVariable is a public variable
var MyPublicVariable = "THIS IS A SIMPLE CALCULATOR"

//MyPrivateVariable is a private variable
var myPrivateVariable = "THIS IS A PRIVATE VARIABLE"

//Add is a function to add two integers
func Add(a, b int) int {
	return a + b
}

//Subtract is a function to add two integers
func Subtract(a, b int) int {
	return a - b
}

//Multiply is a function to add two integers
func Multiply(a, b int) int {
	return a * b
}

//Divide is a function to add two integers
func Divide(a, b int) int {
	if b != 0 {
		return a / b
	}
	return -1
}
