package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

// Compare images using aws API
func Compare(Img1 []byte, Img2 []byte) (float64, error) {
	sess, err := session.NewSession()
	svcR := rekognition.New(sess)

	input := &rekognition.CompareFacesInput{
		SimilarityThreshold: aws.Float64(70),
		SourceImage: &rekognition.Image{
			Bytes: Img1,
		},
		TargetImage: &rekognition.Image{
			Bytes: Img2,
		},
	}

	res, err := svcR.CompareFaces(input)
	if err == nil && len(res.FaceMatches) > 0 {
		for _, matchedFace := range res.FaceMatches {
			confidence := *matchedFace.Similarity
			return confidence, nil
		}
	}
	return 0, err
}

// Read text from images
func Read(img []byte) ([]byte, error) {
	sess, err := session.NewSession()
	svcR := rekognition.New(sess)

	input := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: img,
		},
	}

	res, err := svcR.DetectText(input)
	if err == nil && len(res.TextDetections) > 0 {
		var finalRes []byte
		for _, detectedtext := range res.TextDetections {
			finalRes = append(finalRes, *detectedtext.DetectedText...)
		}
		return finalRes, nil
	}

	return nil, err
}
