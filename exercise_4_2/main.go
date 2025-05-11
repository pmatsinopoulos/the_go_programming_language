package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	var algo = flag.String("sha", "SHA256", "Hash algorithm: SHA256, SHA384 or SHA512")
	flag.Parse()

	var hasher hash.Hash

	switch *algo {
	case "SHA256":
		hasher = sha256.New()
	case "SHA384":
		hasher = sha512.New384()
	case "SHA512":
		hasher = sha512.New()
	default:
		fmt.Fprintf(os.Stderr, "Unsupported algorithm: %s\n", *algo)
		os.Exit(1)
	}

	if _, err := io.Copy(hasher, os.Stdin); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	fmt.Printf("%x\n", hasher.Sum(nil))
}
