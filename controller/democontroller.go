package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tiomayo/face-comparison-api/imagehandler/azure"
	"github.com/tiomayo/face-comparison-api/pii"
)

// SecondEndpoint asdasdasd
func SecondEndpoint(w http.ResponseWriter, r *http.Request) {
	decodedPii, errDecode := pii.DecodeFormPost(r)

	if errDecode != nil {
		log.Println(errDecode)
	}

	piiExist, errExist := decodedPii.Exist()

	if errExist != nil {
		log.Println(errExist)
	}

	asyncProc := make(chan interface{}, 1)

	if piiExist == false {
		go func() {
			id, errSave := decodedPii.Save()
			if errSave != nil {
				log.Println(errSave)
			}
			asyncProc <- id
		}()
	}

	a, _ := azure.FaceIDByImage(decodedPii.PasfotoKTP.Data)
	fmt.Println(a)
	go azure.GetFaceIDByImage(decodedPii.PasfotoKTP.Data, asyncProc)
	go azure.GetFaceIDByImage(decodedPii.FotoSelfie.Data, asyncProc)

	faceID1 := <-asyncProc
	faceID2 := <-asyncProc

	confidence, _ := azure.GetConfidenceRes(fmt.Sprintf("%v", faceID1), fmt.Sprintf("%v", faceID2))

	w.Write([]byte(fmt.Sprintf("%v", confidence)))
}
