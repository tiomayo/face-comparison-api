package azure

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Attr as struct
type Attr []struct {
	FaceID string `json:"faceId"`
}

var detectAPI = fmt.Sprintf("https://westcentralus.api.cognitive.microsoft.com/face/v1.0/detect?returnFaceId=true")

// GetFaceIDByURI request an faceId to azure API face detect
func GetFaceIDByURI(ImgStr string) (Attr, error) {

	var attributes Attr

	jsonimg := []byte(ImgStr)
	req, _ := http.NewRequest("POST", detectAPI, bytes.NewBuffer(jsonimg))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ocp-Apim-Subscription-Key", "b7c28d1c5a4d411f9a83c28412498f9b")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("client fail to react Azure API, may caused by client policy or network connectivity problem")
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&attributes)

	return attributes, nil
}

// GetFaceIDByImage request an facId to azure API face detect
func GetFaceIDByImage(ImgBytes []byte) (Attr, error) {

	var attributes Attr

	req, _ := http.NewRequest("POST", detectAPI, bytes.NewBuffer(ImgBytes))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", "b7c28d1c5a4d411f9a83c28412498f9b")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, errors.New("fail to reach Azure API, may caused by client policy or network connectivity problem")
	}

	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&attributes)

	return attributes, nil
}
