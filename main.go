package main

import (
	// "fmt"

	"image/color"
	"net"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/skip2/go-qrcode"
	"main.go/router"
	// "fyne.io/fyne/v2/widget"
)

func main() {
	var box *fyne.Container
	var port string = ":8080"
	var qr *canvas.Image
	r := router.SetUpRouter()
	a := app.New()
	appname := "PowerClick"
	w := a.NewWindow(appname)
	w.Resize(fyne.NewSquareSize(500))
	ip := getIp()
	portEditText := widget.NewEntry()
	portEditText.SetPlaceHolder("Hangi Portu kullanmak istersiniz ? (seçmezseniz 8080 kullanılacak)")
	color := color.NRGBA{R: 0, G: 255, B: 255, A: 255}
	text := canvas.NewText(ip+port, color)
	text.Alignment = fyne.TextAlign(fyne.TextAlignCenter)
	appName := canvas.NewText(appname, color)
	appName.Alignment = fyne.TextAlign(fyne.TextAlignCenter)

	text.Alignment = fyne.TextAlign(fyne.TextAlignCenter)
	generateQr(ip, port)
	qr = canvas.NewImageFromFile("./qr.png")
	qr.FillMode = canvas.ImageFillOriginal
	button := widget.NewButton("Generate qr", func() {
		if portEditText.Text != "" {
			port = ":" + portEditText.Text
		}
		generateQr(ip, port)
		qr = canvas.NewImageFromFile("./qr.png")
		qr.FillMode = canvas.ImageFillOriginal
		text.Text = ip + port
		text.Alignment = fyne.TextAlign(fyne.TextAlignCenter)
		qr.Refresh()
		text.Refresh()
		box.Refresh()

	})
	startServerButton := widget.NewButton("start", func() {
		if portEditText.Text != "" {
			port = ":" + portEditText.Text
		}
		go r.Run(port)
	})
	linkedinUrl, _ := url.Parse("https://www.linkedin.com/in/emre-aktas-9176a31a6/")
	githubUrl, _ := url.Parse("https://github.com/EmreCogac")
	downloadLink, _ := url.Parse("https://memreaktas.blogspot.com/2024/03/beta-02-update-powerclick-app.html")
	taluyLink, _ := url.Parse("https://www.linkedin.com/company/taluy/?viewAsMember=true")
	linkedin := widget.NewHyperlink("linkedin", linkedinUrl)
	github := widget.NewHyperlink("Github", githubUrl)
	update := widget.NewHyperlink("Update link", downloadLink)
	taluy := widget.NewHyperlink("Taluy", taluyLink)
	box = container.NewVBox(
		appName,
		qr,
		button,
		portEditText,
		text,
		startServerButton,
		linkedin,
		github,
		update,
		taluy,
	)
	w.SetContent(box)

	w.SetFixedSize(false)

	w.ShowAndRun()

}

func generateQr(ip string, port string) {
	_ = qrcode.WriteFile(ip+port, qrcode.High, 500, "./qr.png")
}

func getIp() string {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4.String()
		}
	}
	return " "
}

//label := widget.NewLabel("data")
// button := widget.NewButton("button name", func() {
// 	fmt.Println("deneme")
// })
// url, _ := url.Parse("https://www.linkedin.com/in/emre-aktas-9176a31a6/")
// hyperlink := widget.NewHyperlink("linkedin", url)

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
