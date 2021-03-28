FROM gbaeke/gocv-4.0.0-build as build       
RUN go get -u -d gocv.io/x/gocv       
RUN go get -u -d github.com/disintegration/imaging       
RUN go get -u -d github.com/thebasher21/face_blur_detect_api
RUN cd $GOPATH/src/github.com/thebasher21/face_blur_detect_api && go build -o $GOPATH/bin/face_blur_detect_api ./main.go       
                
FROM gbaeke/gocv-4.0.0-run       
COPY --from=build /go/bin/face_blur_detect_api /face_blur_detect_api
#ADD haarcascade_frontalface_default.xml /       

ENTRYPOINT ["/face_blur_detect_api.exe"]