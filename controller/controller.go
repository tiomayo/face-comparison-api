package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/tiomayo/face-comparison-api/imagehandler/aws"

	"github.com/tiomayo/face-comparison-api/imagehandler"
)

// var urlgambar = `{"url":"http://cdn2.tstatic.net/batam/foto/bank/images/cut-tari-artis-dan-pembawa-acara-televisi.jpg"}`
// var urlgambar2 = `{"url":"https://cdns.klimg.com/kapanlagi.com/selebriti/Ersa_Mayori/p/ersa-mayori-025.jpg"}`

// Identitas struct (Model)
type Identitas struct {
	Nik         string `json:"nik"`
	NamaLengkap string `json:"nama_lengkap"`
	NoHp        string `json:"no_hp"`
}

// Respon struct
type Respon struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Identify as endpoint
func Identify(w http.ResponseWriter, r *http.Request) {
	bufKTP := bytes.NewBuffer(nil)
	bufSelfie := bytes.NewBuffer(nil)
	ch := make(chan []byte, 3)

	// Handle OCR KTP
	imgKTP, headerKTP, err := r.FormFile("imgKTP")
	if err != nil {
		http.Error(w, "Required file not found", http.StatusBadRequest)
		return
	}
	fmt.Println("Reading Image KTP " + headerKTP.Filename)
	defer imgKTP.Close()

	// Handle Selfie image
	imgSelfie, headerSelfie, err := r.FormFile("imgSelfie")
	if err != nil {
		http.Error(w, "Required file not found", http.StatusBadRequest)
		return
	}
	fmt.Println("Reading Image Selfie " + headerSelfie.Filename)
	defer imgSelfie.Close()

	// Write images into buffer byte
	io.Copy(bufKTP, imgKTP)
	io.Copy(bufSelfie, imgSelfie)

	adapter := &imagehandler.AwsAdapter{
		Gateway: &aws.Gateway{
			Region:    "",
			KeyID:     "",
			SecretKey: "",
		},
	}
	go adapter.Compare(bufKTP.Bytes(), bufSelfie.Bytes(), ch)
	go adapter.Read(bufKTP.Bytes(), ch)
	chanVal := <-ch
	chanVal2 := <-ch

	w.Header().Add("content-type", "application/json")
	w.Write(append(chanVal, chanVal2...))
}

// HitDukcapil api hit function
func HitDukcapil() (data interface{}) {
	response, err := http.Get("https://httpbin.org/json")
	if err != nil {
		log.Printf("HTTP request failed with error %s\n", err)
	}
	json.NewDecoder(response.Body).Decode(&data)
	return data
}

// Aisatsu sample get request for testing purpose
func Aisatsu(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Guest"
	}
	log.Printf("こんにちは %s'san\n", name)
	w.Write([]byte(fmt.Sprintf("Hello, %s'san\n", name)))
}
