package AES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)


func Encrypt(key []byte, plaintext []byte) []byte{
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	rand.Read(iv);
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:],plaintext)

	return ciphertext;
}

func Decrypt(key []byte,ciphertext []byte) []byte{
	block, _ := aes.NewCipher(key);
	iv := ciphertext[:aes.BlockSize];
	plaintext:=ciphertext[aes.BlockSize:];
	mode :=cipher.NewCBCDecrypter(block,iv);
	mode.CryptBlocks(ciphertext[aes.BlockSize:],plaintext);
	return plaintext;
}