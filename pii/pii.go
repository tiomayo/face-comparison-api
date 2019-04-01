package pii

import (
	"context"
	"fmt"
	"net/textproto"
	"time"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ImageStruct standart Struct for image
type ImageStruct struct {
	Data   []byte               `schema:"data,omitempty" bson:"data,omitempty"`
	Name   string               `schema:"name,omitempty" bson:"name,omitempty"`
	Size   int64                `schema:"size,omitempty" bson:"size,omitempty"`
	Header textproto.MIMEHeader `schema:"header,omitempty" bson:"header,omitempty"`
}

// Pii stands for Personal Identifying Information
type Pii struct {
	ID               primitive.ObjectID `schema:"_id,omitempty" bson:"_id,omitempty"`
	Nik              string             `schema:"nik,omitempty" bson:"nik,omitempty"`
	EktpStatus       bool               `schema:"ektp_status,omitempty" bson:"ektp_status,omitempty"`
	NamaLengkap      string             `schema:"nama_lengkap,omitempty" bson:"nama_lengkap,omitempty"`
	NoHp             string             `schema:"no_hp,omitempty" bson:"no_hp,omitempty"`
	TanggalLahir     string             `schema:"tanggal_lahir,omitempty" bson:"tanggal_lahir,omitempty"`
	TempatLahir      string             `schema:"tempat_lahir,omitempty" bson:"tempat_lahir,omitempty"`
	PendidikanAkhir  string             `schema:"pendidikan_akhir,omitempty" bson:"pendidikan_akhir,omitempty"`
	NoKK             string             `schema:"no_kk,omitempty" bson:"no_kk,omitempty"`
	Alamat           string             `schema:"alamat,omitempty" bson:"alamat,omitempty"`
	Rt               string             `schema:"rt,omitempty" bson:"rt,omitempty"`
	Rw               string             `schema:"rw,omitempty" bson:"rw,omitempty"`
	Kecamatan        string             `schema:"kecamatan,omitempty" bson:"kecamatan,omitempty"`
	Kabupaten        string             `schema:"kabupaten,omitempty" bson:"kabupaten,omitempty"`
	Provinsi         string             `schema:"provinsi,omitempty" bson:"provinsi,omitempty"`
	Agama            string             `schema:"agama,omitempty" bson:"agama,omitempty"`
	Pekerjaan        string             `schema:"pekerjaan,omitempty" bson:"pekerjaan,omitempty"`
	StatusPerkawinan string             `schema:"status_perkawinan,omitempty" bson:"status_perkawinan,omitempty"`
	FotoKTP          *ImageStruct       `schema:"foto_ktp,omitempty" bson:"foto_ktp,omitempty"`
	FotoSelfie       *ImageStruct       `schema:"foto_selfie,omitempty" bson:"foto_selfie,omitempty"`
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
