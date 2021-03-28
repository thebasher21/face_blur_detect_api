package facecheck

import (
	facecheckCont "github.com/thebasher21/face_blur_detect_api/controller/facecheck"
	facecheckReq "github.com/thebasher21/face_blur_detect_api/requestHandler/facecheck"
	facecheckRes "github.com/thebasher21/face_blur_detect_api/responseHandler/facecheck"

	"github.com/go-chi/chi"
)

func FacecheckRouter(r chi.Router) {
	r.Route("/", FacecheckRoutes)
}

func FacecheckRoutes(r chi.Router) {
	r.With(
		facecheckReq.FaceCheckReq,
		facecheckCont.FaceCheckCont,
	).Post("/facecheck", facecheckRes.FaceCheckRes)
}
