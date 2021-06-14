package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey := hexutil.Encode(privateKeyBytes)
	fmt.Printf("private key:\t%s\n", privKey[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("ewallet: error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey := hexutil.Encode(publicKeyBytes)
	fmt.Printf("public key:\t%s\n", pubKey[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("address:\t%s\n", address)
}
