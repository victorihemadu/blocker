build:
	@go build -o Users/bigDev/Desktop/blocker

run: build
	@.Users/bigDev/Desktop/blocker

test:
	@go test -v ./... 