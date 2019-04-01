package pii

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/tiomayo/face-comparison-api/pii/piimage"

	"github.com/gorilla/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Pii stands for Personal Identifying Information
type Pii struct {
	ID                primitive.ObjectID   `schema:"_id,omitempty" bson:"_id,omitempty"`
	Nik               string               `schema:"nik,omitempty" bson:"nik,omitempty"`
	EktpStatus        bool                 `schema:"ektp_status,omitempty" bson:"ektp_status,omitempty"`
	NamaLengkap       string               `schema:"nama_lengkap,omitempty" bson:"nama_lengkap,omitempty"`
	NoHp              string               `schema:"no_hp,omitempty" bson:"no_hp,omitempty"`
	TanggalLahir      string               `schema:"tanggal_lahir,omitempty" bson:"tanggal_lahir,omitempty"`
	TempatLahir       string               `schema:"tempat_lahir,omitempty" bson:"tempat_lahir,omitempty"`
	PendidikanAkhir   string               `schema:"pendidikan_akhir,omitempty" bson:"pendidikan_akhir,omitempty"`
	NoKK              string               `schema:"no_kk,omitempty" bson:"no_kk,omitempty"`
	Alamat            string               `schema:"alamat,omitempty" bson:"alamat,omitempty"`
	Rt                string               `schema:"rt,omitempty" bson:"rt,omitempty"`
	Rw                string               `schema:"rw,omitempty" bson:"rw,omitempty"`
	Kecamatan         string               `schema:"kecamatan,omitempty" bson:"kecamatan,omitempty"`
	Kabupaten         string               `schema:"kabupaten,omitempty" bson:"kabupaten,omitempty"`
	Provinsi          string               `schema:"provinsi,omitempty" bson:"provinsi,omitempty"`
	Agama             string               `schema:"agama,omitempty" bson:"agama,omitempty"`
	Pekerjaan         string               `schema:"pekerjaan,omitempty" bson:"pekerjaan,omitempty"`
	StatusPerkawinan  string               `schema:"status_perkawinan,omitempty" bson:"status_perkawinan,omitempty"`
	FotoKTP           *piimage.ImageStruct `schema:"foto_ktp,omitempty" bson:"foto_ktp,omitempty"`
	FotoSelfie        *piimage.ImageStruct `schema:"foto_selfie,omitempty" bson:"foto_selfie,omitempty"`
	FotoSelfieWithKTP *piimage.ImageStruct `schema:"foto_selfie_with_ktp,omitempty" bson:"foto_selfie_with_ktp,omitempty"`
	PasfotoKTP        *piimage.ImageStruct `schema:"pasfoto_ktp,omitempty" bson:"pasfoto_ktp,omitempty"`
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

// DecodeFormPost decode the formPost data in requests form-data and assign it to Pii Struct
func DecodeFormPost(r *http.Request) (*Pii, error) {

	r.ParseMultipartForm(10 << 20)

	fd := r.PostForm
	newPii := new(Pii)
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	parsedtgllahir, errParse := time.Parse("2006-01-02", fd.Get("tanggal_lahir"))
	if errParse != nil {
		return nil, errors.New("Cannot parse tanggallahir decode")
	}

	tgllahirString := parsedtgllahir.String()
	fd.Set("tanggal_lahir", tgllahirString)

	err := decoder.Decode(newPii, fd)

	if err != nil {
		errdetail := fmt.Sprintf("Fail to decode request form-data to new Pii data Struct : %s\n", err)
		return nil, errors.New(errdetail)
	}

	newPii.FotoKTP, err = piimage.ImageStructHandler("foto_ktp", r)
	newPii.FotoSelfie, err = piimage.ImageStructHandler("foto_selfie", r)
	newPii.FotoSelfieWithKTP, err = piimage.ImageStructHandler("foto_selfie_with_ktp", r)
	newPii.PasfotoKTP, err = piimage.ImageStructHandler("pasfoto_ktp", r)

	return newPii, nil
}

// Save current Personal Identifying Information
func (p *Pii) Save() (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("facecomparison").Collection("pii")
	res, err := collection.InsertOne(ctx, p)
	defer cancel()

	if err != nil {
		return nil, err
	}

	id := res.InsertedID

	return id, nil
}
