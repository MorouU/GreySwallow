package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

type AESData struct {
	Key    []byte
	Source []byte
	Result []byte
	Method string
}

func (A *AESData) GetKey(key []byte) (bool, error) {
	if len(key) == 16 || len(key) == 24 || len(key) == 32 {
		A.Key = key
		return true, nil
	}
	return false, errors.New("Key is not correct~")
}

func (A *AESData) Encrypt() (bool, error) {
	switch A.Method {
	case "CBC":
		return A.EncryptCBC()
	case "ECB":
		return A.EncryptECB()
	case "CFB":
		return A.EncryptCFB()
	}
	return false, errors.New("Error method~")
}

func (A *AESData) Decrypt() (bool, error) {
	switch A.Method {
	case "CBC":
		return A.DecryptCBC()
	case "ECB":
		return A.DecryptECB()
	case "CFB":
		return A.DecryptCFB()
	}
	return false, errors.New("Error method~")
}

// CBC
func (A *AESData) EncryptCBC() (bool, error) {
	block, err := aes.NewCipher(A.Key) // 16,24,32 -> AES128,AES192,AES256
	if err != nil {
		return false, err
	}
	originData := pkcs5Padding(A.Source, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, A.Key[:block.BlockSize()])
	encrypted := make([]byte, len(originData))
	blockMode.CryptBlocks(encrypted, originData)
	A.Result = encrypted
	return true, nil
}

func (A *AESData) DecryptCBC() (bool, error) {
	block, err := aes.NewCipher(A.Key) // 16,24,32 -> AES128,AES192,AES256
	if err != nil {
		return false, err
	}
	size := block.BlockSize() // Get BlockSize
	blockMode := cipher.NewCBCDecrypter(block, A.Key[:size])
	decrypted := make([]byte, len(A.Source))
	blockMode.CryptBlocks(decrypted, A.Source)
	decrypted = pkcs5UnPadding(decrypted)
	A.Result = decrypted
	return true, nil
}

// ECB
func (A *AESData) EncryptECB() (bool, error) {
	getKey, err := aes.NewCipher(func(key []byte) (genKey []byte) {
		genKey = make([]byte, 16)
		copy(genKey, A.Key)
		for i := 16; i < len(key); {
			for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
				genKey[j] ^= A.Key[i]
			}
		}
		return
	}(A.Key))
	if err != nil {
		return false, err
	}
	length := (len(A.Source) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, A.Source)
	pad := byte(len(plain) - len(A.Source))
	for i := len(A.Source); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, getKey.BlockSize(); bs <= len(A.Source); bs, be = bs+getKey.BlockSize(), be+getKey.BlockSize() {
		getKey.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	A.Result = encrypted
	return true, nil
}

func (A *AESData) DecryptECB() (bool, error) {
	getKey, err := aes.NewCipher(func(key []byte) (genKey []byte) {
		genKey = make([]byte, 16)
		copy(genKey, A.Key)
		for i := 16; i < len(key); {
			for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
				genKey[j] ^= A.Key[i]
			}
		}
		return
	}(A.Key))
	if err != nil {
		return false, err
	}
	decrypted := make([]byte, len(A.Source))
	for bs, be := 0, getKey.BlockSize(); bs < len(A.Source); bs, be = bs+getKey.BlockSize(), be+getKey.BlockSize() {
		getKey.Decrypt(decrypted[bs:be], A.Source[bs:be])
	}

	if len(decrypted) > 0 {
		decrypted = decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]
		A.Result = decrypted
		return true, nil
	} else {
		return false, errors.New("Can't decrypt~")
	}
}

// CFB
func (A *AESData) EncryptCFB() (bool, error) {
	block, err := aes.NewCipher(A.Key)
	if err != nil {
		return false, err
	}
	encrypted := make([]byte, aes.BlockSize+len(A.Source))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return false, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], A.Source)
	A.Result = encrypted
	return true, nil
}

func (A *AESData) DecryptCFB() (bool, error) {
	block, err := aes.NewCipher(A.Key)
	if err != nil {
		return false, err
	}
	if len(A.Source) < aes.BlockSize {
		return false, errors.New("Can't decrypt~")
	}
	iv := A.Source[:aes.BlockSize]
	decrypted := A.Source[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)
	A.Result = decrypted
	return true, nil

}
