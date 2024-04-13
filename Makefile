run: build
	@./tat img ./testdata/gopher.png -x 40

build:
	@go build -o tat .
