package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Uygulamayı oluşturuyoruz
	myApp := app.New()

	// Pencerenin ismi
	myWindow := myApp.NewWindow("Whois")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(600, 600))

	// İcon için fotoğraf path'i veriyoruz
	icon, err := fyne.LoadResourceFromPath("images/icon.png")
	if err != nil {
		log.Fatal(err)
	}
	myWindow.SetIcon(icon)

	// Arkaplan için fotoğraf path'i veriyoruz
	image2, err := fyne.LoadResourceFromPath("images/YILDIZ.png")
	if err != nil {
		log.Fatal(err)
	}
	backgroundWallpaper := canvas.NewImageFromResource(image2)
	backgroundWallpaper.FillMode = canvas.ImageFillStretch

	// Domain giriş alanı
	domainEntry := widget.NewEntry()
	domainEntry.SetPlaceHolder("Domain adı girin...")

	// Whois sonuçlarını gösterecek Label
	whoisResult := widget.NewLabel("Whois bilgileri burada görünecek")
	whoisResult.Wrapping = fyne.TextWrapWord // Metni sarmalayacak

	// Sosyal medya bağlantıları için ikonlar ve linkler
	linkedinURL, _ := url.Parse("https://www.linkedin.com/in/uslanozan")
	githubURL, _ := url.Parse("https://github.com/uslanozan")

	linkedinLink := widget.NewHyperlink("LinkedIn", linkedinURL)
	githubLink := widget.NewHyperlink("GitHub", githubURL)

	// İkonları yüklüyoruz
	lnkIcon, _ := fyne.LoadResourceFromPath("images/linkedin.png")
	linkedinIcon := canvas.NewImageFromResource(lnkIcon)
	linkedinIcon.SetMinSize(fyne.NewSize(48, 48)) // İkon boyutunu ayarlıyoruz

	gtIcon, _ := fyne.LoadResourceFromPath("images/github.png")
	githubIcon := canvas.NewImageFromResource(gtIcon)
	githubIcon.SetMinSize(fyne.NewSize(48, 48)) // İkon boyutunu ayarlıyoruz

	// Sosyal medya bağlantılarını yan yana koyuyoruz
	linkedinContainer := container.NewHBox(linkedinIcon, linkedinLink)
	githubContainer := container.NewHBox(githubIcon, githubLink)

	socialLinks := container.NewHBox(linkedinContainer, githubContainer)

	// Whois sorgulaması için buton
	whoisButton := widget.NewButton("Whois Sorgula", func() {
		domain := domainEntry.Text
		if domain == "" {
			whoisResult.SetText("Lütfen bir domain adı girin.")
			return
		}
		whoisInfo := getWhoisInfo(domain)
		whoisResult.SetText(whoisInfo)
	})

	// Çıktıyı kaydet butonu
	saveButton := widget.NewButton("Çıktıyı Kaydet", func() {
		domain := domainEntry.Text
		if domain == "" {
			dialog.ShowInformation("Hata", "Lütfen bir domain adı girin.", myWindow)
			return
		}
		saveWhoisInfo(domain, whoisResult.Text)
	})

	// İçerikleri burada oluşturuyoruz
	content := container.NewVBox(
		socialLinks, // Sosyal medya bağlantıları
		domainEntry, // Domain giriş alanı
		whoisButton, // Whois sorgulama butonu
		saveButton,  // Çıktıyı kaydet butonu
		whoisResult, // Whois sonuçları
	)

	// Arka planı ve diğer içerikleri stack'e ekliyoruz
	myWindow.SetContent(container.NewStack(
		backgroundWallpaper, // Arkaplan resmi
		content,             // Diğer içerikler
	))

	myWindow.ShowAndRun()
}

// Whois bilgilerini almak için fonksiyon
func getWhoisInfo(domain string) string {
	conn, err := net.Dial("tcp", "whois.verisign-grs.com:43")
	if err != nil {
		return fmt.Sprintf("Whois sunucusuna bağlanılamadı: %v", err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte(domain + "\r\n"))
	if err != nil {
		return fmt.Sprintf("Whois sorgusu gönderilemedi: %v", err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return fmt.Sprintf("Whois yanıtı alınamadı: %v", err)
	}

	return string(buf[:n])
}

// Whois bilgilerini kaydetmek için fonksiyon
func saveWhoisInfo(domain string, info string) {
	outputsDir := "outputs"
	if err := os.MkdirAll(outputsDir, os.ModePerm); err != nil {
		log.Fatalf("outputs klasörü oluşturulamadı: %v", err)
	}

	domainName := strings.TrimSuffix(domain, ".com")
	fileName := fmt.Sprintf("%s/%s_whois_info.txt", outputsDir, domainName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Dosya oluşturulamadı: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(info)
	if err != nil {
		log.Fatalf("Dosyaya yazılamadı: %v", err)
	}

	dialog.ShowInformation("Başarılı", fmt.Sprintf("%s bilgileri başarıyla kaydedildi.", fileName), fyne.CurrentApp().NewWindow("Success"))
}
