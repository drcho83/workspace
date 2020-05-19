package main

import (
	"math"
)

func main() {
	//example1 (오버 플로우)
	var n1 uint8 = math.MaxUint8 + 1
	var n2 uint16 = math.MaxUint16 + 1
	var n3 uint32 = math.MaxUint32 + 1
	var n4 uint64 = math.MaxUint64 + 1

	//example2
	var n1 uint8 = -1
	var n5 uint16 = -1
	var n6 uint32 = -1
	var n7 uint64 = -1

}
