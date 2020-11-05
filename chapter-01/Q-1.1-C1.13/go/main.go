package main

import "fmt"

func printTable(n int) {
	printRecursive(n, 10)
}

func printRecursive(n, base int) {
	if base < 1 {
		return
	}
	printRecursive(n, base-1)
	fmt.Printf("%d * %d = %d\n", n, base, (n * base))
}

func main() {
	printTable(8)
}
