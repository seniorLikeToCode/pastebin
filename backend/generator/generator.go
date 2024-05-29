package generator

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return string(encoded)
}

func GenerateTag(data string) string {
	dataHashBytes := sha256Of(data)
	generatedNumbers := new(big.Int).SetBytes(dataHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumbers)))
	return finalString[:6]
}
