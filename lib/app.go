package lib

import (
	"strconv"
	"strings"

	"github.com/morouu/utils"
)

var (
	// Con get Grey
	Con = &Grey{}
)

var (
	// Error
	ListenError         = "Please input Listen format\n -listen (TCP/UDP):(IP):(PORT)"
	ConnectError        = "Please input Connect format\n -connect (TCP/UDP):(IP):(PORT)"
	MethodError         = "Please use (TCP/UDP) , not %s"
	IPError             = "Please use correct IP , not %s"
	PortError           = "Please use correct Port(0~65535) , not %d"
	EncryptMehtodError  = "Please use correct EncryptMethod(DES/AES/RSA/NULL) , not %s"
	EncryptModeError    = "Please use correct EncryptMode(CBC/ECB/CFB) , not %s"
	EncryptKeyFromError = "Please use correct EncryptKeyFrom(string/file/URL/TCP) , not %s"
	B64DecodeKeyError   = "Please use correct B64DecodeKey(True/False) , not %v"
	TurnMethodError     = "Please use correct TurnMethod(en/de/null) , not %s"
	AcceptMethodError   = "Please use correct AcceptMethod(en/de/null) , not %s"
	URLMethodError      = "Please use correct URLMethod(GET/POST) , not %s"
	TcpAddressError     = "Please use correct TcpAddress(IP:PORT) , not %s"
)

var (
	B64DecodeFailed = "Failed to decode Base64 , < %s >"
	FileOpenFailed  = "Failed to open file <%s>"
)

func ParseValue(data map[string]interface{}) bool {

	var connectData = data["connect"].([]string)
	var listenData = data["listen"].([]string)
	var encryptMethod = data["encryptMethod"].(string)
	var encryptMode = data["encryptMode"].(string)
	var encryptSource = data["encryptSource"].(string)
	var encryptFrom = data["encryptFrom"].(string)
	var b64DecodeKey = data["b64DecodeKey"].(bool)
	var turnMethod = data["turnMethod"].(string)
	var acceptMethod = data["acceptMethod"].(string)
	var urlMethod = data["urlMethod"].(string)
	var urlParams = data["urlParams"].(string)
	var urlData = data["urlData"].(string)

	connectPort, _ := strconv.ParseInt(connectData[2], 10, 64)
	listenPort, _ := strconv.ParseInt(listenData[2], 10, 64)

	CheckPort(connectPort, ConnectError)
	CheckPort(listenPort, ListenError)
	CheckMethod(strings.ToLower(connectData[0]), MethodError)
	CheckMethod(strings.ToLower(listenData[0]), MethodError)
	CheckIP(connectData[1], IPError)
	CheckIP(listenData[1], IPError)
	CheckEncryptMethod(encryptMethod)
	CheckEncryptMode(encryptMode)
	CheckKeyFrom(encryptFrom)
	CheckB64DecodeKey(b64DecodeKey)
	CheckTurnMethod(turnMethod)
	CheckAcceptMethod(acceptMethod)

	Con.Listen = &connect{
		IP:   listenData[1],
		Port: listenPort,
		Type: listenData[0],
	}
	Con.Connect = &connect{
		IP:   connectData[1],
		Port: connectPort,
		Type: connectData[0],
	}

	Con.EncryptMethod = strings.ToUpper(encryptMethod)
	Con.EncryptMode = strings.ToUpper(encryptMode)
	Con.EncryptSource = encryptSource
	Con.EncryptSourceB64Decode = b64DecodeKey
	Con.EncryptFrom = strings.ToLower(encryptFrom)
	Con.TurnMethod = strings.ToLower(turnMethod)
	Con.AcceptMethod = strings.ToLower(acceptMethod)

	if Con.EncryptFrom == "url" {
		CheckUrlMethod(urlMethod)
		utils.Set("urlMethod", urlMethod, true)
		utils.Set("params", urlParams, true)
		utils.Set("data", urlData, true)
	}

	return true

}
