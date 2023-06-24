package main

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"strings"
)

func main() {
	// Using the default options
	//salt, encodedPwd := password.Encode("123456", nil)
	//check := password.Verify("123456", salt, encodedPwd, nil)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	//fmt.Println(check) // true

	// Using custom options
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("123456aa", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	passwordInfo := strings.Split(newPassword, "$")
	check := password.Verify("123456aa", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(salt)
	//fmt.Println(encodedPwd)
	fmt.Println(passwordInfo[0]) // ""
	fmt.Println(passwordInfo[1]) // "pbkdf2-sha512"
	fmt.Println(check)           // true
	fmt.Println(newPassword)
}
