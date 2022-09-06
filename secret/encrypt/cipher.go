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
	cipherText, err := preparePayload(plain)

	if err != nil {
		return "", err
	}

	iv := cipherText[:aes.BlockSize]

	stream, err := encryptStream(key, iv)
	if err != nil {
		return "", err
	}

	stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plain))

	return fmt.Sprintf("%x", cipherText), nil

}

// EncryptWriter will return a writer that will write encrypted data into a writer (e.g. file)
func EncryptWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream, err := encryptStream(key, iv)
	if err != nil {
		return nil, err
	}

	n, err := w.Write(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encyrpt: unable to write full iv")
	}

	return &cipher.StreamWriter{S: stream, W: w}, nil
}

func Decrypt(key, cipherHex string) (string, error) {
	cipherText, err := hex.DecodeString(cipherHex)

	if err != nil {
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		return "", errors.New("encyrpt: cipher is too short")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream, err := decryptStream(key, iv)
	if err != nil {
		return "", err
	}

	stream.XORKeyStream(cipherText, cipherText)

	return fmt.Sprintf("%s", cipherText), nil
}

// DecryptReader will return reader that will read encrypted data from reader (e.g. file) and decrypt it
func DecryptReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize)

	n, err := r.Read(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New("encyrpt: unable to read full iv")
	}

	stream, err := decryptStream(key, iv)
	if err != nil {
		return nil, err
	}

	return &cipher.StreamReader{S: stream, R: r}, nil
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

func encryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := aes.NewCipher(hash(key))

	if err != nil {
		return nil, err
	}

	return cipher.NewCFBEncrypter(block, iv), nil
}

func decryptStream(key string, iv []byte) (cipher.Stream, error) {
	block, err := aes.NewCipher(hash(key))

	if err != nil {
		return nil, err
	}

	return cipher.NewCFBDecrypter(block, iv), nil
}
