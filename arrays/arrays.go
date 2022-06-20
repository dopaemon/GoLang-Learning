package main

import "fmt"

func main() {
	// Dòng x[4] = 100 có nghĩa là gán giá trị 100 cho phần tử thứ 5 trong mảng, các phần tử không được gán giá trị sẽ có giá trị mặc định là 0. Lý do phần tử đó là phần tử thứ 5 chứ không phải thứ 4 là vì các phần tử trong array được đánh số thứ tự từ 0.
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// Đoạn code trên tính giá trị trung bình của các phần tử trong mảng. Khi chạy chương trình sẽ in kết quả ra 20
	var a [5]float64
	a[0] = 10
	a[1] = 15
	a[2] = 20
	a[3] = 25
	a[4] = 30

	var total float64 = 0
	for i:=0; i<5; i++ {
		total = total + a[i]
	}
	fmt.Println(total / 5)

	// Một kiểu khác dùng len
	// đếm các phần tử trong mảng bằng len
	var b [5]float64
	b[0] = 10
	b[1] = 15
	b[2] = 20
	b[3] = 25
	b[4] = 30
	var trungbinh float64 = 0
	for i := 0; i < len(a); i++ {
		trungbinh = trungbinh + b[i]
	}
	fmt.Println(trungbinh / float64(len(b)))

	// Một cách khác Ngắn hơn
	var c [5]float64
	c[0] = 10
	c[1] = 15
	c[2] = 20
	c[3] = 25
	c[4] = 30
	var tb float64 = 0
	// _ là biến giả, vì golang khai báo biến không dùng sẻ gây lỗi
	for _, value := range c {
		tb += value
	}
	fmt.Println(tb / float64(len(c)))

	// Khai báo Array Nhanh
	d := [5]float64{10, 15, 20, 25, 30}
	fmt.Println(d)

	// Kiểu dữ liệu Slice
	var e []float64
	fmt.Println(e)

	// Hàm Make
	// chúng ta có thể tạo một slice bằng hàm make()
	f := make([]float64, 5)
	fmt.Println(f)

	// Chúng ta có thể tạo slice từ một array bằng cách dùng biểu thức [low:high]
	arr := [5]float64{10, 15, 20, 25, 30}
	g := arr[0:5]
	h := arr[:] // Lấy toàn bộ array
	i := arr[0:] // Lấy toàn bộ phần tử từ 0
	fmt.Println(g[3])
	fmt.Println(h)
	fmt.Println(i)

	// Ngoài ra trong còn Go có 2 hàm hỗ trợ làm việc với slice là append() và copy()
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1,"\n", slice2)

	// Ví dụ đối với hàm copy()
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := make([]int, 4)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)

	// Hàm Map
	j := make(map[string]int)
	j["key"] = 10
	fmt.Println(j["key"])
	// delete(j, "key") để xoá phần tử key trong mảng

	// Thực chất khi truy xuất 
	// một phần tử của map chúng ta nhận 
	// được 2 giá trị trả về chứ không chỉ 
	// có giá trị của khóa, giá trị thứ 2 
	// là một giá trị kiểu boolean cho 
	// biết việc truy xuất có thành công 
	// hay không. Nếu chúng ta truy xuất một 
	// khóa có tồn tại trong map thì giá trị 
	// boolean trả về true, ngược lại trả về 
	// false
	k := make(map[string]int)
	k["key2"] = 10

	value2, ok := k["key2"]
	fmt.Println(value2, ok)

	value3, ok2 := k["key3"]
	fmt.Println(value3, ok2)

	// Vừa khai báo giá trị vừa gán vào map
	elements := map[string]string{
		"H": "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be" : "Beryllium",
		"B" : "Boron",
		"C" : "Carbon",
		"N" : "Nitrogen",
		"O" : "Oxygen",
		"F" : "Fluorine",
		"Ne" : "Neon",
	}
	fmt.Println(elements)
}