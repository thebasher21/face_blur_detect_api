package facecheck

import (
	facecheckCont "face_blur_detect_api/controller/facecheck"
	facecheckReq "face_blur_detect_api/requestHandler/facecheck"
	facecheckRes "face_blur_detect_api/responseHandler/facecheck"

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
