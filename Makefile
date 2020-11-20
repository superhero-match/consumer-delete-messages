prepare:
	go mod download

run:
	go build -o bin/main cmd/consumer/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/consumer/main.go
	chmod +x bin/main

dkb:
	docker build -t consumer-delete-messages .

dkr:
	docker run -p "8010:8010" consumer-delete-messages

launch: dkb dkr

cr-log:
	docker logs consumer-delete-messages -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

cr-ssh:
	docker exec -it consumer-delete-messages /bin/bash


PHONY: prepare build dkb dkr launch cr-log db-log es-log cr-ssh db-ssh es-ssh rmc rmi clear