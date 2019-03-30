package pii

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ImageStruct standart Struct for image
type ImageStruct struct {
	Data   []byte `json:"data,omitempty"`
	Name   string `json:"name,omitempty"`
	Size   string `json:"size,omitempty"`
	Header string `json:"header,omitempty"`
}

// Pii stands for Personal Identifying Information
type Pii struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nik             string             `json:"nik" bson:"nik"`
	EktpStatus      bool               `json:"ektp_status" bson:"ektp_status"`
	NamaLengkap     string             `json:"nama_lengkap" bson:"nama_lengkap"`
	NoHp            string             `json:"no_hp" bson:"no_hp"`
	TanggalLahir    *time.Time         `json:"tanggal_lahir" bson:"tanggal_lahir"`
	TempatLahir     string             `json:"tempat_lahir" bson:"tempat_lahir"`
	PendidikanAkhir string             `json:"pendidikan_akhir" bson:"pendidikan_akhir"`
	NoKK            string             `json:"no_kk" bson:"no_kk"`
	Alamat          string             `json:"alamat" bson:"alamat"`
	Rt              string             `json:"rt" bson:"rt"`
	Rw              string             `json:"rw" bson:"rw"`
	Kecamatan       string             `json:"kecamatan" bson:"kecamatan"`
	Kabupaten       string             `json:"kabupaten" bson:"kabupaten"`
	Provinsi        string             `json:"provinsi" bson:"provinsi"`
	Agama           string             `json:"agama" bson:"agama"`
	Pekerjaan       string             `json:"pekerjaan" bson:"pekerjaan"`
	StatusKawin     string             `json:"status_kawin" bson:"status_kawin"`
	FotoKTP         *ImageStruct       `json:"foto_ktp" bsono:"foto_ktp"`
}

// GetLocalPii func to get Personal Information based on given nik (param)
// return (Pii, nil) Struct , and (nil, error) if data is not exist or somethong went wrong
func GetLocalPii(nik string) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("facecomparison").Collection("pii")
	cursor, _ := collection.Find(ctx, bson.M{"nik": nik})

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Pii
		cursor.Decode(&person)
		fmt.Println("nik : ", person.Nik)
	}

}
