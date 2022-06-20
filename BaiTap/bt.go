package main
import "fmt"

func main() {
	var a int
	var b int
	var c int

	fmt.Printf("Nhập vào a: ")
	fmt.Scanln(&a)
	fmt.Printf("Nhập vào b: ")
	fmt.Scanln(&b)

	c = a + b

	fmt.Println(a, "+", b, "=", c)
}