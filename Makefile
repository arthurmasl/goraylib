dylib:
	@go build -buildmode=c-shared -o resources/library.dylib cmd/library/main.go

plugin:
	@go build -buildmode=plugin -o resources/library.so cmd/library/main.go
	@go build -o ./tmp/main ./cmd/app
	
