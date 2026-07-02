package main

import (
	"fmt"
	"os"
	"text/tabwriter"
) // Extra formation, table writer

func main() {
	fmt.Println("---------------------------------")
	fmt.Println("Unsigned integer types")
	fmt.Println("---------------------------------")

	uit := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(uit, "Type\tMin\tMax")

	// Rows
	var u1min uint8 = 0
	var u1max uint8 = 255
	fmt.Fprintf(uit, "uint8\t%v\t%v\n", u1min, u1max)

	var u2min uint16 = 0
	var u2max uint16 = 65535
	fmt.Fprintf(uit, "uint16\t%v\t%v\n", u2min, u2max)

	var u3min uint32 = 0
	var u3max uint32 = 4294967295
	fmt.Fprintf(uit, "uint32\t%v\t%v\n", u3min, u3max)

	var u4min uint64 = 0
	var u4max uint64 = 18446744073709551615
	fmt.Fprintf(uit, "uint64\t%v\t%v\n", u4min, u4max)

	var byteValMin byte = 0   // byte is an alias for uint8
	var byteValMax byte = 255 // byte is an alias for uint8
	fmt.Fprintf(uit, "byte\t%v\t%v\n", byteValMin, byteValMax)

	uit.Flush()

	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("Signed integer types")
	fmt.Println("---------------------------------")

	it := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(it, "Type\tMin\tMax")

	// Rows
	var i1min int8 = -128
	var i1max int8 = 127
	fmt.Fprintf(it, "int8\t%v\t%v\n", i1min, i1max)

	var i2min int16 = -32768
	var i2max int16 = 32767
	fmt.Fprintf(it, "int16\t%v\t%v\n", i2min, i2max)

	var i3min int32 = -2147483648
	var i3max int32 = 2147483647
	fmt.Fprintf(it, "int32\t%v\t%v\n", i3min, i3max)

	var i4min int64 = -9223372036854775808
	var i4max int64 = 9223372036854775807
	fmt.Fprintf(it, "int64\t%v\t%v\n", i4min, i4max)

	var runeValMin rune = -2147483648 // rune is an alias for int32
	var runeValMax rune = 2147483647  // rune is an alias for int32
	fmt.Fprintf(it, "rune\t%v\t%v\n", runeValMin, runeValMax)

	it.Flush()

	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("Floating-point types")
	fmt.Println("---------------------------------")

	ft := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(ft, "Type\tMin\tMax")

	// Rows
	var f1min float32 = -3.402823466e+38
	var f1max float32 = 3.402823466e+38
	fmt.Fprintf(ft, "float32\t%v\t%v\n", f1min, f1max)

	var f2min float64 = -1.7976931348623157e+308
	var f2max float64 = 1.7976931348623157e+308
	fmt.Fprintf(ft, "float64\t%v\t%v\n", f2min, f2max)

	ft.Flush()

	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("Complex types")
	fmt.Println("---------------------------------")

	ct := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(ct, "Type\tMin\tMax")

	// Rows
	// float32 + imaginary float32
	var c1min complex64 = complex(-3.402823466e+38, -3.402823466e+38)
	var c1max complex64 = complex(3.402823466e+38, 3.402823466e+38)
	fmt.Fprintf(ct, "complex64\t%v\t%v\n", c1min, c1max)

	// float64 + imaginary float64
	var c2min complex128 = complex(-1.7976931348623157e+308, -1.7976931348623157e+308)
	var c2max complex128 = complex(1.7976931348623157e+308, 1.7976931348623157e+308)
	fmt.Fprintf(ct, "complex128\t%v\t%v\n", c2min, c2max)

	ct.Flush()

	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("String type")
	fmt.Println("---------------------------------")

	st := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(st, "Type\tValue")

	// Rows
	var str1 string = "Hello, World!"
	fmt.Fprintf(st, "string\t%v\n", str1)

	// String with special characters
	var str2 string = "Hello, \nWorld!\t😊"
	fmt.Fprintf(st, "string with special characters\t%v\n", str2)

	// Multiline string using backticks
	var str3 string = `Hello,
	World!`
	fmt.Fprintf(st, "multiline string\t%v\n", str3)

	fmt.Fprintf(st, "string length\t%v\n", len(str1))
	fmt.Fprintf(st, "unicode char at position 7\t%v\n", string(str1[7]))

	st.Flush()

	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("Boolean type")
	fmt.Println("---------------------------------")

	bt := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	// Header
	fmt.Fprintln(bt, "Type\tValue")

	// Rows
	var b1 bool = true
	fmt.Fprintf(bt, "bool\t%v\n", b1)

	var b2 bool = false
	fmt.Fprintf(bt, "bool\t%v\n", b2)

	bt.Flush()

	// Another way to declare a variable
	var name, surname string = "John", "Doe"
	fmt.Println("")
	fmt.Println("---------------------------------")
	fmt.Println("Multiple variable declaration")
	fmt.Println("---------------------------------")
	fmt.Printf("Name: %s, Surname: %s\n", name, surname)

	// Another one
	var (
		age  int    = 30
		city string = "New York"
	)
	fmt.Printf("Age: %d, City: %s\n", age, city)

	// Short variable declaration
	country := "USA"
	fmt.Printf("Country: %s\n", country)

}
