package main
import "fmt"
import "os"
import "os/exec"

func main() {
	var b int = 0

	var c int
	var e int = 0
	fmt.Printf("Nhập vào số phần tử: ")
	fmt.Scanln(&c)

	var d = make([]int, c)
	for i := 0; i < c; i++ {
		fmt.Print("Nhập vào phần tử ", i, ": ")
		fmt.Scanln(&e)
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
        cmd.Run()
		d[i] = e
	}
	fmt.Println("Mảng của bạn là:", d)

	for i := 0; i < c; i++ {
		b = b + d[i]

	}
	fmt.Println("Kết quả mảng trên là: ", b)
}