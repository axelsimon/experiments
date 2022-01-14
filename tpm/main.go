// Heavily inspired by the example on / thanks to
// https://ericchiang.github.io/post/tpm-keys/

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-tpm/tpm2"
)

func main() {
	f, err := os.OpenFile("/dev/tpmrm0", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf("Opening TPM: %v\nDo you have sufficient permissions?", err)
	}
	defer f.Close()

	var reqb uint
	flag.UintVar(&reqb, "n", 16, "Number of random bytes to return from TPM")
	flag.Parse()

	reqb16 := uint16(reqb)
	out, err := tpm2.GetRandom(f, reqb16)
	if err != nil {
		log.Fatalf("Getting random bytes from TPM: %v", err)
	}

	fmt.Printf("%x\n", out)
}
