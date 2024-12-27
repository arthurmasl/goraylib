app: 
	@go run ./cmd/app

plugin:
	@go build -buildmode=plugin -o resources/library.so cmd/library/main.go

dev:
	make -j2 app air

	
