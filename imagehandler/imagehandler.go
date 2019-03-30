package imagehandler

import (
	"fmt"

	"github.com/tiomayo/middleware/imagehandler/azure"
)

// Comparator interface to compare two images
type Comparator interface {
	Compare(string, string, chan []byte)
}

// Azure using Azure API to compare
type Azure struct{}

// Compare two images using Azure
func (a Azure) Compare(img1 string, img2 string, ch chan []byte) {
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

// AWS using AWS API to compare
type AWS struct{}

// Compare two images using AWS
func (aws AWS) Compare(img1 string, img2 string) (string, error) {
	return "Comparing using aws not implemented", nil
}

// Google using Google Vision API to compare
type Google struct{}

// Compare two images using Goolge
func (g Google) Compare(img1 string, img2 string) (string, error) {
	return "Comparing using google not implemented", nil
}
