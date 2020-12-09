package lib

type connect struct {
	IP   string
	Port int64
	Type string
}

// Grey struct
type Grey struct {
	Listen                 *connect
	Connect                *connect
	EncryptMethod          string
	EncryptMode            string
	EncryptSource          string
	EncryptSourceB64Decode bool
	EncryptKey             []byte
	EncryptFrom            string
	TurnMethod             string
	AcceptMethod           string
}
