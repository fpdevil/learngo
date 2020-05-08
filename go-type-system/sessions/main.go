package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
	"unsafe"
)

// type definitions usually declared at the package level
//
// EXERCISE: Move the declaration into main()'s scope
//
type (
	gram  float64 // float64 is the underlying type of gram
	ounce float64 //float64 is the underlying type of ounce
)

func main() {
	fmt.Println("--- 01-bits ---")

	// 1 bit can encode 2 different state: 0 or 1
	fmt.Printf("%d -> %[1]b, %d -> %[2]b\n", 0, 1)

	// 2 bits can encode 4 different states
	// 0 0
	// 0 1
	// 1 0
	// 1 1
	fmt.Println("2 bits can encode 4 different states")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d => %02[1]b\n", i)
	}

	// %08b prints 8 bits with leading zeros if any
	fmt.Println()
	fmt.Printf("%d => %08[1]b\n", 1)
	fmt.Printf("%d => %08[1]b\n", 2)
	fmt.Printf("%d => %08[1]b\n", 4)
	fmt.Printf("%d => %08[1]b\n", 8)
	fmt.Printf("%d => %08[1]b\n", 16)
	fmt.Printf("%d => %08[1]b\n", 32)
	fmt.Printf("%d => %08[1]b\n", 64)
	fmt.Printf("%d => %08[1]b\n", 128)

	// calculation of bits
	fmt.Println("bit calculation...")
	i, _ := strconv.ParseInt("00000100", 2, 64)
	fmt.Println("00000100 = ", i)
	i, _ = strconv.ParseInt("00010110", 2, 64)
	fmt.Println("00010110 = ", i)

	fmt.Println("--- 02-bytes ---")
	// a byte is an integer number with 8 bits (1 byte)
	var b byte

	// all bits are either empty or 0
	b = 0 // min number a byte can have
	fmt.Printf("%08b = %d\n", b, b)

	// all bits are either full or 1
	b = 255 // max number a byte can have
	fmt.Printf("%08b = %d\n", b, b)

	fmt.Println()
	fmt.Println("---- 03-predeclared-types ----")

	// getting limits of numeric types
	// int32 and rune types are same; rune is an alias to int32
	fmt.Println("limits of integer types...")
	fmt.Println("int8	:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16	:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32	:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64	:", math.MinInt64, math.MaxInt64)

	// unsigned numbers take only positive numbers
	// byte and uint8 are same types and byte is an alias o unit8
	fmt.Println()
	fmt.Println("limits of unsigned number types")
	fmt.Println("uint8		:", 0, math.MaxUint8)
	fmt.Println("uint16		:", 0, math.MaxUint16)
	fmt.Println("uint32		:", 0, math.MaxUint32)
	fmt.Println("uint64		:", 0, uint64(math.MaxUint64))
	// MaxUint64 is a very large number which cannot be printed
	// directly; it can only be used in constant expressions

	// floating point types
	fmt.Println()
	fmt.Println("limits of floating point number types")
	fmt.Println("float32	:", math.SmallestNonzeroFloat32)
	fmt.Println("float64	:", math.SmallestNonzeroFloat64)

	// memory limits
	fmt.Println()
	fmt.Println("size/memory of types...")
	fmt.Println("int	:", unsafe.Sizeof(int(1)), "bytes")
	fmt.Println("int8	:", unsafe.Sizeof(int8(1)), "bytes")
	fmt.Println("int16	:", unsafe.Sizeof(int16(1)), "bytes")
	fmt.Println("int32	:", unsafe.Sizeof(int32(1)), "bytes")
	fmt.Println("int64	:", unsafe.Sizeof(int64(1)), "bytes")

	fmt.Println()
	fmt.Println("size/memory of unsigned types...")
	fmt.Println("uint	:", unsafe.Sizeof(uint(1)), "bytes")
	fmt.Println("uint8	:", unsafe.Sizeof(uint8(1)), "bytes")
	fmt.Println("uint16	:", unsafe.Sizeof(uint16(1)), "bytes")
	fmt.Println("uint32	:", unsafe.Sizeof(uint32(1)), "bytes")
	fmt.Println("uint64	:", unsafe.Sizeof(uint64(1)), "bytes")

	fmt.Println()
	fmt.Println("size/memory of floating types...")
	fmt.Println("float32	:", unsafe.Sizeof(float32(1)), "bytes")
	fmt.Println("float64	:", unsafe.Sizeof(float64(1)), "bytes")

	fmt.Println()
	fmt.Println("--- testing destruction ---")
	big := uint16(65535)
	// uint8 destroys its value
	// to its own max value which is 255
	//
	// 65535 - 255 is lost.
	small := uint8(big)
	fmt.Printf("%016b %[1]d\n", big)
	fmt.Printf("%016b %[1]d\n", small)

	fmt.Println()
	fmt.Println("---- 01-duration-example ----")
	// a defined type is a type that can be created from another
	// existing type by someone
	// A Defined type is also called a Named Type
	h, _ := time.ParseDuration("4h30m")
	fmt.Println("4h32m", h.Hours(), "hours")
	fmt.Println("4h30m", h.Minutes(), "minutes")

	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	// underlying types
	fmt.Printf("Tyoe of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))

	// creating own type definitions
	fmt.Println("02-type-definition-create-your-own-type")
	var g gram = 1000
	var o ounce

	// ounce and gram are different types
	// o = g * 0.03524

	o = ounce(g) * 0.03524
	fmt.Printf("%g grams is %.2f ounce\n", g, o)

	fmt.Println("")
	fmt.Println("---- 06-aliased-types ----")

	// aliased types are same types but with different names
	// for readability and refactoring
	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	uint8Val = byteVal // this works

	var (
		runeVal  rune
		int32Val int32
	)

	runeVal = int32Val // this works

	runeVal = rune(int32Val)
	runeVal = rune(runeVal)

	// for silencing the compiler
	_, _ = uint8Val, intVal

}
