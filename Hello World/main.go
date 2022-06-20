// packages main sẻ chứa điểm thực thi của chương trình
package main

// Import Lib Giúp in ra màn hình
import "fmt"
// Chúng ta import package CGO thông qua câu lệnh import "C". Chương trình trên chưa thực hiện bất kì thao tác nào với CGO, chỉ mới thông báo sẵn sàng cho việc lập trình với CGO. Mặc dù chưa sử dụng gì đến CGO nhưng lệnh go build vẫn sẽ gọi trình biên dịch gcc trong suốt quá trình biên dịch do đây được là một chương trình CGO hoàn chỉnh.
// import "C"

// Scope đọc ở dưới
var e string = "Xin Chào GoLang -  Scope"

// là hàm đầu tiên được chạy trong chương trình
func main() {
	// Hello World
	fmt.Println("Xin chào GoLang\nDòng 2")

	// Toán
	fmt.Println("1 + 1 = ", 1 + 1)

	// Kiểu Giá Trị
	/*
	 * uint8	0 – 255
	 * uint16	0 – 65535
	 * uint32	0 – 4294967295
	 * uint64	0 – 18446744073709551615
	 * int8	    -128 – 127
	 * int16	-32768 – 32767
	 * int32	-2147483648 – 2147483647
	 * int64	-9223372036854775808 – 9223372036854775807
	 */
	fmt.Println(len("Xin Chào GoLang"))
	fmt.Println("Xin Chào GoLang"[1])
	fmt.Println("Xin Chào " + " GoLang")

	// Boolean
	/*
	 * BIỂU THỨC	GIÁ TRỊ
	 * true && true		true
	 * true && false	false
	 * false && true	false
	 * false && false	false
	 * Phép OR:
	 *
	 * BIỂU THỨC	GIÁ TRỊ
	 * true || true	true
	 * true || false	true
	 * false || true	true
	 * false || false	false
	 * Phép NOT:

	 * BIỂU THỨC	GIÁ TRỊ
	 * !true	false
	 * !false	true
	 */

	fmt.Println(true && false)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Var Khai báo
	var a string = "Xin chào GoLang"
	fmt.Println(a)

	// var Khai báo 2
	var b string
	b = "Xin Chào GoLang"
	fmt.Println(b)

	// var khai báo 3
	var c string
	c = "Xin Chào"
	fmt.Println(c)
	c = "GoLang"
	fmt.Println(c)

 	// Var Khai báo 4
	// Ở đây chúng ta không dùng từ khóa var, không khai báo kiểu dữ liệu, thay vào đó chúng ta ghi tên biến rồi dùng toán tử := theo sau là giá trị để khai báo nhanh một biến và gán giá trị ngay tại chỗ. Trình biên dịch Go sẽ tự động nhận diển kiểu dữ liệu dựa vào giá trị mà bạn gán cho biến. Chẳng hạn như ở đây Go thấy giá trị là “Hello World”, tức là một string nên sẽ tự động cho biến x kiểu dữ liệu string.
	// Tóm lại nó giống var d string
	d := "Xin Chào GoLang"
	fmt.Println(d)

	// Tên Biến
	nameserver := "Doraemon"
	fmt.Println("Tên Server:", nameserver)

	// Scope Như trên
	fmt.Println(e)

	// Hằng Số - Constants
	// Hằng số đơn giản chỉ là các biến không thể thay đổi dược giá trị. Cách khai báo và gán giá trị cho hằng số cũng giống như với biến, chỉ khác một chỗ là thay từ khóa var bằng từ khóa const
	const f string = "Xin chào GoLang"
	fmt.Println(f)

	// Khai báo nhiều biến
	// Chúng ta dùng từ khóa var (hoặc const), theo sau là cặp dấu ngoặc tròn (), rồi tới danh sách các biến và giá trị của chúng.
	var (
		g = 1
		h = 2
		i = 3
	)
}

func main1() {
	fmt.Println(e)
}