package main
import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Print(i, " even")
		} else {
			fmt.Print(i, " odd")
		}
		switch i {
			case 0: fmt.Println(" Zero")
			case 1: fmt.Println(" One")
			case 2: fmt.Println(" Two")
			case 3: fmt.Println(" Three")
			case 4: fmt.Println(" Four")
			default: fmt.Println("")
		}
	}
	for i:= 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		}
	}
}

