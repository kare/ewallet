package main

import (
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
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

func addressToChecksumCase(address string) string {
	a := common.HexToAddress(address)
	return a.Hex()
}

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	out := flag.CommandLine.Output()
	newCmd.Usage = func() {
		fmt.Fprintf(out, "Usage: ewallet new [-h]\n")
		fmt.Fprintf(out, "Generate new private key\n")
	}
	addressCmd := flag.NewFlagSet("address", flag.ExitOnError)
	addressCmd.Usage = func() {
		fmt.Fprintf(out, "Usage: ewallet address [-h] private_key\n")
		fmt.Fprintf(out, "Convert given private key to address\n")
	}
	publicCmd := flag.NewFlagSet("public", flag.ExitOnError)
	publicCmd.Usage = func() {
		fmt.Fprintf(out, "Usage: ewallet public [-h] private_key\n")
		fmt.Fprintf(out, "Convert given private key to public key\n")
	}
	checksumCmd := flag.NewFlagSet("checksum", flag.ExitOnError)
	checksumCmd.Usage = func() {
		fmt.Fprintf(out, "Usage: ewallet checksum [-h] private_key\n")
		fmt.Fprintf(out, "Convert given address to checksum case\n")
	}
	help := flag.Bool("h", false, "help message")
	flag.Parse()
	flag.Usage = func() {
		cmdName := os.Args[0]
		fmt.Fprintf(out, "Usage: %s command\n", cmdName)
		fmt.Fprintf(out, "Flags:\n")
		flag.PrintDefaults()
		usageMessage := `Commands:
	new		Generate new private key
	address		Convert given private key to address
	public		Convert given private key to public key
	checksum	Convert given address to checksum case
`
		fmt.Printf("%s", usageMessage)
	}
	if len(os.Args) <= 1 || *help {
		flag.Usage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "new":
		if err := newCmd.Parse(os.Args[2:]); err == nil {
			privateKey, err := crypto.GenerateKey()
			if err != nil {
				log.Fatalf("error while generating private key: %v", err)
			}
			privateKeyBytes := crypto.FromECDSA(privateKey)
			privKey := hexutil.Encode(privateKeyBytes)
			key := privKey[2:]
			fmt.Printf("%s\n", key)
		} else {
			log.Fatalf("new private key flag parse error: %v", err)
		}
	case "address":
		if len(os.Args) < 3 {
			fmt.Fprintf(flag.CommandLine.Output(), "private key is a required argument\n")
			addressCmd.Usage()
			os.Exit(1)
		}
		if err := addressCmd.Parse(os.Args[2:]); err == nil {
			address := privateKeyToAddress(addressCmd.Args()[0])
			fmt.Printf("%s\n", address)
		} else {
			log.Fatalf("private key to address flag parse error: %v", err)
		}
	case "public":
		if len(os.Args) < 3 {
			fmt.Fprintf(flag.CommandLine.Output(), "private key is a required argument\n")
			publicCmd.Usage()
			os.Exit(1)
		}
		if err := publicCmd.Parse(os.Args[2:]); err == nil {
			public := privateKeyToPublic(publicCmd.Args()[0])
			fmt.Printf("%s\n", public)
		} else {
			log.Fatalf("private key to public key flag parse error: %v", err)
		}
	case "checksum":
		if len(os.Args) < 3 {
			fmt.Fprintf(flag.CommandLine.Output(), "private key is a required argument\n")
			checksumCmd.Usage()
			os.Exit(1)
		}
		if err := checksumCmd.Parse(os.Args[2:]); err == nil {
			address := addressToChecksumCase(checksumCmd.Args()[0])
			fmt.Printf("%s\n", address)
		} else {
			log.Fatalf("private key to public key flag parse error: %v", err)
		}
	}
}
