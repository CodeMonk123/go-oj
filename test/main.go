package main

import "fmt"

func f(a *[]int) {
	*a = append(*a, 1)
}

func g(a []int) {
	a[0] = 3
}

func main() {
	a := []int{0}
	fmt.Println(a)
	f(&a)
	fmt.Println(a)
	g(a)
	fmt.Println(a)

	b := "Hello世界"
	for _, c := range b {
		fmt.Printf("%c\n", c)
	}

	for _, c := range []byte(b) {
		fmt.Printf("%c\n", c)
	}
}
