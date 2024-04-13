run: build
	@./tat img ./testdata/gopher.png

build:
	@go build -o tat .
