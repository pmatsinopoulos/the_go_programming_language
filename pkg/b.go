package pkg

import "fmt"

func init() {
	fmt.Println("init() in package 'pkg' file 'b.go'")
}

func FromB() {
	fmt.Println("FromB() in package 'pkg' file 'b.go'")
}
