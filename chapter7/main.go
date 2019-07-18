package main
import "fmt"

func average(xs ...float64) (float64, int) {
	//panic("Not implemented")
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs)), len(xs)
}

// This is basically shorthand for having a "setup" function along
// with the function that does the real work. This Generator should
// only be called once, since that will run the "setup" as well.
// The output of the generator is what can be called multiple times.
// So the value should persist as long as the variable holding the
// function returned by this Generator is in scope.
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// defer, panic and recover
func learnToPanic() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()		// these round brackets simply call the function that
			// was just defined!
	panic("PANIC")
}

func sum(numbers []int) int {
	total := 0
	for _,v := range numbers {
		total += v
	}
	return total
}

func halfEven(a int) (int, bool) {
	convertToBool := func(b int) bool {
		if (b > 0) {
			return true
		} else {
			return false
		}
	}
	return a/2, convertToBool((a/2) % 2)
}

func kanye(args ...int) int {
	k := 0
	for _, v := range args {
		if k < v {
			k = v
		}
	}
	return k
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (r uint) {
		r = i
		i += 2
		return
	}
}

func fib(x uint) uint {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println(fib(8))
	fmt.Println("")

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println("")

	ints := []int{5,6,7,8}
	fmt.Println(kanye(ints...))
	fmt.Println("")

	fmt.Println(halfEven(5))
	fmt.Println(halfEven(6))
	fmt.Println("")

	xs := []float64{98,93,77,82,83}
	addToSlice := func(x float64) {
		xs = append(xs, x)
	}

	addToSlice(77)
	addToSlice(93)
	fmt.Println(average(xs...))
	for _,v := range xs {
		fmt.Println(v)
	}

	fmt.Println("")
	nextEven := makeEvenGenerator()
	fmt.Println(makeEvenGenerator()) // returns a function pointer
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	fmt.Println("")
	fmt.Println(factorial(5))

	fmt.Println("")
	learnToPanic()

	fmt.Println("")
	fmt.Println(sum([]int{1,2,3}))
}

