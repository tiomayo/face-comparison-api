package controller

import (
	"encoding/json"
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

// Identify as endpoint
func Identify(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var identitas Identitas
	var respon Respon

	if err := json.NewDecoder(r.Body).Decode(&identitas); err != nil {
		respon.Status = 404
		respon.Message = "Format salah"
		json.NewEncoder(w).Encode(respon)
		return
	}

	jsonidentitas, err := json.Marshal(identitas)
	if err != nil {
		panic(err)
	}
	w.Write(jsonidentitas)
}
