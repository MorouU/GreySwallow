package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/morouu/lib"
	f "github.com/spf13/pflag"
)

func main() {

	var banner string = `
=================================================
	Grey swallow _V1.05 welcome!
=================================================
`
	var listenData string
	var connectData string
	var encryptMethod string
	var encryptMode string
	var encryptSource string
	var encryptFrom string
	var b64DecodeKey bool
	var turnMethod string
	var acceptMethod string
	var urlMethod string
	var urlParams string
	var urlData string

	f.StringVarP(&listenData, "listen", "l", "", "Listen -> [TCP/UDP:IP:PORT]")
	f.StringVarP(&connectData, "connect", "c", "", "Connect -> [TCP/UDP:IP:PORT]")
	f.StringVarP(&encryptMethod, "encrypt-method", "m", "NULL", "Encrypt-Method -> [AES128/AES192/AES256/DES/RSA/NULL]")
	f.StringVarP(&encryptMode, "encrypt-mode", "M", "CBC", "Encrypt-Mode -> [CBC/ECB/CFB]")
	f.StringVarP(&encryptSource, "encrypt-source", "s", "", "Encrypt-Source -> [KEY/Filename/URL/(IP):(PORT)]")
	f.StringVarP(&encryptFrom, "encrypt-from", "f", "String", "Encrypt-From -> [String/File/URL/TCP]")
	f.BoolVarP(&b64DecodeKey, "source-b64-de", "b", false, "Source-Base64-Decode -> [True/False]")
	f.StringVarP(&turnMethod, "turn-method", "t", "NULL", "Turn-Method -> [EN/DE/NULL]")
	f.StringVarP(&acceptMethod, "accept-method", "a", "NULL", "Accept-Method -> [EN/DE/NULL]")
	f.StringVarP(&urlMethod, "url-method", "r", "GET", "Url-Method -> [GET/POST]")
	f.StringVarP(&urlParams, "url-params", "p", "", "Url-Params -> [QueryString]")
	f.StringVarP(&urlData, "url-data", "d", "", "Url-Data -> [BodyString]")

	fmt.Println(banner)
	f.CommandLine.SortFlags = false
	f.Parse()

	listenGroup := strings.FieldsFunc(listenData, func(r rune) bool { return r == ':' })
	connectGroup := strings.FieldsFunc(connectData, func(r rune) bool { return r == ':' })

	if listenData == "" || len(listenGroup) < 3 {
		log.Println(lib.ListenError)
		os.Exit(3)
	}
	if connectData == "" || len(connectGroup) < 3 {
		log.Println(lib.ConnectError)
		os.Exit(3)
	}

	lib.ParseValue(map[string]interface{}{
		"listen":        listenGroup,
		"connect":       connectGroup,
		"encryptMethod": encryptMethod,
		"encryptMode":   encryptMode,
		"encryptSource": encryptSource,
		"encryptFrom":   encryptFrom,
		"b64DecodeKey":  b64DecodeKey,
		"turnMethod":    turnMethod,
		"acceptMethod":  acceptMethod,
		"urlMethod":     urlMethod,
		"urlParams":     urlParams,
		"urlData":       urlData,
	})

	lib.Con.Run_()
}
