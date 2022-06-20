package main
import "fmt"

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	fmt.Println(add(1, 2, 3))
}

 /*
  *Chúng ta thêm 3 dấu chấm "..." 
  * vào trước tên kiểu dữ liệu của 
  * tham số cuối cùng, Go sẽ hiểu 
  * rằng tham số này có thể có 0 hoặc 
  * nhiều giá trị được truyền vào, khi 
  * gọi hàm chúng ta có thể truyền vào 0 
  * hoặc nhiều giá trị, ngăn cách nhau bởi 
  * dấu phẩy, đặc tính này cho phép hàm 
  * nhận tham số một cách linh hoạt hơn.
  */