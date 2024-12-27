clear:
	rm -rf ./examples/hot/resources && mkdir ./examples/hot/resources

app: 
	@go run ./examples/hot/cmd/app/main.go

FILENAME = $(shell date +%Y%m%d%H%M%S)
plugin:
	@go build -buildmode=plugin -o examples/hot/resources/$(FILENAME).so examples/hot/cmd/library/library.go

make hot:
	@air

dev:
	make -j2 clear hot app
