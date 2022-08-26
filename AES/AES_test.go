package AES

import (
	"crypto/rand"
	"log"
	"testing"
)

func TestAESCBC(t *testing.T){
	key:=make([]byte,32);
	rand.Read(key);
	plaintext:=make([]byte,129*16);
	rand.Read(plaintext);
	log.Println(string(plaintext));
	ciphertext:=Encrypt(key,plaintext);
	log.Println("ciphertext=",string(ciphertext));
	plaintext1:=Decrypt(key,ciphertext);
	log.Println("plaintext",string(plaintext1));
}

func BenchmarkAES(t *testing.B){
	key:=make([]byte,32);
	rand.Read(key);
	plaintext:=make([]byte,1024*1024*16);
	rand.Read(plaintext);
	t.Run("AESEnc", func(b *testing.B) {
		for i:=0;i<b.N;i++{
			Encrypt(key,plaintext);
		}
	})
	ciphertext:=Encrypt(key,plaintext);
	t.Run("AESDec", func(b *testing.B) {
		for i:=0;i<b.N;i++{
			Decrypt(key,ciphertext);
		}
	})
}

func BenchmarkAESEnc(b *testing.B){
	key:=make([]byte,32);
	rand.Read(key);
	plaintext:=make([]byte,1024*1024*16);
	rand.Read(plaintext);
	for i:=0;i<b.N;i++{
		Encrypt(key,plaintext);
	}
}