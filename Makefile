run:bin
	./build/bobobox

bin:
	go build -o build/bobobox app/web-service/main.go

dep:
	dep ensure -v -vendor-only