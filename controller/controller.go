package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

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

type Attr []struct {
	FaceID string `json:"faceId"`
}

// Identify as endpoint
func Identify(w http.ResponseWriter, r *http.Request) {
	var identitas Identitas
	var respon Respon

	// Mock picture as url
	urlgambar := `{"url":"http://cdn2.tstatic.net/batam/foto/bank/images/cut-tari-artis-dan-pembawa-acara-televisi.jpg"}`
	urlgambar2 := `{"url":"https://cdns.klimg.com/kapanlagi.com/selebriti/Ersa_Mayori/p/ersa-mayori-025.jpg"}`
	// faceID1, _ := json.Marshal(GetFaceID(urlgambar))
	faceID1 := GetFaceID(urlgambar)
	faceID2 := GetFaceID(urlgambar2)

	// Insert faceID into array
	jsonFaceID := []byte(`{"faceId1":"` + faceID1[0].FaceID + `","faceId2":"` + faceID2[0].FaceID + `"}`)

	// get Verify
	jsonVerify, _ := json.Marshal(Verify(jsonFaceID))

	if err := json.NewDecoder(r.Body).Decode(&identitas); err != nil {
		respon.Status = 404
		respon.Message = "Format salah"
		json.NewEncoder(w).Encode(respon)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(jsonVerify)
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

// GetFaceID func hit azure api and return faceId
func GetFaceID(ImgStr string) Attr {
	var attributes Attr
	jsonimg := []byte(ImgStr)
	url := fmt.Sprintf("https://westcentralus.api.cognitive.microsoft.com/face/v1.0/detect?returnFaceId=true")
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonimg))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "b7c28d1c5a4d411f9a83c28412498f9b")

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&attributes)
	return attributes
}

// Verify function
func Verify(jsonStr []byte) (data interface{}) {
	// jsonStr = []byte(jsonStr)
	url := fmt.Sprint("https://westcentralus.api.cognitive.microsoft.com/face/v1.0/verify")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Printf("NewRequest: ", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "b7c28d1c5a4d411f9a83c28412498f9b")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Do: ", err)
		return
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}
