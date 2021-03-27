package main

import (
	"face_blur_detect_api/router/facecheck"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	/*
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("health ok"))
		})
	*/
	r.Route("/", facecheck.FacecheckRouter)
	log.Println("go-api server started")
	http.ListenAndServe(":4312", r)
}
