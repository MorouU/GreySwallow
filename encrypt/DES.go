package encrypt

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"io"
)

type DESData struct {
	Key    []byte
	Source []byte
	Result []byte
	Method string
}

func (D *DESData) GetKey(key []byte) (bool, error) {
	if len(key) == 8 {
		D.Key = key
		return true, nil
	}
	return false, errors.New("Key is not correct~")
}

func (D *DESData) Encrypt() (bool, error) {
	switch D.Method {
	case "CBC":
		return D.EncryptCBC()
	case "ECB":
		return D.EncryptECB()
	case "CFB":
		return D.EncryptCFB()
	}
	return false, errors.New("Error method~")
}

func (D *DESData) Decrypt() (bool, error) {
	switch D.Method {
	case "CBC":
		return D.DecryptCBC()
	case "ECB":
		return D.DecryptECB()
	case "CFB":
		return D.DecryptCFB()
	}
	return false, errors.New("Error method~")
}

// CBC
func (D *DESData) EncryptCBC() (bool, error) {
	block, err := des.NewCipher(D.Key)
	if err != nil {
		return false, err
	}
	originData := pkcs5Padding(D.Source, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, D.Key)
	encrypted := make([]byte, len(originData))
	blockMode.CryptBlocks(encrypted, originData)
	D.Result = encrypted
	return true, nil
}

func (D *DESData) DecryptCBC() (bool, error) {
	block, err := des.NewCipher(D.Key)
	if err != nil {
		return false, err
	}
	blockMode := cipher.NewCBCDecrypter(block, D.Key)
	originData := make([]byte, len(D.Source))
	blockMode.CryptBlocks(originData, D.Source)
	originData = pkcs5UnPadding(originData)
	D.Result = originData
	return true, nil
}

// ECB
func (D *DESData) EncryptECB() (bool, error) {
	getKey, err := des.NewCipher(func(key []byte) (genKey []byte) {
		genKey = make([]byte, 8)
		copy(genKey, D.Key)
		for i := 8; i < len(key); {
			for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
				genKey[j] ^= D.Key[i]
			}
		}
		return
	}(D.Key))
	if err != nil {
		return false, err
	}
	length := (len(D.Source) + des.BlockSize) / des.BlockSize
	plain := make([]byte, length*des.BlockSize)
	copy(plain, D.Source)
	pad := byte(len(plain) - len(D.Source))
	for i := len(D.Source); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0, getKey.BlockSize(); bs <= len(D.Source); bs, be = bs+getKey.BlockSize(), be+getKey.BlockSize() {
		getKey.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	D.Result = encrypted
	return true, nil
}

func (D *DESData) DecryptECB() (bool, error) {
	getKey, err := des.NewCipher(func(key []byte) (genKey []byte) {
		genKey = make([]byte, 8)
		copy(genKey, D.Key)
		for i := 8; i < len(key); {
			for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
				genKey[j] ^= D.Key[i]
			}
		}
		return
	}(D.Key))
	if err != nil {
		return false, err
	}
	decrypted := make([]byte, len(D.Source))
	for bs, be := 0, getKey.BlockSize(); bs < len(D.Source); bs, be = bs+getKey.BlockSize(), be+getKey.BlockSize() {
		getKey.Decrypt(decrypted[bs:be], D.Source[bs:be])
	}

	if len(decrypted) > 0 {
		decrypted = decrypted[:len(decrypted)-int(decrypted[len(decrypted)-1])]
		D.Result = decrypted
		return true, nil
	} else {
		return false, errors.New("Can't decrypt~")
	}
}

// CFB
func (D *DESData) EncryptCFB() (bool, error) {
	block, err := des.NewCipher(D.Key)

	if err != nil {
		return false, err
	}
	encrypted := make([]byte, des.BlockSize+len(D.Source))
	iv := encrypted[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return false, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[des.BlockSize:], D.Source)
	D.Result = encrypted
	return true, nil
}

func (D *DESData) DecryptCFB() (bool, error) {
	block, err := des.NewCipher(D.Key)
	if err != nil {
		return false, err
	}
	if len(D.Source) < des.BlockSize {
		return false, errors.New("Can't decrypt~")
	}
	iv := D.Source[:des.BlockSize]
	decrypted := D.Source[des.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decrypted, decrypted)
	D.Result = decrypted
	return true, nil

}
