package main

import (
	"fmt"
	"image/png"
	"net/http"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func generateQRCode(w http.ResponseWriter, r *http.Request) {
	// Get the data for the QR code from the request
	t := r.URL.Query().Get("data")
	if t == "" {
		http.Error(w, "No data provided", http.StatusBadRequest)
		return
	}

	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil || size <= 0 {
		http.Error(w, "Invalid size provided", http.StatusBadRequest)
		return
	}

	// Generate the QR code
	qrCode, _ := qr.Encode(t, qr.M, qr.Auto)
	qrCode, _ = barcode.Scale(qrCode, size, size)

	// Encode the QR code as a PNG
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, qrCode)
}

func serveHTML(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "index.html")
}

func main() {
	port := "3001"

	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/qr/generate", generateQRCode)

	fmt.Printf("Starting server on port %s\n", port)
	if err := http.ListenAndServe(":" + port, nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
