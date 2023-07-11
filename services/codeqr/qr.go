package codeqr

import (
	"image"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func NewQRCode(text string, size int) image.Image {
	qr, err := qrcode.New(text, qrcode.Highest)
	if err != nil {
		log.Println(err)
	}
	qrImg := qr.Image(size)
	return qrImg
}
