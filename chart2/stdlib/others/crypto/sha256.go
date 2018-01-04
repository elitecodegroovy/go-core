package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"io"
	"encoding/hex"
)

func hashSHA256str(){
	s := "sha256 芳华"

	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Printf("origin: %s, sha256 hash: %x\n", s, bs)
}

func hashSHA256File(filePath string) (string, error){
	var hashValue string
	file, err := os.Open(filePath)
	if err != nil {
		return hashValue, err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return hashValue,  err
	}
	hashInBytes := hash.Sum(nil)
	hashValue = hex.EncodeToString(hashInBytes)
	return hashValue, nil

}

func testHashSHA256File(){
	filePath := "./sha256.go"
	if hash , err := hashSHA256File(filePath); err != nil {
		fmt.Printf(" %s, sha256 value: %s ", filePath,  hash)
	}else {
		fmt.Printf(" %s, sha256 hash: %s ", filePath,  hash)
	}
}
func main(){
	hashSHA256str()
	testHashSHA256File()

	//cmd :
	//>go build -o sha256.exe
	/*
	origin: sha256 芳华, sha256 hash: ddaddfd55d4ce82e5c3e8da4c77d70093f94c89b0b1a282a29edf49395b6b4e9
	./sha256.go, sha256 hash: 02a264bb85843644155801f6ac245ef6304f385df729ac57ec9ed2c215f49cd7

	*/

}
