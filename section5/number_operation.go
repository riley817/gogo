package main

import (
	"math"
)

func main()  {
	// Example 1 (overflows Error : 범위 초과)
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint16 = math.MaxUint16 + 1
	var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1

	// Example 2 (overflows Error : 범위 초과)
	var n5 uint8 = -1
	var n6 uint16 = -1
	var n7 uint32 = -1
	var n8 uint64 = -1
}
