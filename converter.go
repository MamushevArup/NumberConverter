package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	for  {
		fmt.Printf("\nWhat you need a) decimal to (octal, binary, hex) b) binary to (decimal, octal, hex) c) octal to (decimal, binary, hex) d) hex to (decimal, binary, octal)\n")
		var choice string
		fmt.Scan(&choice)
		pointers(choice)	
	}
}
func pointers(choose string){
  if choose == "a"{
    decimality()
  }else if choose == "b"{
	binaryti()
  }else if choose == "c"{
	octality()
  }else if choose == "d"{
	hexality()
  }
}
func decimality() {
	var choices string
	fmt.Println("decimal to a) octal b) binary c) hex")
	fmt.Scan(&choices)
	fmt.Println("Input the decimal value")
	var convert int
	fmt.Scan(&convert)
	if choices == "a" {
		fmt.Printf("Convert decimal %v to octal %v", convert, decimalToOctal(convert))
	}else if choices == "b"{
		fmt.Printf("Convert decimal %v to binary %v", convert, decimalToBinary(convert))
	}else if choices == "c"{
		fmt.Printf("Convert decimal %v to hex %v", convert, decimalToHex(convert))
	}

}
func binaryti(){
	var choices string
	fmt.Println("Binary to a) decimal b) hex c) octal")
	fmt.Scan(&choices)
	fmt.Println("Input the binary value")
	var convert string
	fmt.Scan(&convert)
	if choices == "a"{
		fmt.Printf("Convert binary %v to decimal %v", convert, binarytodec(convert))
	}else if choices == "b" {
		fmt.Printf("Convert binary %v to hex %v", convert, binaryToHex(convert))
	}else if choices == "c" {
		fmt.Printf("Convert binary %v to octal %v", convert, binaryToOctal(convert))
	}
}
func octality(){
	var choices string
	fmt.Println("Octal to a) decimal b) binary c) hex")
	fmt.Scan(&choices)
	fmt.Println("Input the octal value")
	var convert int
	fmt.Scan(&convert)
	if choices == "a"{
		fmt.Printf("Convert octal %v to decimal %v", convert, octalToDecimal(convert))
	}else if choices == "b"{
		fmt.Printf("Convert octal %v to binary %v", convert, octalToBinary(convert))
	}else if choices == "c"{
		fmt.Printf("Convert octal %v to hex %v", convert, octalToHex(convert))
	}
}
func hexality(){
	var choices string
	fmt.Println("Hex to a) decimal b) binary c) hex")
	fmt.Scan(&choices)
	fmt.Println("Input hex value")
	var convert string
	fmt.Scan(&convert)
	if choices == "a" {
		fmt.Printf("Convert hex %v to decimal %v", convert, hexToDecimal(convert))
	}else if choices == "b"{
		fmt.Printf("Convert hex %v to binary %v", convert, hexToBinary(convert))
	}else if choices == "c"{
		fmt.Printf("Convert hex %v to octal %v", convert, hexToOctal(convert))
	}
}
func decimalToBinary(x int) string {
	res := ""
	for x >= 1 {
		if x%2 == 0 {
			res += "0"
		} else {
			res += "1"
		}
		x = int(math.Floor(float64(x) / 2))
	}
	return reverseString(res)
}
func decimalToOctal(x int) int {
	res := 0
	i := 1
	for x != 0 {
		res += (x % 8) * i
		x /= 8
		i *= 10
	}
	return res
}
func decimalToHex(value int) string {
	res := ""
	arr := [6]string{"A", "B", "C", "D", "E", "F"}
	for value >= 1 {
		if value%16 > 9 {
			res += arr[(value%16)%10]
		}else {
			res += strconv.Itoa(value % 16)
		}
		value = int(value / 16)
	}
	return reverseString(res)
}
func reverseString(value string) string {
	res := ""
	for i := len(value) - 1; i >= 0; i-- {
		res += string(value[i])
	}
	return res
}

func binarytodec(x string) int {
	for i := 0; i < len(x); i++ {
		curr, _ := strconv.Atoi(string(x[i]))
		if curr >= 2 || curr < 0{
			fmt.Printf("You input wrong value\n")
			return 0
		}
	}
	res := 0
	base := 1
	for i := len(x) - 1; i >= 0; i-- {
		if string(x[i]) == "1" {
			res += base
		}
		base = base * 2
	}
	return res
}
func binaryToHex(value string) string{
	for i := 0; i < len(value); i++ {
		curr, _ := strconv.Atoi(string(value[i]))
		if curr >= 2 || curr < 0{
			fmt.Printf("You input wrong value\n")
			return "0"
		}
	}
	res := binarytodec(value)
	res2 := decimalToHex(res)
	return res2
}
func binaryToOctal(x string) int{
	for i := 0; i < len(x); i++ {
		curr, _ := strconv.Atoi(string(x[i]))
		if curr >= 2 || curr < 0{
			fmt.Printf("You input wrong value\n")
			return 0
		}
	}
	res := binarytodec(x)
	res2 := decimalToOctal(res)
	return res2
}
func octalToDecimal(x int) int{
	vr := strconv.Itoa(x)
	for i := 0; i < len(vr); i++ {
		if rune(vr[i]) >= 57 || rune(vr[i]) <= 47{
			fmt.Println("You input wrong value")
			return 0
		}
	}
	base := 1
    res := 0
	for x >= 1{
		digit := x % 10
		if (digit > 8){
			return 0
		}
		res += digit * base
		x = x / 10
		base = base * 8
	}
	return res
}
func octalToBinary(x int) string{
	vr := strconv.Itoa(x)
	for i := 0; i < len(vr); i++ {
		if rune(vr[i]) >= 57 || rune(vr[i]) <= 47{
			fmt.Println("You input wrong value")
			return "0"
		}
	}
	res := octalToDecimal(x)
	res2 := decimalToBinary(res)
	return res2
}
func octalToHex(x int) string {
	vr := strconv.Itoa(x)
	for i := 0; i < len(vr); i++ {
		if rune(vr[i]) >= 57 || rune(vr[i]) <= 47{
			fmt.Println("You input wrong value")
			return "0"
		}
	}
	res := octalToDecimal(x)
	res2 := decimalToHex(res)
	return res2
}

func hexToDecimal(x string)int{
	mape := map[string]string{
		"A" : "10",
		"B" : "11", 
		"C" : "12" ,
		"D" : "13" ,
		"E" : "14" ,
		"F" : "15" ,
	}
	arr := []string{}
	for i := 0; i < len(x); i++ {
		a := string(x[i])
		if rune(x[i]) >= 65 && rune(x[i]) <= 70 {
			arr = append(arr, mape[a])
		}else {
			arr = append(arr, a)
		}
	}
	res := 0
	base := 1
	for i := len(arr)-1; i >= 0; i-- {
		conv, _ := strconv.Atoi(arr[i])
		res += base * conv
		base *= 16
	}
	return res
}
func hexToBinary(x string) string{
	res := hexToDecimal(x)
	res2 := decimalToBinary(res)
	return res2
}
func hexToOctal(x string) int{
	res := hexToDecimal(x)
	res2 := decimalToOctal(res)
	return res2
}