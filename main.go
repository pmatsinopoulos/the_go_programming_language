package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"the_go_programming_language/count_diff_bits"
	"the_go_programming_language/pkg"
	"the_go_programming_language/popcount"
)

func init() {
	fmt.Println("init() in package 'pkg' file 'main.go'")
}

func main() {
	fmt.Println("Hello, World!")

	var a [3]int
	a[0] = 5
	a[1] = 10
	a[2] = 15

	fmt.Println(a[0])
	fmt.Println("==========")
	fmt.Println(a[len(a)-1])
	fmt.Println("--------------")

	for i, v := range a {
		fmt.Println(i, v)
	}

	for _, v := range a {
		fmt.Println(v)
	}

	var q [3]int = [3]int{8, 15, 20}

	for _, v := range q {
		fmt.Println(v)
	}

	fmt.Println("Array literal initialization")
	var r [3]int = [3]int{1, 3, 5} // this is an array literal
	for _, v := range r {
		fmt.Println(v)
	}

	// I don't have to set value for all the elements of the array
	var s [3]int = [3]int{1, 5}
	for _, v := range s {
		fmt.Println(v)
	}

	fmt.Println("Array length can be determined by the number of elements in the initializer")
	var b = [...]int{8, 9, 15, 32}
	for _, v := range b {
		fmt.Println(v)
	}
	fmt.Printf("Printing type of a variable b = %T\n", b)

	fmt.Printf("Working with Currency symbol")

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "E", GBP: "P", RMB: "R"}
	fmt.Printf("%v\n", symbol)
	for _, v := range symbol {
		fmt.Println(v)
	}
	fmt.Println(symbol[USD])

	fmt.Println("-----------------")

	c := [...]int{99: -1} // this means that element at position 99 takes the value -1. All other, 0 to 98, have value 0.
	fmt.Println("len(c) = ", len(c))
	for _, v := range c {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	fmt.Println("comparable arrays")

	d := [2]int{1, 2}
	e := [...]int{1, 2}
	f := [2]int{1, 3}
	fmt.Println(d == e, d == f, e == f)

	fmt.Println("Print digest of x and X")

	x := "x"
	X := "X"

	fmt.Println(x, X)

	c1 := sha256.Sum256([]byte(x))
	c2 := sha256.Sum256([]byte(X))
	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("Type of c1 %T\n", c1)

	g := [3]int{8, 15, -3}
	doSomething(g)

	for _, v := range g {
		fmt.Println(v)
	}

	doSomething2(&g)

	for _, v := range g {
		fmt.Println(v)
	}

	zeroElements(&g)

	for _, v := range g {
		fmt.Println(v)
	}

	h := [32]byte(sha256.Sum256([]byte("x")))
	i := [32]byte(sha256.Sum256([]byte("X")))

	numberOfDifferentBits := count_diff_bits.CountDifferentBits(&h, &i)
	fmt.Printf("x = %b\nX = %b\n\tNumber of different bits = %v\n", h, i, numberOfDifferentBits)

	pkg.FromA()
	pkg.FromB()

	var aUint64 uint64 = rand.Uint64()
	fmt.Printf("%x\n\t%v\n\t%b\n", aUint64, aUint64, aUint64)

	var numberOfBitsSet int = popcount.PopCount(aUint64)
	fmt.Printf("Number of bits: %v\n", numberOfBitsSet)

	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December",
	}
	fmt.Printf("length of months: %d, First month: %s\n", len(months), months[1])

	slice1 := months[2:6]
	for _, v := range slice1 {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Length of slice1 = %d, capacity of slice1 = %d\n", len(slice1), cap(slice1))
	slice2 := slice1[:6]
	fmt.Printf("slice2[5] = %s\n", slice2[5])
}

func doSomething(a [3]int) {
	for i := range len(a) {
		a[i] += 1
	}

	for _, v := range a {
		fmt.Println(v)
	}
}

func doSomething2(a *[3]int) {
	for i := range a { // look how do I loop over all elements of the array that the pointer points to. Amazing!
		a[i] += 1
	}
}

func zeroElements(ptr *[3]int) {
	*ptr = [3]int{}
}
