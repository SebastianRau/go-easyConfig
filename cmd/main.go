package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sebastianrau/go-easyConfig/pkg/encryption"
)

func main() {
	var (
		createKeyFiles    = flag.String("c", "", "Create a new AES Key Pair with given filename")
		stringToEncrypt   = flag.String("s", "", "String to encrypt")
		encryptionKeyPath = flag.String("k", "", "Encryption Key File")
	)
	flag.Parse()

	fmt.Println("Easy Config KeyGen & String encryption")

	//Check if Pathes are given, else use home directory/.kiosk/
	if *createKeyFiles != "" {
		log.Printf("Creating key pair: %s", *createKeyFiles)
		err := encryption.CreateKeyFile(*createKeyFiles)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		os.Exit(0)
	}

	if *stringToEncrypt == "" {
		fmt.Println("no string to encrypt given")
		os.Exit(-2)
	}
	if *encryptionKeyPath == "" {
		fmt.Println("no keyfiles to encrypt given")
		os.Exit(-2)
	}

	key, err := os.ReadFile(*encryptionKeyPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(-3)
	}

	enc, err := encryption.EncryptString(*stringToEncrypt, key)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-2)
	}
	fmt.Println("Encrytion of string is:")
	fmt.Printf("${%s}\n", enc)
}
