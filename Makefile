build:
	go build .

fmt:
	go fmt .

test:
	go test . -v

clean:
	rm ./AliveCor

run: build
	./AliveCor
