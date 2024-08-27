# **Go Whois Desktop App** 🌐

Bu uygulama, kullanıcıların girdiği bir domain adı için Whois bilgilerini sorgulamak amacıyla geliştirilmiştir.
Kullanıcılar,domain adını girdikten sonra ilgili Whois bilgilerini görüntüleyebilir ve bu bilgileri TXT dosyası olarak kaydedebilirler.

![image](https://github.com/user-attachments/assets/59d68a37-9b45-4e4c-9dc3-b1e6269d1294)


## ✨ Özellikler

- Kullanıcı dostu arayüz
- Whois bilgilerini görüntüleme
- TXT dosyası olarak kayıt etme

## 🔧 Kurulum
- go get fyne.io/fyne/v2@latest
- go install fyne.io/fyne/v2/cmd/fyne@latest

### 🖇️ Gereksinimler

- Go 1.18+ sürümü
- Google Search API erişimi
- Go Fyne
- GCC

### 🎯 Adımlar

1. Repoyu klonlayın:
   ```bash
   git clone https://github.com/uslanozan/go-whois-desktop-app.git
   cd go-whois-desktop-app

2. Gerekli bağımlılıkları yükleyin:
   ```bash
   go get fyne.io/fyne/v2
   go install fyne.io/fyne/v2/cmd/fyne@latest

3. Projeyi çalıştırın:
   ```bash
   go run main.go


## 📖 Kullanım
- Domain adı girin.
- "Whois Sorgula" butonuna tıklayın.
- Whois bilgileri görüntülenecek.
- "Çıktıyı Kaydet" butonuna tıklayarak bilgileri TXT dosyası olarak kaydedin.

## 📦 Paketleme
1.  ### Windows 🪟:
    ```bash
    fyne package -os windows -icon images/YILDIZ.png

2. ### macOS 🍎:
     ```bash
     fyne package -os darwin -icon images/YILDIZ.png

3. ### Linux 🐧:
   ```bash
   fyne package -os linux -icon images/YILDIZ.png
