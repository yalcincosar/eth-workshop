package wallet

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func GenerateWallet() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	encodedPrivateKey := hexutil.Encode(privateKeyBytes)[2:]

	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	walletAddress := hexutil.Encode(hash.Sum(nil)[12:])

	return encodedPrivateKey, walletAddress
}
