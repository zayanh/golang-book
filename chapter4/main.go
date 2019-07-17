package main
import "fmt"

const a string = "global"

func main() {
	var x string
	x = "first"
	fmt.Println(x)
	x += "second"
	fmt.Println(x)

	var y = "hello"
	z := "world"
	fmt.Println(y == z)

	fmt.Println(a)
	f()

	var(
	  b = 10
	  c = 15
	)
	fmt.Print("Enter temp in farenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * (5.0/9.0)

	fmt.Println("b = ", b, "\nc = ", c, "\ntemp in celcius = ", output)
}

func f() {
	fmt.Println(a)
}
