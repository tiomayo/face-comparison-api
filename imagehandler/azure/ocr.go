package azure

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const uriOCR = "https://westcentralus.api.cognitive.microsoft.com/vision/v2.0/ocr"
const apiKeyOCR = "9e9da3f626c74ad8afda993049b94f81"

type Words struct {
	Text string `json:"text"`
}

type Lines struct {
	Lines []*Words `json:"words"`
}

type Regions struct {
	Lines []*Lines `json:"lines"`
}

type OCRResponseJSON struct {
	Language string     `json:"language"`
	Regions  []*Regions `json:"regions"`
}

// Read function to read identity from images using Azure OCR API
func Read(ImgKTP []byte) ([]byte, error) {
	var res OCRResponseJSON
	req, _ := http.NewRequest("POST", uriOCR, bytes.NewBuffer(ImgKTP))
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", apiKeyOCR)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("fail to reach Azure Face Detection API, may caused by client policy or network connectivity problem")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &res)
	jsonFormatted, _ := json.Marshal(res)
	return jsonFormatted, nil
}
