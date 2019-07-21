package main
import "fmt"
import "math"

//Shapes
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// interfaces define a set of methods
// if a type includes all these methods, it
// "implements" this interface
type Shape interface {
	area() float64
	perimeter() float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r*c.r
}

// function with a "receiver" is called a method
func(c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

func (c *Circle) perimeter() float64 {
	return math.Pi * 2*c.r
}

func (r *Rectangle) length() float64 {
	return distance(r.x1, r.y1, r.x1, r.y2)
}

func (r* Rectangle) width() float64 {
	return distance(r.x1, r.y1, r.x2, r.y1)
}

func (r *Rectangle) area() float64 {
	return r.length() * r.width()
}

func (r *Rectangle) perimeter() float64 {
	return 2*(r.length() + r.width())
}

func totalArea(shapes ...Shape) (float64, float64) {
	var area, perimeter float64
	for _, s := range shapes {
		area += s.area()
		fmt.Println("perimeter: ", s.perimeter())
		perimeter += s.perimeter()
	}
	return area, perimeter
}

// interface as a field in another type
type MultiShape struct {
	shapes []Shape
}

// turn MultiShape into a Shape (MultiShape itself can implement Shape)
func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// How to use MultiShape???
func printAreas(mul *MultiShape) {
	for _, s := range mul.shapes {
		fmt.Println("area: ", s.area())
	}
}

//People
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person		// embedded type "Android is a Person" - can call
	Model string	// any Person methods directly
}

func main() {
	// Shapes
	//========
	//var c Circle
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	r := Rectangle{0, 0, 10, 10}

	fmt.Println("circle coordinates: ", c.x, c.y, c.r)
	fmt.Println("circle area: ", circleArea(&c))
	fmt.Println("circle area: ", c.area())
	fmt.Println("rectangle area: ", r.area())

	ar, pr := totalArea(&c, &r)
	fmt.Println("total area: ", ar, "\ntotal perimeter: ", pr)

	//shapeSlice := [2]Shape{&c, &r}
	//m := MultiShape{shapeSlice}

	//m := new(MultiShape)
	//m.shapes.append(Circle{0,0,4})


	fmt.Println()
	// People
	//========
	a := new(Android)
	a.Name = "Marcus"
	a.Model = "12cd"
	a.Person.Talk()
	a.Talk()	// embedded types

	b := Android{Person{"Jane"}, "33q"}
	b.Talk()
}

