package lib

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/morouu/encrypt"
	"github.com/morouu/utils"
)

func B64decode(s string) []byte {
	b64de, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		log.Printf(B64DecodeFailed, s)
		os.Exit(3)
	}
	return b64de
}

func FileRead(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		log.Printf(FileOpenFailed, path)
		os.Exit(3)
	}
	defer file.Close()
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	return buf
}

func AESCrypto(buf []byte, key []byte, method string, turn string) []byte {
	a := encrypt.AESData{
		Key:    key,
		Source: buf,
		Method: method,
	}
	switch turn {
	case "en":
		_, err := a.Encrypt()
		if err != nil {
			log.Println(err)
		}
		break
	case "de":
		_, err := a.Decrypt()
		if err != nil {
			log.Println(err)
		}
		break
	}
	return a.Result
}

func DESCrypto(buf []byte, key []byte, method string, turn string) []byte {
	d := encrypt.DESData{
		Key:    key,
		Source: buf,
		Method: method,
	}
	switch turn {
	case "en":
		_, err := d.Encrypt()
		if err != nil {
			log.Println(err)
		}
		break
	case "de":
		_, err := d.Decrypt()
		if err != nil {
			log.Println(err)
		}
		break
	}
	return d.Result
}

func NetURL(url string, method string) []byte {
	switch strings.ToUpper(method) {
	case "GET":
		return func() []byte {
			client := &http.Client{Timeout: time.Second * 30}
			resp, err := client.Get(url + "?" + utils.Get("params").(string))
			if err != nil {
				log.Println(err)
				os.Exit(3)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				os.Exit(3)
			}
			return body
		}()
	case "POST":
		return func() []byte {
			client := &http.Client{Timeout: time.Second * 30}
			resp, err := client.Post(url+"?"+utils.Get("params").(string), "application/x-www-form-urlencoded", strings.NewReader(utils.Get("data").(string)))
			if err != nil {
				log.Println(err)
				os.Exit(3)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				os.Exit(3)
			}
			return body
		}()
	}
	return []byte("")
}

func NetTcp(addr string) []byte {

	address := strings.FieldsFunc(addr, func(r rune) bool { return r == ':' })
	if len(address) != 2 {
		log.Printf(TcpAddressError, addr)
		os.Exit(3)
	}
	ip := address[0]
	port, _ := strconv.ParseInt(address[1], 10, 64)

	if port < 0 || port > 65535 {
		log.Printf(PortError, port)
		os.Exit(3)
	}
	connect, err := net.Dial("TCP", ip)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	defer connect.Close()

	key, err := ioutil.ReadAll(connect)
	if err != nil {
		log.Println(err)
		os.Exit(3)
	}
	return key

}
