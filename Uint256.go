package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/holiman/uint256"
)

type UintWrap uint256.Int

//go:noescape
func _U256Add(a, b, res unsafe.Pointer)

func (z *UintWrap) AsmAdd(x,y *UintWrap) *UintWrap {
	_U256Add(unsafe.Pointer(x), unsafe.Pointer(y), unsafe.Pointer(z))
	return z;
}

func ufromhex(hex string) uint256.Int {
	a, _ := uint256.FromHex(hex)
	return *a;
}

func main() {
	N := 100000000
	start := time.Now()

	fmt.Println("c2go asm benchmark")
	aa := UintWrap(ufromhex("0x123456789abcdeffedcba9876543210f2f3f4f5f6f7f8f9fff3f4f5f6f7f8f9"))
	bb := UintWrap(ufromhex("0x123456789abcdeffedcba9876543210f2f3f4f5f6f7f8f9fff3f4f5f6f7f8f9"))

	start = time.Now()
	for i := 0; i < N; i++ {
		bb.AsmAdd(&aa, &bb)
	}
	fmt.Println(float64(time.Since(start).Nanoseconds()) / float64(N))

	bbp := uint256.Int(bb)
	fmt.Println(bbp)
	
	fmt.Println("holiman/uint256 asm benchmark")
	a := ufromhex("0x123456789abcdeffedcba9876543210f2f3f4f5f6f7f8f9fff3f4f5f6f7f8f9")
	b := ufromhex("0x123456789abcdeffedcba9876543210f2f3f4f5f6f7f8f9fff3f4f5f6f7f8f9")


	start = time.Now()
	for i := 0; i < N; i++ {
		b.Add(&a, &b)
	}
	fmt.Println(float64(time.Since(start).Nanoseconds()) / float64(N))
	fmt.Println(b)
	
}
