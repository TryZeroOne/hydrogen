package protection

import (
	"bytes"
	"compress/flate"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"io"
	"os"

	"github.com/andybalholm/brotli"
)

func Encryption(output string) error {

	_plain, err := os.ReadFile(output)

	if err != nil {
		return errors.New("Can't read: " + output)
	}

	plain, err := CompressPlain(_plain)

	if err != nil {
		return err
	}

	payload, err := generatePayload(output, plain)

	if err != nil {
		return err
	}

	file, err := os.Open(output)
	if err != nil {
		return errors.New("Can't read: " + output)
	}
	defer file.Close()

	file.WriteString(payload)

	return nil
}

func generatePayload(output string, plain []byte) (string, error) {
	content, _ := os.ReadFile(output)

	key := sha512.Sum512_256(content)

	c, err := aes.NewCipher(key[:])
	if err != nil {
		return "", errors.New("Encryption error: " + err.Error())
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", errors.New("Encryption error: " + err.Error())
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", errors.New("Encryption error: " + err.Error())
	}
	_ciphertext := gcm.Seal(nonce, nonce, plain, nil)

	for i := range _ciphertext {
		_ciphertext[i] = reverseBytes(_ciphertext[i])
	}

	ciphertext := string(reverseArrayBytes(_ciphertext))

	return ciphertext, nil

}

func reverseBytes(bytes byte) (res byte) {
	for i := 0; i < 8; i++ {
		res <<= 1
		res |= bytes & 1
		bytes >>= 1
	}
	return res
}

func reverseArrayBytes(input []byte) []byte {
	var res = make([]byte, 0)

	for i := range input {
		x := input[len(input)-1-i]
		res = append(res, x)
	}

	return res
}

func compressFlate(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w, err := flate.NewWriter(&b, 9)
	if err != nil {
		return nil, err
	}
	w.Write(data)
	w.Close()
	return b.Bytes(), nil
}

func compressBrotli(data []byte) []byte {
	var b bytes.Buffer
	w := brotli.NewWriterLevel(&b, brotli.BestCompression)
	w.Write(data)
	w.Close()
	return b.Bytes()
}

func CompressPlain(plain []byte) ([]byte, error) {
	res, err := compressFlate(plain)
	if err != nil {
		return nil, errors.New("Encryption error: " + err.Error())
	}

	res = compressBrotli(res)

	return []byte(base64.RawStdEncoding.EncodeToString(res)), nil

}
