package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

const secret = "0jejn021hnfd082n-982bnt8923nf8923nfh"

func Encrypt(key, plain string) (string, error) {
	block, err := aes.NewCipher(hash(key))

	if err != nil {
		return "", err
	}

	cipherText, err := preparePayload(plain)

	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plain))

	return fmt.Sprintf("%x", cipherText), nil
}

func Decrypt(key, cipherHex string) (string, error) {
	block, err := aes.NewCipher(hash(key))

	if err != nil {
		return "", err
	}

	cipherText, err := hex.DecodeString(cipherHex)

	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("encyrpt: cipher is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}

func hash(s string) []byte {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hasher.Sum(nil)
}

func preparePayload(s string) ([]byte, error) {
	cipherText := make([]byte, aes.BlockSize+len(s))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	return cipherText, nil
}
