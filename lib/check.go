package lib

import (
	"log"
	"net"
	"os"
	"strings"
)

func CheckPort(port int64, errString string) {
	if port < 0 || port > 65535 {
		log.Printf(errString, port)
		os.Exit(3)
	}
}

func CheckMethod(method string, errString string) {
	if method != "udp" && method != "tcp" {
		log.Printf(errString, method)
		os.Exit(3)
	}
}

func CheckIP(IP string, errString string) {
	var target net.IP
	if target = net.ParseIP(IP); target == nil {
		log.Printf(errString, IP)
		os.Exit(3)
	} else if target = target.To4(); target == nil {
		log.Printf(errString, IP)
		os.Exit(3)
	}
}

func CheckEncryptMethod(method string) {
	method = strings.ToUpper(method)
	for _, v := range []string{
		"AES128",
		"AES192",
		"AES256",
		"DES",
		"RSA",
		"NULL",
	} {
		if method == v {
			return
		}
	}
	log.Printf(EncryptMehtodError, method)
	os.Exit(3)
}

func CheckEncryptMode(mode string) {
	mode = strings.ToUpper(mode)
	for _, v := range []string{
		"CBC",
		"ECB",
		"CFB",
	} {
		if mode == v {
			return
		}
	}
	log.Printf(EncryptModeError, mode)
	os.Exit(3)
}

func CheckKeyFrom(from string) {
	from = strings.ToLower(from)
	for _, v := range []string{
		"string",
		"file",
		"url",
		"tcp",
	} {
		if from == v {
			return
		}
	}
	log.Printf(EncryptKeyFromError, from)
	os.Exit(3)
}

func CheckB64DecodeKey(b bool) {
	if b != true && b != false {
		log.Printf(B64DecodeKeyError, b)
		os.Exit(3)
	}
}

func CheckTurnMethod(method string) {
	method = strings.ToLower(method)
	if method != "en" && method != "de" && method != "null" {
		log.Printf(TurnMethodError, method)
		os.Exit(3)
	}
}

func CheckAcceptMethod(method string) {
	method = strings.ToLower(method)
	if method != "en" && method != "de" && method != "null" {
		log.Printf(AcceptMethodError, method)
		os.Exit(3)
	}
}

func CheckUrlMethod(method string) {
	method = strings.ToUpper(method)
	if method != "GET" && method != "POST" {
		log.Printf(URLMethodError, method)
		os.Exit(3)
	}
}
