package main

import (
	"log"
	"net/http"

	"github.com/thebasher21/face_blur_detect_api/router/facecheck"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/", facecheck.FacecheckRouter)
	log.Println("go-api server started")
	http.ListenAndServe(":4312", r)
}
