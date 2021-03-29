.PHONY: build run stop clean logs
.SILENT: build run stop clean logs

init:
	docker network create finboxnet

build:
	docker build -t face_blur_detect_api .

run:
	docker run -p 4312:4312 --name face_blur_detect_api --network finboxnet --hostname localhost -d face_blur_detect_api 

stop:
	docker stop $(shell docker ps | grep "face_blur_detect_api" | awk '{ print $$1 }')

clean:
	docker container prune
	kill -9 $(shell lsof -t -i:4312)

logs:
	docker logs $(shell docker ps | grep 'face_blur_detect_api' | awk '{ print $$1 }') -f
