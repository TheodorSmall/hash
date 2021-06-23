package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func hash(password string) string {
	hash := sha512.New()
	io.WriteString(hash, password)
	decodedhash := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(decodedhash)
}

func main() {
	if len(os.Args[1:]) < 1 {
		fmt.Printf("No command line argument to hash")
		os.Exit(1)
	}

	password := os.Args[1]
	fmt.Printf("%s\n", hash(password))
}
