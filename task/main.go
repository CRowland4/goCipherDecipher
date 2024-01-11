package main

import "fmt"

func main() {
	g, p := getLetters()
	fmt.Printf("g=%d and p=%d", g, p)
	return
}

func getLetters() (g, p int) {
	fmt.Scanf("g is %d and p is %d", &g, &p)
	return g, p
}
