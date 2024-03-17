package main

import (
	// "fmt"

	"net"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"github.com/skip2/go-qrcode"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("deneme")
	w.Resize(fyne.NewSquareSize(700))
	//var box *fyne.Container
	//label := widget.NewLabel("data")
	// button := widget.NewButton("button name", func() {
	// 	fmt.Println("deneme")
	// })
	// url, _ := url.Parse("https://www.linkedin.com/in/emre-aktas-9176a31a6/")
	// hyperlink := widget.NewHyperlink("linkedin", url)
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			_ = qrcode.WriteFile(ipv4.String(), qrcode.High, 500, "./qr.png")
			// color := color.NRGBA{R: 0, G: 255, B: 0, A: 255}

			// text := canvas.NewText(ipv4.String(), color)
			// box = container.NewVBox(
			// 	text,

			//)

		}

	}
	images := canvas.NewImageFromFile("./qr.png")
	w.SetContent(images)

	w.SetFixedSize(true)
	w.ShowAndRun()
}

// func getIP(color color.Color) (*canvas.Text, *canvas.Image) {
// 	var textNull *canvas.Text = canvas.NewText("", color)
// 	var imageNull *canvas.Image
// 	host, _ := os.Hostname()
// 	addrs, _ := net.LookupIP(host)
// 	for _, addr := range addrs {
// 		if ipv4 := addr.To4(); ipv4 != nil {
// 			png, err := qrcode.Encode(ipv4.String(), qrcode.High, 500)
// 			if err != nil {
// 				text := canvas.NewText("qrerr", color)
// 				return text, imageNull
// 			}
// 			img := image.NewGray(image.Rect(0, 0, 100, 100))
// 			img.Pix = png
// 			images := canvas.NewImageFromImage(img)
// 			text := canvas.NewText(ipv4.String(), color)
// 			return text, images

// 		}

// 	}
// 	return textNull, imageNull
// }
