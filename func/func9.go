package main
import "fmt"

func main(x uint) unit {
	if x == 0 {
		return 1
	}
	return x * main(x - 1)
}

// Đệ Quy