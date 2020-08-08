package main

import (
	"fmt"
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
	fmt.Println(what(true))
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


	case bool:
		return fmt.Sprintf("Boolean: %v", v)	

	default:
		return "Unknown type"
	}

}
