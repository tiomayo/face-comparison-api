package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"ekyc/imagehandler"
)

var urlgambar = `{"url":"http://cdn2.tstatic.net/batam/foto/bank/images/cut-tari-artis-dan-pembawa-acara-televisi.jpg"}`
var urlgambar2 = `{"url":"https://cdns.klimg.com/kapanlagi.com/selebriti/Ersa_Mayori/p/ersa-mayori-025.jpg"}`

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
	ch := make(chan []byte, 3)
	var c imagehandler.Comparator = imagehandler.Azure{}

	go c.Compare(urlgambar, urlgambar2, ch)
	v := <-ch

	w.Header().Add("content-type", "application/json")
	w.Write(v)
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
