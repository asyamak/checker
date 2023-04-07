run:
	go run ./cmd/main.go

build:
	go build -o app ./cmd/main.go


dbuild:
	docker image build -t tsarka-api-image .

drun: 
	docker container run -p 9090:9090 -d --name tsarka-api-container tsarka-api-image

dstop:
	docker stop tsarka-api-container

drm: 
	docker rm tsarka-api-container

drim: 
	docker rmi tsarka-api-image

dclear:
	docker system prune -a

test:
	go test -v ./...