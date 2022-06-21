package main
import (
	"fmt"
	"math"
)

func delta(a, b, c int) int {
	var delta int
	delta = b * b - 4 * a * c
	return delta
}

func ptb2(a, b, c, d int) int {
	var x, x1, x2 float64
	var a1, b1, d1 float64

	a1 = float64(a)
	b1 = float64(b)
	d1 = float64(d)

	if a == 0 {
		fmt.Println("Bạn cần nhập a > 0")
	} else if d < 0 {
		fmt.Println("Phương trình vô nghiệm")
	} else if d == 0 {
		x = -b1 / (2 * a1);
		fmt.Println("Phương Trình Có 2 Nghiệm Phân Biệt X1 = X2 =", x)
	} else if d > 0 {
		x1 = (-b1 + math.Sqrt(d1) ) / (2 * a1)
        x2 = (-b1 - math.Sqrt(d1) ) / (2 * a1)
        fmt.Println("Nghiệm X1 =", x1)
		fmt.Println("Nghiệm X2 =", x2)
	}
	return 0;
}

func main() {
	var aa, bb, cc int
	fmt.Print("Nhập vào a: ")
	fmt.Scanln(&aa)
	fmt.Print("Nhập vào b: ")
	fmt.Scanln(&bb)
	fmt.Print("Nhập vào c: ")
	fmt.Scanln(&cc)
	arr := []int{aa, bb, cc}
	d := delta(arr[0], arr[1], arr[2])
	ptb2(arr[0], arr[1], arr[2], d)
}