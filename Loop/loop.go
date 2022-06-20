package main

import "fmt"

func main() {
	// Vòng Lặp For
	var i int = 1
    for i <= 10 {
        fmt.Println(i)
        i  = i + 1
    }

	// Vòng Lặp For Dạng 2
	for a := 1; a <= 10 ; a++ {
		fmt.Println(a)
	}


	// Else If
	for b := 1; b <= 10; b++ {
		if b % 2 == 0 {
			fmt.Println(b, "Chẵn")
		} else {
			fmt.Println(i, "Lẻ")
		}
	}

	var c int = 2
	if c == 0 {
		fmt.Println("Khong")
	} else if c == 1 {
		fmt.Println("Mot")
	} else if c == 2 {
		fmt.Println("Hai")
	} else if c == 3 {
		fmt.Println("Ba")
	} else if c == 4 {
		fmt.Println("Bon")
	} else if c == 5 {
		fmt.Println("Nam")
	}

	// Switch Case
	var d int = 5
	switch d {
		case 0:  fmt.Println("Khong")
		case 1:  fmt.Println("Mot")
		case 2:  fmt.Println("Hai")
		case 3:  fmt.Println("Ba")
		case 4:  fmt.Println("Bon")
		case 5:  fmt.Println("Nam")
		default: fmt.Println("Khong biet")
		}
 }