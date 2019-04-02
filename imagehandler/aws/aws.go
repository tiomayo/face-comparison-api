package aws

import (
	"github.com/aws/aws-sdk-go/aws/credentials"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type Lines struct {
	Text string `json:"DetectedText"`
}

type Res struct {
	DetectedText []*Lines
}

// Compare images using aws API
func Compare(Img1 []byte, Img2 []byte) (float64, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
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
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	svcR := rekognition.New(sess)

	input := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			Bytes: img,
		},
	}

	res, err := svcR.DetectText(input)
	if err == nil && len(res.TextDetections) > 0 {
		finalRes := []byte(`{"DetectedText": "`)
		for _, detectedtext := range res.TextDetections {
			if aws.StringValue(detectedtext.Type) == "LINE" {
				finalRes = append(finalRes, aws.StringValue(detectedtext.DetectedText)+" "...)
			}
		}
		finalRes = append(finalRes, `"}]`...)
		return finalRes, nil
	}

	return nil, err
}
