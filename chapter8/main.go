package main
import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swapInts(x *int, y *int) {
	z := *x
	*x = *y
	*y = z
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)

	y := 1.5
	square(&y)
	fmt.Println(y)

	x = 10;
	z := 15;
	swapInts(&x, &z)
	fmt.Println(x, z)
}
