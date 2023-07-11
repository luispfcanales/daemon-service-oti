package gui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/luispfcanales/daemon-service-oti/services/codeqr"
)

func GetCanvasQRCode(text string) *canvas.Raster {
	var sizeQR float32 = 300
	urlQR := "https://oti.vercel.app/api/upload?device="
	imageQR := codeqr.NewQRCode(
		fmt.Sprintf("%s%s", urlQR, text),
		int(sizeQR),
	)

	imageSize := fyne.Size{Width: sizeQR, Height: sizeQR}

	qrCanvas := canvas.NewRasterFromImage(imageQR)
	qrCanvas.SetMinSize(imageSize)

	return qrCanvas
}
