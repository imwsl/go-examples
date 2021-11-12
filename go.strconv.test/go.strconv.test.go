// Strconv package 
// data type(int, string etc) convert utils
// 

package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	// convert string to int
	num1, err1 := strconv.Atoi("199")
	if err1 != nil {
		fmt.Println("convert string to int err: %v\n", err1)
		return
	} else {
		fmt.Printf("type: %T value: %d\n", num1, num1)
	}

	// convert int to string. strconv.Itoa. notice: this function has no error return 
	numstring := strconv.Itoa(999)
	fmt.Printf("convert int to string. type: %T value: %s\n", numstring, numstring)

	// Parse functions
	// func ParseBool(s string) (value bool, err error)
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// bitSize: 0, 8, 16, 32, 64 => int, int8, int16, int32, int64
	// func ParseUnit(s string, base int, bitSize int) (n uint64, err error)
	// func ParseFloat(s string, base int, bitSize int) (f float64, err error)
	fmt.Println("............ Parse functions .......................")
	bStrFlag, bStrError := strconv.ParseBool("f")
	if bStrError != nil {
		fmt.Println("strconv.ParseBool err => %v\n", bStrError)
	} else {
		fmt.Printf("strconv.ParseBool(\"f\") => %v\n", bStrFlag)
	}
	
	// strconv.ParseInt always return in64 whatever bitSize is setted. 
	// you must use int32 or int16 to convert the return value.
	iNum, iError := strconv.ParseInt("78", 10, 32)
	if iError != nil {
		fmt.Println("ParseInt error: %v\n", iError)
	} else {
		fmt.Printf("strconv.ParseInt(\"78\") %v Type: %T\n", iNum, iNum)
	}
	
	// as the ParseInt, ParseUint always return uint64...
	uNum, uError := strconv.ParseUint("79", 10, 8)
	if uError != nil {
		fmt.Println("strconv.ParseUint error => %v\n", uError)
	} else {
		fmt.Printf("strconv.ParseUint(\"79\", 10, 8) ==> %v  Type: %T\n", uNum, uNum)
	}
	
	// ParseFloat always return float64
	fNum, fError := strconv.ParseFloat("199.99", 32)
	if fError != nil {
		fmt.Println("strconv.ParseFloat error ==> %v\n", fError)
	} else {
		fmt.Printf("strconv.ParseFloat(\"199.99\", 32) ==> %v  Type: %T\n", fNum, fNum)
	}

	// Format functions
	// FormatBool(b bool) string
	// FormatInt(i int64, base int) string
	// FormatUint(i uint64, base int) string
	// FormatFloat(f float64, fmt byte,prec int, bitSize int) string

	// isPrint(r rune) bool
	// CanBackquote(s string) bool
}

