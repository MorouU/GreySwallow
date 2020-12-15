package lib

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/morouu/utils"
)

const MessageBlock = 8192

func (G *Grey) GetKey_() {

	if G.EncryptSourceB64Decode == true {
		G.EncryptSource = string(B64decode(G.EncryptSource))
	}

	switch G.EncryptFrom {
	case "string":
		G.EncryptKey = []byte(G.EncryptSource)
		break
	case "file":
		G.EncryptKey = FileRead(G.EncryptSource)
		break
	case "url":
		G.EncryptKey = NetURL(G.EncryptSource, utils.Get("urlMethod").(string))
		break
	case "tcp":
		G.EncryptKey = NetTcp(G.EncryptSource)
		break
	}
}

func (G *Grey) CheckKey_() {
	switch G.EncryptMethod {
	case "AES128":
		G.EncryptKey = G.EncryptKey[:16]
		break
	case "AES192":
		G.EncryptKey = G.EncryptKey[:24]
		break
	case "AES256":
		G.EncryptKey = G.EncryptKey[:32]
		break
	case "DES":
		G.EncryptKey = G.EncryptKey[:8]
		break
	case "RSA":
		G.EncryptKey = G.EncryptKey
		break
	}
}

func (G *Grey) Run_() {
	G.GetKey_()
	G.CheckKey_()
	G.listen_()

}

func (G *Grey) listen_() {
	listen, err := net.Listen(G.Listen.Type, G.Listen.IP+":"+strconv.FormatInt(G.Listen.Port, 10))
	if err != nil {
		log.Println(err)
		return
	}
	defer listen.Close()
	for {
		client, err := listen.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		go G.connect_(client)
	}
}

func (G *Grey) connect_(client net.Conn) {
	connect, err := net.Dial(G.Connect.Type, G.Connect.IP+":"+strconv.FormatInt(G.Connect.Port, 10))
	if err != nil {
		log.Println(err)
		return
	}
	switch G.EncryptMethod {
	case "NULL":
		go io.Copy(connect, client)
		go io.Copy(client, connect)
		break
	case "AES128", "AES192", "AES256":
		if G.TurnMethod != "null" {
			go func(key []byte, mode string, turn string) {
				buf := make([]byte, MessageBlock)

				for {
					size, err := client.Read(buf)
					if err == nil {
						result := AESCrypto(buf[:size], key, mode, turn)
						connect.Write(result)
					}
				}
			}(G.EncryptKey, G.EncryptMode, G.TurnMethod)
		} else {
			go io.Copy(connect, client)
		}
		if G.AcceptMethod != "null" {
			go func(key []byte, mode string, turn string) {
				buf := make([]byte, MessageBlock)

				for {
					size, err := connect.Read(buf)
					if err == nil {
						result := AESCrypto(buf[:size], key, mode, turn)
						client.Write(result)
					}
				}
			}(G.EncryptKey, G.EncryptMode, G.AcceptMethod)
		} else {
			go io.Copy(client, connect)
		}
		break

	case "DES":
		if G.TurnMethod != "null" {
			go func(key []byte, mode string, turn string) {
				buf := make([]byte, MessageBlock)

				for {
					size, err := client.Read(buf)
					if err == nil {
						result := DESCrypto(buf[:size], key, mode, turn)
						connect.Write(result)
					}
				}
			}(G.EncryptKey, G.EncryptMode, G.TurnMethod)
		} else {
			go io.Copy(connect, client)
		}
		if G.AcceptMethod != "null" {
			go func(key []byte, mode string, turn string) {
				buf := make([]byte, MessageBlock)

				for {
					size, err := connect.Read(buf)
					if err == nil {
						result := DESCrypto(buf[:size], key, mode, turn)
						client.Write(result)
					}
				}
			}(G.EncryptKey, G.EncryptMode, G.AcceptMethod)
		} else {
			go io.Copy(client, connect)
		}
		break

	case "RSA":
		if G.TurnMethod != "null" {
			go func(key []byte, turn string) {
				buf := make([]byte, MessageBlock)

				for {
					size, err := client.Read(buf)
					if err == nil {
						result := RSACrypto(buf[:size], key, turn)
						connect.Write(result)
					}
				}
			}(G.EncryptKey, G.TurnMethod)
		} else {
			go io.Copy(connect, client)
		}
		if G.AcceptMethod != "null" {
			go func(key []byte, turn string) {
				buf := make([]byte, MessageBlock)
				for {
					size, err := connect.Read(buf)
					if err == nil {
						result := RSACrypto(buf[:size], key, turn)
						client.Write(result)
					}
				}
			}(G.EncryptKey, G.AcceptMethod)
		} else {
			go io.Copy(client, connect)
		}
	}

}
