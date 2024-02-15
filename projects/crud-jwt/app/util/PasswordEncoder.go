package util

import (
	"crud/app/error"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

type PasswordEncoder struct {
	secret string
	bytes  []byte
}

func (passwordEncoder PasswordEncoder) encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func (passwordEncoder *PasswordEncoder) Encrypt(data string) (string, *error.AppError) {
	bytes := passwordEncoder.bytes
	block, err := aes.NewCipher([]byte(passwordEncoder.secret))
	if err != nil {
		return "", error.NewUnexpectedError(err.Error())
	}
	dataInBytes := []byte(data)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(dataInBytes))
	cfb.XORKeyStream(cipherText, dataInBytes)
	return passwordEncoder.encode(cipherText), nil
}
func (passwordEncoder *PasswordEncoder) Match(data, matchingData string) (bool, *error.AppError) {
	encryptedData, err := passwordEncoder.Encrypt(data)
	if err != nil || encryptedData != matchingData {
		return false, err
	}
	return true, nil
}

func Encoder(secret string) *PasswordEncoder {
	return &PasswordEncoder{
		secret: secret,
		bytes:  []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05},
	}
}
