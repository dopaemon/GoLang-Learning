package main
import "fmt"

func chuvi(arr []int) int {
	var a, b, c int
	a = arr[0]
	b = arr[1]
	c = (a + b) * 2
	return c
}

func main () {
	var d, e int
	fmt.Print("Nhập vào cạnh a: ")
	fmt.Scanln(&d)
	fmt.Print("Nhập vào cạnh b: ")
	fmt.Scanln(&e)
	xs := []int{d, e}
	fmt.Println("Diện tích hình chữ nhật cạnh a=", d, "Cạnh b=", e, "là", chuvi(xs))
}

// (a + b) * 2