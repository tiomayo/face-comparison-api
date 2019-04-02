package azure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const uriDetect = "https://southeastasia.api.cognitive.microsoft.com/face/v1.0/detect?returnFaceId=true"
const uriVerify = "https://southeastasia.api.cognitive.microsoft.com/face/v1.0/verify"
const apiKey = ""

type faceAttr []struct {
	FaceID string `json:"faceId"`
}

// GetConfidence of two images using Azure Face Verify API
func GetConfidence(imageID1 string, imageID2 string) ([]byte, error) {
	res := new(bytes.Buffer)
	imgStr := `{"faceId1":"` + imageID1 + `","faceId2":"` + imageID2 + `"}`
	req, _ := http.NewRequest("POST", uriVerify, strings.NewReader(imgStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("fail to reach Azure Face Verify API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()
	res.ReadFrom(resp.Body)
	return []byte(res.String()), nil
}

// FaceID get face id from Azure Face Detection API using url as a source
func FaceID(source string) (string, error) {
	var res faceAttr
	req, _ := http.NewRequest("POST", uriDetect, strings.NewReader(source))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("fail to reach Azure Face Detection API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res[0].FaceID, nil
}

// GetFaceIDByImage get face id from Azure Face Detection API using image as a source
func GetFaceIDByImage(ImgBytes []byte, ch chan interface{}) error {
	var res faceAttr
	req, _ := http.NewRequest("POST", uriDetect, bytes.NewBuffer(ImgBytes))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{}
	resp, errDo := client.Do(req)
	if errDo != nil {
		fmt.Println("ERROR DO")
		return errors.New("fail to reach Azure Face Detection API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()
	errDecode := json.NewDecoder(resp.Body).Decode(&res)
	if errDecode != nil {
		fmt.Println("ERROR DECODE", resp.Body)
		return errDecode
	}

	fmt.Println("DIBAWAH INI REPOSNSE")
	fmt.Println(resp)

	ch <- res[0].FaceID

	return nil
}

// FaceIDByImage get face id from Azure Face Detection API using image as a source
func FaceIDByImage(ImgBytes []byte) (string, error) {
	var res faceAttr
	req, _ := http.NewRequest("POST", uriDetect, bytes.NewBuffer(ImgBytes))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", errors.New("fail to reach Azure Face Detection API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	fmt.Println("DIBAWAH INI REPOSNSE FaceIDByImage")
	fmt.Println(res[0])
	return res[0].FaceID, nil
}

type confidenceRes struct {
	IsIdentical bool `json:"is_identical"`
	Confidence  bool `json:"confidence"`
}

// GetConfidenceRes of two images using Azure Face Verify API
func GetConfidenceRes(imageID1 string, imageID2 string) (interface{}, error) {
	res := new(bytes.Buffer)
	imgStr := `{"faceId1":"` + imageID1 + `","faceId2":"` + imageID2 + `"}`
	req, _ := http.NewRequest("POST", uriVerify, strings.NewReader(imgStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("fail to reach Azure Face Verify API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()
	res.ReadFrom(resp.Body)
	return res, nil
}
