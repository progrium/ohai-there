build/container: build/ohaithere Dockerfile
	docker build -t ohaithere .
	touch build/container

build/ohaithere: *.go
	go build -o build/ohaithere

.PHONY: clean
clean:
	rm -rf build