package main

import (
	"fmt"
	"math/rand"
)

func main() {
	g, p := getGAndP()
	fmt.Println("OK")

	b := rand.Intn(p-3) + 2
	B := calculateB(g, p, b)

	A := getA()
	S := calculateSharedSecret(A, b, p)

	fmt.Print("B is ", B, "\n")
	fmt.Print("A is ", A, "\n")
	fmt.Print("S is ", S, "\n")
	return
}

func getGAndP() (g, p int) {
	fmt.Scanf("g is %d and p is %d\n", &g, &p)
	return g, p
}

func calculateB(g, p, b int) (B int) {
	B = 1
	for i := 1; i <= b; i++ {
		B = (B * g) % p
	}

	return B
}

func getA() (A int) {
	fmt.Scanf("A is %d\n", &A)
	return A
}

func calculateSharedSecret(A, b, p int) (S int) {
	S = 1
	for i := 1; i <= b; i++ {
		S = (S * A) % p
	}

	return S
}
