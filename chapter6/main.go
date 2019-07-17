package main
import "fmt"

func array() {
	fmt.Println("Arrays")

	var a[5]int
	x := [5]float64{98, 93, 77, 82, 83}

	a[4] = 100
	fmt.Println(a)

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println(total / float64(len(x)))

	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func slices() {
	fmt.Println("\nSlices")

	//x := make([]float64, 5, 10)
	//arr := [5]float64{1,2,3,4,5}
	//y := arr[1:]

	s1 := []int{1,2,3}
	s2 := append(s1, 4, 5)
	fmt.Println(s1, s2)

	z := make([]int, 3, 9)
	fmt.Println(len(z))

	x := []int{
		48, 96, 86, 68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	min := x[0]
	for _, value := range x {
		if value < min {
			min = value
		}
	}
	fmt.Println("min =", min)
}

func maps() {
	fmt.Println("\nMaps")

	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])

	elements := map[string]map[string]string {
		"H": map[string]string{
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name": "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name": "Lithim",
			"state": "solid",
		},
		"Be": map[string]string{
			"name": "Berylium",
			"state": "solid",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	} else {
		fmt.Println("element Li not found")
	}
}

func main() {
	array()
	slices()
	maps()
}
