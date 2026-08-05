package main

import (
	crypto_rand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func integersAndFloats() {
	// rand.Seed(time.Now().UnixNano()) // Starting with Go version 1.20, the math/rand package automatically seeds the global random number generator, making it random by default.
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(20))

	fmt.Println(rand.Float64())
	fmt.Println(5 + rand.Float64()*10) // will generate floating numbers in the interval of [5, 15)
}

func secureRandomNumbers() {
	randomNumber, _ := crypto_rand.Int(crypto_rand.Reader, big.NewInt(100))
	fmt.Println(randomNumber)

	randomPrimeNumber, _ := crypto_rand.Prime(crypto_rand.Reader, 8)
	fmt.Println(randomPrimeNumber)
}

func randomCharacter() {
	charset := "abcdefghijklmnopqrstuvwxyz"
	randomPosition := rand.Intn(len(charset))
	c := charset[randomPosition]

	fmt.Println(string(c))
}

func randomString() {
	charset := "abcdefghijklmnopqrstuvwxyz"
	length := 10

	random := make([]byte, length)
	for i := 0; i < length; i++ {
		randomPosition := rand.Intn(len(charset))
		c := charset[randomPosition]
		random[i] = c
		// random[i] = charset[rand.Intn(len(charset))]
	}

	str := string(random)
	fmt.Println(str)
}

func randomStringsUsingASCII() {
	// rand.Seed(time.Now().UnixNano())
	length := 10

	randomUpperCaseString := make([]byte, length)
	for i := 0; i < length; i++ {
		// upper-case letters lie within [65, 90
		randomUpperCaseString[i] = byte(65 + rand.Intn(26))
	}

	str1 := string(randomUpperCaseString)
	fmt.Println(str1)

	randomLowerCaseString := make([]byte, length)
	for i := 0; i < length; i++ {
		// lower-case letters are in the interval of [97, 122]
		randomLowerCaseString[i] = byte(97 + rand.Intn(26))
	}

	str2 := string(randomLowerCaseString)
	fmt.Println(str2)
}

func main() {
	fmt.Println("Generating random numbers and strings")

	// integersAndFloats()
	// secureRandomNumbers()
	// randomCharacter()
	// randomString()
	randomStringsUsingASCII()
}
