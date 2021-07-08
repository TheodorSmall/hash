package main

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

import (
	_ "crypto/md5"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "golang.org/x/crypto/blake2b"
	_ "golang.org/x/crypto/blake2s"
	_ "golang.org/x/crypto/md4"
	_ "golang.org/x/crypto/ripemd160"
	_ "golang.org/x/crypto/sha3"
)

func calcHash(algorithm crypto.Hash, password string) string {
	if !algorithm.Available() {
		fmt.Println("No implementation found")
		os.Exit(1)
	}

	hash := algorithm.New()
	io.WriteString(hash, password)
	decodedhash := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(decodedhash)
}

func calcAlgorithm(algorithm string) crypto.Hash {
	switch strings.ToLower(algorithm) {
	case "md4":
		return crypto.MD4
	case "md5":
		return crypto.MD5
	case "sha1":
		return crypto.SHA1
	case "sha224":
		return crypto.SHA224
	case "sha256":
		return crypto.SHA256
	case "sha384":
		return crypto.SHA384
	case "sha512":
		return crypto.SHA512
	case "md5sha1":
		return crypto.MD5SHA1
	case "ripemd160":
		return crypto.RIPEMD160
	case "sha3-224":
		return crypto.SHA3_224
	case "sha3-256":
		return crypto.SHA3_256
	case "sha3-384":
		return crypto.SHA3_384
	case "sha3-512":
		return crypto.SHA3_512
	case "sha512-224":
		return crypto.SHA512_224
	case "sha512-256":
		return crypto.SHA512_256
	case "blake2s-256":
		return crypto.BLAKE2s_256
	case "blake2b-256":
		return crypto.BLAKE2b_256
	case "blake2b-384":
		return crypto.BLAKE2b_384
	case "blake2b-512":
		return crypto.BLAKE2b_512

	default:
		fmt.Println("Unknown algorithm")
		os.Exit(1)
		return 0
	}
}

func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Usage: hash <algorithm> <password>")
		os.Exit(1)
	}

	algorithm := calcAlgorithm(os.Args[1])
	passhash := calcHash(algorithm, os.Args[2])

	fmt.Printf("%s\n", passhash)
}
