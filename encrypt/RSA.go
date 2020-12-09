package encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

type RSAData struct {
	Source  []byte
	Result  []byte
	Content []byte
}

func (R *RSAData) Encrypt() (bool, error) {

	block, _ := pem.Decode(R.Content)
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false, err
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, R.Source)
	if err != nil {
		return false, err
	}
	R.Result = encrypted
	return true, nil
}

func (R *RSAData) Decrypt() (bool, error) {
	block, _ := pem.Decode(R.Content)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return false, err
	}
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, R.Source)
	if err != nil {
		return false, err
	}
	R.Result = decrypted
	return true, nil
}
