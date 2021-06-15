package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {

	fmt.Printf("\t\tUsing unicode package\n")
	ch := 'A'
	fmt.Println(unicode.IsLetter(ch))
	fmt.Println(unicode.IsDigit(ch))
	fmt.Println(unicode.IsSpace(ch))

	fmt.Printf(`This is a raw string and in not interpreted \n\n`)

	fmt.Printf("\n\n\t\tUsing strings\n\n")
	// String Concatination with '+' operator
	s := "Hel" + "lo "
	s += "World!"
	fmt.Printf("The concatinated string is %s with a length of %d \n\n", s, len(s))

	// Checking for the following in a given string
	var str string = "This is a sample string"
	// 1 - Prefix
	fmt.Println(strings.HasPrefix(str, "Thi"))
	// 2 - Suffix
	fmt.Println(strings.HasSuffix(str, "ting"))
	// 3 - substring
	fmt.Println(strings.Contains(str, "a s"))

	// Checking for the first and last index of pattern in a string
	str = "Hi, I'm Marc, Hi"
	fmt.Println(strings.Index(str, "Hi"))
	fmt.Println(strings.LastIndex(str, "Hi"))

	// Count number of non-overlapping pattern in a string
	str1 := "ggggggggg" // Nine continuous g
	fmt.Println(strings.Count(str1, "gg"))

	// Replacing a substring
	// The last argument 'n' denotes the first n occurence to replace
	// if n = - 1 , it replaces all the old pattern with new pattern
	fmt.Println(strings.Replace(str, "Marc", "Samuel", -1))
	fmt.Println(strings.Replace(str, "Hi", "Hello", -1))

	// Repeat a string
	fmt.Println(strings.Repeat("Hi! ", 3))

	// Changing Case of a string
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

	// Triming in strings
	s = "\t  Hello World!   "
	fmt.Println(s)
	fmt.Println(strings.TrimSpace(s))
	fmt.Println(strings.Trim(s, "\t"))

	//Split string using whitespaces
	fmt.Println(strings.Fields(str))

	// Spliting strings on basis of a separator
	fmt.Println(strings.Split(str, ","))

	// Using strconv package
	fmt.Printf("\n\t\t Using strconv package\n\n")

	var orig string = "666"
	var num int
	var newStr string
	num, _ = strconv.Atoi(orig)
	fmt.Printf("String converted to a number represented here: %d\n", num)
	num += 5
	newStr = strconv.Itoa(num)
	fmt.Printf("The new string is %s\n", newStr)

}
