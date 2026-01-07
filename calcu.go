package main

import ("fmt")
func main() {
	var a, b float64
	var op string
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan operator (+, -, *, /): ")
	fmt.Scan(&op)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scan(&b)

	var hasil float64

	switch op {
	case "+":
		hasil = a + b
	case "-":
		hasil = a - b
	case "*":
		hasil = a * b
	case "/":
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol tidak diperbolehkan.")
			return
		}
		hasil = a / b
	default:
		fmt.Println("Operator tidak valid.")
		return
	}

	fmt.Printf("Hasil: %.2f\n", hasil)
}
