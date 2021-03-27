package facecheck

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

func FaceCheckReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		type requestStruct struct {
			//auth string `json:"auth"`
			Path string `json:"path" validate:"required"`
		}

		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		var reqObj requestStruct
		err := decoder.Decode(&reqObj)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		//bytes, err := json.Marshal(r.URL.Query())
		if err != nil {
			log.Println(err)
			panic(err)
		}
		//var reqObj map[string][]string

		//err = json.Unmarshal(bytes, &reqObj)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		v := validator.New()

		err = v.Struct(reqObj)

		if err != nil {
			log.Println(err)
			panic(err)
		}

		var attributes = make(map[string]string)
		attributes["path"] = reqObj.Path
		ctx := context.WithValue(r.Context(), "attributes", attributes)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
