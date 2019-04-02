package imagehandler

import (
	"fmt"

	"github.com/tiomayo/face-comparison-api/imagehandler/aws"
	"github.com/tiomayo/face-comparison-api/imagehandler/azure"
)

// Comparator interface to compare two images
type Comparator interface {
	CompareByURL(string, string, chan []byte)
	CompareByImages([]byte, []byte, chan []byte)
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

// AWS using AWS API to compare
type AWS struct{}

// CompareByImages of two images using AWS
func (a AWS) CompareByImages(img1 []byte, img2 []byte, ch chan []byte) {
	res, err := aws.Compare(img1, img2)
	if err != nil {
		fmt.Sprintln(err)
	}
	str := fmt.Sprintf("%f", res)
	finalres := []byte("Confidence: " + str)
	ch <- finalres
}

func (a AWS) Read(img []byte, ch chan []byte) {
	res, err := aws.Read(img)
	if err != nil {
		fmt.Sprintln(err)
	}
	ch <- res
}

// Google using Google Vision API to compare
type Google struct{}

// CompareByURL of two images using Goolge
func (g Google) CompareByURL(img1 string, img2 string) (string, error) {
	return "Comparing using google not implemented", nil
}
