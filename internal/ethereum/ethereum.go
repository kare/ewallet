package ethereum

import (
	"crypto/ecdsa"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyToPublicKey converts given private key to ECDSA public key.
func PrivateKeyToPublicKey(key string) *ecdsa.PublicKey {
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

// PrivateKeyToAddress converts given private key to an address in string format.
func PrivateKeyToAddress(key string) string {
	publicKeyECDSA := PrivateKeyToPublicKey(key)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	hexAddress := address.Hex()
	return hexAddress
}

// PrivateKeyToPublic converts given private key to public key in string format.
func PrivateKeyToPublic(key string) string {
	publicKeyECDSA := PrivateKeyToPublicKey(key)
	publicKey := PublicKeyToString(publicKeyECDSA)
	return publicKey
}

// PublicKeyToString converts ECDSA public key to hex format string.
func PublicKeyToString(publicKeyECDSA *ecdsa.PublicKey) string {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey := hexutil.Encode(publicKeyBytes)
	return pubKey[4:]
}

// AddressToChecksumCase converts an Ethereum address to checksum case.
func AddressToChecksumCase(address string) string {
	a := common.HexToAddress(address)
	return a.Hex()
}
