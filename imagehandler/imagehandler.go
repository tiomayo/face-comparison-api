package imagehandler

import (
	"fmt"

	"github.com/tiomayo/face-comparison-api/imagehandler/aws"
	"github.com/tiomayo/face-comparison-api/imagehandler/azure"
)

// Comparator interface to compare two images
type Comparator interface {
	Compare(img1 []byte, img2 []byte, ch chan []byte)
	// CompareByURL(string, string, chan []byte)
	// CompareByImages([]byte, []byte, chan []byte)
}

// OCRReader interface to read identity from images
type OCRReader interface {
	Read([]byte, chan []byte)
}

// Azure using Azure API to compare
type Azure struct{}

// CompareByURL of two images URL using Azure
func (a Azure) CompareByURL(img1 string, img2 string, ch chan []byte) {
	faceIDKTP, err := azure.FaceID(img1)
	if err != nil {
		fmt.Sprintln(err)
	}
	faceIDSelfie, err := azure.FaceID(img2)
	if err != nil {
		fmt.Sprintln(err)
	}
	imgJSON, err := azure.GetConfidence(faceIDKTP, faceIDSelfie)
	ch <- imgJSON
}

// CompareByImages of two images using Azure
func (a Azure) CompareByImages(img1 []byte, img2 []byte, ch chan []byte) {
	faceIDKTP, err := azure.FaceIDByImage(img1)
	if err != nil {
		fmt.Sprintln(err)
	}
	faceIDSelfie, err := azure.FaceIDByImage(img2)
	if err != nil {
		fmt.Sprintln(err)
	}
	imgJSON, err := azure.GetConfidence(faceIDKTP, faceIDSelfie)
	ch <- imgJSON
}

func (a Azure) Read(imgktp []byte, ch chan []byte) {
	jsonIdentity, err := azure.Read(imgktp)
	if err != nil {
		fmt.Sprintln(err)
	}
	ch <- jsonIdentity
}

// AwsAdapter adapt AWS function
type AwsAdapter struct {
	Gateway *aws.Gateway
}

// Compare of two images using AWS
func (b *AwsAdapter) Compare(img1 []byte, img2 []byte, ch chan []byte) {
	p := &aws.CompareParam{
		ImgKTP:    img1,
		ImgSelfie: img2,
	}

	if res, err := b.Gateway.Compare(p); err == nil {
		finalres := []byte(`{"Confidence": "` + fmt.Sprintf("%f", res) + `"}`)
		ch <- finalres
	}
}

func (b *AwsAdapter) Read(img []byte, ch chan []byte) {
	if res, err := b.Gateway.Read(img); err == nil {
		ch <- res
	}
}
