BIN_FOLDER=./bin
BIN_NAME=app


run: build
	$(BIN_FOLDER)/$(BIN_NAME)

build:
	@go build -o $(BIN_FOLDER)/$(BIN_NAME) .

clean:
	@rm -rf $(BIN_FOLDER)

