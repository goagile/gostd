package main

import (
	"fmt"
	"math"
)

func main() {

	
	fmt.Println()
	fmt.Println(what(uint8(255)))

	fmt.Println()
	fmt.Println(what(uint16(65535)))

	fmt.Println()
	fmt.Println(what(uint32(4294967295)))

	fmt.Println()
	fmt.Println(what(uint64(18446744073709551615)))

	fmt.Println()
	fmt.Println(what(uint(18446744073709551615)))

	fmt.Println()
	fmt.Println(what(uintptr(18446744073709551615)))


	fmt.Println()
	fmt.Println("byte 255 is ", what(byte(255)))


	fmt.Println()
	fmt.Println(what(int8(-128)))
	fmt.Println(what(int8( 127)))

	fmt.Println()
	fmt.Println(what(int16(-32768)))
	fmt.Println(what(int16( 32767)))
	
	fmt.Println()
	fmt.Println(what(int32(-2147483648)))
	fmt.Println(what(int32( 2147483647)))

	fmt.Println()
	fmt.Println(what(int64(-9223372036854775808)))
	fmt.Println(what(int64( 9223372036854775807)))

	fmt.Println()
	fmt.Println(what(int(-9223372036854775808)))
	fmt.Println(what(int( 9223372036854775807)))


	fmt.Println()
	fmt.Println("math.MaxFloat32 is ", math.MaxFloat32)
	fmt.Println(what(float32(math.MaxFloat32)))

	fmt.Println()
	fmt.Println("math.MaxFloat64 is ", math.MaxFloat64)
	fmt.Println(what(float64(math.MaxFloat64)))


	fmt.Println()
	fmt.Println(what(complex(math.MaxFloat64, math.MaxFloat64)))


	fmt.Println()
	fmt.Println(what(true))


	fmt.Println()
	fmt.Println("rune 'a' is ", what(rune('a')))

	fmt.Println()
	fmt.Println(what("Hello"))

}

func what(v interface{}) string {

	switch v.(type) {
	

	case uint8:
		return fmt.Sprintf("Unsigned Integer 8 bit: %v", v)

	case uint16:
		return fmt.Sprintf("Unsigned Integer 16 bit: %v", v)

	case uint32:
		return fmt.Sprintf("Unsigned Integer 32 bit: %v", v)

	case uint64:
		return fmt.Sprintf("Unsigned Integer 64 bit: %v", v)

	case uint:
		return fmt.Sprintf("Unsigned Integer: %v", v)

	case uintptr:
		return fmt.Sprintf("Unsigned Integer Pointer: %v", v)


	case int8:
		return fmt.Sprintf("Integer 8 bit: %v", v)

	case int16:
		return fmt.Sprintf("Integer 16 bit: %v", v)		

	case int32:
		return fmt.Sprintf("Integer 32 bit: %v", v)

	case int64:
		return fmt.Sprintf("Integer 64 bit: %v", v)

	case int:
		return fmt.Sprintf("Integer: %v", v)


	case float32:
		return fmt.Sprintf("Float32: %v", v)

	case float64:
		return fmt.Sprintf("Float64: %v", v)


	case complex64:
		return fmt.Sprintf("complex64: %v", v)

	case complex128:
		return fmt.Sprintf("complex128: %v", v)


	case bool:
		return fmt.Sprintf("Boolean: %v", v)	


	case string:
		return fmt.Sprintf("String: %v", v)	


	default:
		return "Unknown type"
	}

}
