package shortener

import (
	"crypto/sha256"
	"fmt"
	"github.com/itchyny/base58-go"
	"os"
)
//hashing the initilial input
func sha2560f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}
//encodes the sha256 byte
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}
func GenerateShortLink(initialLink string, userId string) string {
	urlHashBytes := sha2560f(initialLink + userId) //combines two arguments and then passes it to the sha2560 functions
	//created big number from hashbytes
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	//apply base58 to the integer and picks first 8 characters
	finalString := base58Encoded([]byte(fmt.Sprintf("%d" generatedNumber)))
	return finalString[:8]
}