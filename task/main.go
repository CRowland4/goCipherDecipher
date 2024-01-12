package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	g, p := getGAndP()
	fmt.Println("OK")

	A := getA()

	b := rand.Intn(p-3) + 2
	B := calculateB(g, p, b)
	fmt.Printf("B is %d\n", B)

	S := calculateEncryptionKey(A, b, p)
	fmt.Println(encryptString(S, "Will you marry me?"))
	response := getProposalResponse()
	reaction := getResponseReaction(S, response)

	if reaction != "" {
		fmt.Println(encryptString(S, reaction))
	}

	return
}

// getGAndP reads and returns the variables g and p from StdIn, with the input format of "g is _ and p is _".
func getGAndP() (g, p int) {
	fmt.Scanf("g is %d and p is %d\n", &g, &p)
	return g, p
}

// calculateB uses the encryption variables g, p, and b to calculate and return the encryption variable B.
func calculateB(g, p, b int) (B int) {
	B = 1
	for i := 1; i <= b; i++ {
		B = (B * g) % p
	}

	return B
}

// getA reads and returns the variable A from StdIn, with the input format of "A is _".
func getA() (A int) {
	fmt.Scanf("A is %d\n", &A)
	return A
}

// calculateEncryptionKey uses the encryption variables A, b, and p to calculate and return the shared key S.
func calculateEncryptionKey(A, b, p int) (S int) {
	S = 1
	for i := 1; i <= b; i++ {
		S = (S * A) % p
	}

	return S
}

// encryptString uses the shared encryption key S to encrypt and return string_ using a Caesar Cipher
func encryptString(S int, string_ string) (encryptedString string) {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < len(string_); i++ {
		char := string(string_[i])
		alphabetIndex := strings.Index(alphabet, strings.ToLower(char))
		encryptedIndex := (alphabetIndex + (S % 26)) % 26

		if alphabetIndex == -1 { // Char is not a letter, but punctuation or symbol
			encryptedString += char
		} else if char == strings.ToUpper(char) {
			encryptedString += strings.ToUpper(string(alphabet[encryptedIndex]))
		} else {
			encryptedString += string(alphabet[encryptedIndex])
		}
	}

	return encryptedString
}

// getProposalResponse reads and returns one full line from StdIn.
func getProposalResponse() (proposalResponse string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	proposalResponse = scanner.Text()
	return proposalResponse
}

// getResponseReaction uses the shared encryption key S to determine and return the correct reaction to the response.
func getResponseReaction(S int, response string) (reaction string) {
	if encryptString(S, response) == encryptString(S, "Yeah, okay!") {
		reaction = "Great!"
	} else if encryptString(S, response) == encryptString(S, "Let's be friends.") {
		reaction = "What a pity!"
	} else {
		return ""
	}

	return reaction
}
