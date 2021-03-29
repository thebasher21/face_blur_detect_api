FROM denismakogon/gocv-alpine:4.0.1-buildstage as build-stage

LABEL maintainer="Denis Makogon. mail: lildee1991@gmail.com"

ENV PKG_CONFIG_PATH /usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH /usr/local/lib64
ENV CGO_CPPFLAGS -I/usr/local/include
ENV CGO_CXXFLAGS "--std=c++1z"
ENV CGO_LDFLAGS "-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking"

#RUN go get -u -d gocv.io/x/gocv
#RUN cd $GOPATH/src/gocv.io/x/gocv && go build -o $GOPATH/bin/gocv-version ./cmd/version/main.go

WORKDIR /app

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .


FROM denismakogon/gocv-alpine:4.0.1-runtime

# OpenCV 3.4.1 shared objects from build-stage
COPY --from=build-stage /usr/local/lib /usr/local/lib
COPY --from=build-stage /usr/local/lib64 /usr/local/lib64
COPY --from=build-stage /usr/local/lib64/pkgconfig /usr/local/lib64/pkgconfig
COPY --from=build-stage /usr/local/include/opencv /usr/local/include/opencv
COPY --from=build-stage /usr/local/include/opencv2 /usr/local/include/opencv2

ENV PKG_CONFIG_PATH /usr/local/lib64/pkgconfig
ENV LD_LIBRARY_PATH /usr/local/lib64
ENV CGO_CPPFLAGS -I/usr/local/include
ENV CGO_CXXFLAGS "--std=c++1z"
ENV CGO_LDFLAGS "-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d -lopencv_plot -lopencv_tracking"




#Command to run the executable
#CMD ["./main"]



COPY --from=build-stage /app/main . 
EXPOSE 4312
ENTRYPOINT ["/main"]

#FROM golang:alpine as builder

#RUN go get -u -d gocv.io/x/gocv

#RUN cd $GOPATH/pkg/mod/gocv.io/x/gocv@v0.26.0

#COPY . $GOPATH/src/gocv.io/x/gocv

#WORKDIR $GOPATH/pkg/mod/gocv.io/x/gocv@v0.26.0

#RUN cd $GOPATH/pkg/mod/gocv.io/x/gocv@v0.26.0

#RUN apt-get -y update

#RUN	apt-get -y install unzip wget build-essential cmake curl git libgtk2.0-dev pkg-config libavcodec-dev libavformat-dev libswscale-dev libtbb2 libtbb-dev libjpeg-dev libpng-dev libtiff-dev libdc1394-22-dev

#RUN make download

#RUN make build

#RUN make sudo_install

#RUN go run ./cmd/version/main.go

#RUN make clean

# ENV GO111MODULE=on

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 

#FROM alpine:latest

#RUN apk --no-cache add ca-certificates

#RUN apk add --no-cache tzdata

#WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
#COPY --from=builder /app/main .


# Start a new stage from scratch
#FROM alpine:latest

#RUN apk --no-cache add ca-certificates

#RUN apk add --no-cache tzdata

#WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
#COPY --from=builder /app/main .       

# Expose port 8080 to the outside world
#EXPOSE 4312

#Command to run the executable
#CMD ["./main"]

#RUN mkdir $GOPATH/src/face_blur_detect_api

#WORKDIR $GOPATH/src/face_blur_detect_api

#COPY go.mod .

#COPY go.sum .

#COPY main.go .

#RUN cd $GOPATH/src/face_blur_detect_api && go build -o $GOPATH/bin/face_blur_detect_api ./main.go       

#ENTRYPOINT ["/face_blur_detect_api.exe"]