package main

import (
	"fmt"
	"strings"

	"crypto/sha512"
	"github.com/anaskhan96/go-password-encoder"
)

func main() {
	//// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	// Using custom options
	options := &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}
	salt, encodedPwd := password.Encode("generic password", options)
	dbPassword := fmt.Sprintf("%s$%s", salt, encodedPwd)
	fmt.Println(len(dbPassword))
	fmt.Println(dbPassword)
	passwords := strings.Split(dbPassword, "$")
	fmt.Println(passwords)
	check := password.Verify("generic password", passwords[0], passwords[1], options)
	fmt.Println(check) // true
}
