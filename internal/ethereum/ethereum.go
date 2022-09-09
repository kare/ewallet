package ethereum

import (
	"crypto/ecdsa"
	"errors"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// PrivateKeyToPublicKey converts given private key to ECDSA public key.
func PrivateKeyToPublicKey(key string) (*ecdsa.PublicKey, error) {
	key = strings.TrimPrefix(key, "0x")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return nil, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA := publicKey.(*ecdsa.PublicKey)
	return publicKeyECDSA, nil
}

// PrivateKeyToAddress converts given private key to an address in string format.
func PrivateKeyToAddress(key string) (string, error) {
	publicKeyECDSA, err := PrivateKeyToPublicKey(key)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	hexAddress := address.Hex()
	return hexAddress, nil
}

// PrivateKeyToPublic converts given private key to public key in string format.
func PrivateKeyToPublic(key string) (string, error) {
	publicKeyECDSA, err := PrivateKeyToPublicKey(key)
	if err != nil {
		return "", err
	}
	publicKey := PublicKeyToString(publicKeyECDSA)
	return publicKey, nil
}

// PublicKeyToString converts ECDSA public key to hex format string.
func PublicKeyToString(publicKeyECDSA *ecdsa.PublicKey) string {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey := hexutil.Encode(publicKeyBytes)
	return pubKey[4:]
}

// AddressToChecksumCase converts an Ethereum address to checksum case.
func AddressToChecksumCase(address string) (string, error) {
	a := common.HexToAddress(address)
	h := a.Hex()
	const zeroAddress = "0x0000000000000000000000000000000000000000"
	if h == zeroAddress {
		return "", errors.New("ethereum: given address is not an Ethereum address")
	}
	return h, nil
}
