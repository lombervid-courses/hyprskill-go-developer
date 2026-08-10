package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hashValueOfString() {
	fmt.Println("\nComputing the hash value of a string")

	md5Hash := md5.New()
	sha256Hash := sha256.New()
	sha512Hash := sha512.New()

	md5Hash.Write([]byte("JetBrains Academy"))
	sha256Hash.Write([]byte("JetBrains Academy"))
	sha512Hash.Write([]byte("JetBrains Academy"))

	fmt.Printf("%x\n", md5Hash.Sum(nil))
	fmt.Printf("%x\n", sha256Hash.Sum(nil))
	fmt.Printf("%x\n", sha512Hash.Sum(nil))

	// fmt.Printf("%x\n", md5.Sum([]byte("JetBrains Academy")))
	// fmt.Printf("%x\n", sha256.Sum256([]byte("JetBrains Academy")))
	// fmt.Printf("%x\n", sha512.Sum512([]byte("JetBrains Academy")))

	// Output:
	// dc5740934090c9ed7aa0b3ec8ac755f3
	// 83ac28f753df3cd80fee3f8ce1770da805856afa2b48c2917aefe5123723c4c9
	// 97e5ee749844c330b4e99779bf2d6487cd22497fcff0c49cc2d736fcf95374d1...
}

func bcryptForPasswordHashing() {
	fmt.Println("\nUsing bcrypt for password hashing")

	password := "TrustNo1"
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b)) // $2a$10$pvp/Bsheb5WvHHHOX4zZPO4LbO7depXdUu4ASc4OI.HZ9yGEpD/mi
}

func bcryptHashAndPlainTextPassword() {
	fmt.Println("\nComparing a bcrypt hash and a plaintext password")

	// TrustNo1
	hash := "$2a$10$pvp/Bsheb5WvHHHOX4zZPO4LbO7depXdUu4ASc4OI.HZ9yGEpD/mi"

	var enteredPassword string
	fmt.Scanln(&enteredPassword) // Ask the user to enter a plaintext password

	// Compares the 'bcrypt' hashed password with its possible plaintext equivalent:
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(enteredPassword)); err != nil {
		log.Fatal(err) // Exit the program if the hashes of the two passwords do not match
	}
	fmt.Println("Passwords matched!")

	// Input: TrustNo1
	// Output: Passwords match!

	// Input: TrustNoOne
	// Output: 2022/03/29 22:11:45 crypto/bcrypt: hashedPassword is not the hash of the given password
}

func hashValueOfFile() {
	fmt.Println("\nComputing the hash value of a file")

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// sha256Hash := md5.New()
	sha256Hash := sha256.New()
	// Copy the data from 'hello.txt' to the 'sha256Hash' interface until reaching EOF:
	if _, err := io.Copy(sha256Hash, file); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%x\n", sha256Hash.Sum(nil)) // eab80bcb0f01e951dbe4edbbe22735383c32081e95ed03b085f8bfd1e4858445
}

func main() {
	fmt.Println("Hashing strings and files")
	// go get golang.org/x/crypto/bcrypt

	hashValueOfString()
	bcryptForPasswordHashing()
	bcryptHashAndPlainTextPassword()
	hashValueOfFile()
}
