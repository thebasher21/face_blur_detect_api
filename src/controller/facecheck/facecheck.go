package facecheck

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	cv "gocv.io/x/gocv"
	"gonum.org/v1/gonum/stat"
)

func FaceCheckCont(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attributes := r.Context().Value("attributes").(map[string]string)
		path := attributes["path"]
		resp, err := http.Get(path)
		if err != nil {
			log.Println("Unable to get image from path")
			panic(err)
		}
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Unable to read image")
			panic(err)
		}
		defer resp.Body.Close()
		img, err := cv.IMDecode(data, cv.IMReadAnyColor)
		if err != nil {
			log.Println("Unable to decode image")
			panic(err)
		}
		response := ""
		/*
			window := cv.NewWindow("Image")
			defer window.Close()
			window.IMShow(img)
			window.WaitKey(500)
		*/
		newmat := cv.NewMatWithSize(img.Rows(), img.Cols(), cv.MatTypeCV64FC1)
		defer newmat.Close()
		cv.Sobel(img, &newmat, cv.MatTypeCV64FC1, 2, 2, 3, 1.0, 0.0, cv.BorderDefault)
		imgFloat, err := newmat.DataPtrFloat64()
		if err != nil {
			log.Println(err)
			panic(err)
		}
		//numFaces := 0
		variance := stat.Variance(imgFloat, nil)
		if variance <= 30 {
			response = "Selfie is not clear"
		} else {
			classifier := cv.NewCascadeClassifier()
			defer classifier.Close()
			xmlPath, err := filepath.Abs("classifier.xml")
			if err != nil {
				log.Println("Unable to read classifier xml file")
				panic(err)
			}
			classifier.Load(xmlPath)

			faces := classifier.DetectMultiScale(img)
			if len(faces) == 0 {
				response = "No face detected in selfie"
			}
			//numFaces = len(faces)
		}

		if response == "" {
			response = "success"
		}
		resData := make(map[string]interface{})
		resData["result"] = response
		//resData["faces"] = numFaces
		ctx := context.WithValue(r.Context(), "resData", resData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
