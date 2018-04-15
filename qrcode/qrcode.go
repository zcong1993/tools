package qrcode

import q "github.com/skip2/go-qrcode"

// GenQrcode can generate qrcode image for input string
func GenQrcode(in string) ([]byte, error) {
	return q.Encode(in, q.Medium, 256)
}
