package pkg

import "fmt"

func init() {
	fmt.Println("init() for package 'pkg' file 'a.go'")
}

func FromA() {
	fmt.Println("Function FromA in package 'pkg' file 'a.go'")
}
