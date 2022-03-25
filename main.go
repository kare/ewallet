package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func privateKeyToPublicECDSA(key string) *ecdsa.PublicKey {
	key = strings.TrimPrefix(key, "0x")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("error while converting hex private key to ECDSA: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	return publicKeyECDSA
}

func privateKeyToAddress(key string) string {
	publicKeyECDSA := privateKeyToPublicECDSA(key)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	hexAddress := address.Hex()
	return hexAddress
}

func privateKeyToPublic(key string) string {
	publicKeyECDSA := privateKeyToPublicECDSA(key)
	publicKey := publicKeyToString(publicKeyECDSA)
	return publicKey
}

func publicKeyToString(publicKeyECDSA *ecdsa.PublicKey) string {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey := hexutil.Encode(publicKeyBytes)
	return pubKey[4:]
}

func main() {
	addressCmd := flag.NewFlagSet("address", flag.ExitOnError)
	publicCmd := flag.NewFlagSet("public", flag.ExitOnError)
	flag.Parse()
	flag.Usage = func() {
		out := flag.CommandLine.Output()
		fmt.Fprintf(out, "Usage of %s: command\n", os.Args[0])
		flag.PrintDefaults()
		usageMessage := `Commands:
	new		Generate new private key
	address		Convert given private key to address
	public		Convert given private key to public key
`
		fmt.Printf("%s", usageMessage)
	}
	if len(os.Args) <= 1 {
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatalf("error while generating private key: %v", err)
		}
		privateKeyBytes := crypto.FromECDSA(privateKey)
		privKey := hexutil.Encode(privateKeyBytes)
		key := privKey[2:]
		fmt.Printf("%s\n", key)
	case "address":
		if err := addressCmd.Parse(os.Args[2:]); err == nil {
			address := privateKeyToAddress(addressCmd.Args()[0])
			fmt.Printf("%s\n", address)
		} else {
			log.Fatalf("private key to address flag parse error: %v", err)
		}
	case "public":
		if err := publicCmd.Parse(os.Args[2:]); err == nil {
			public := privateKeyToPublic(publicCmd.Args()[0])
			fmt.Printf("%s\n", public)
		} else {
			log.Fatalf("private key to public key flag parse error: %v", err)
		}
	}
}
