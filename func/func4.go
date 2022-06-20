package main
import "fmt"

func f1() (r int) {
	r = 99999
	return
}

func main() {
	a := f1()
	fmt.Print("Tổng của " , a, " Và ", a, " Là ", a + a)
}