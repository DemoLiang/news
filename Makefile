
all:build dockerd dockerp

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

dockerd:
	docker build -t news -f Dockfile .
