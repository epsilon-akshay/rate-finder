APP_EXECUTABLE= out/currency-converter
MAIN_FILE_PATH = ./main.go
RUN_TEST_COMMAND = go test -tags spec -coverpkg=./... -coverprofile=c.out
RUN_BUILD_COMMAND = go build

test:
	go test ./...

start-server: build
	./$(APP_EXECUTABLE)

build:
	mkdir -p out/
	GO111MODULE=on $(RUN_BUILD_COMMAND) -o $(APP_EXECUTABLE) $(MAIN_FILE_PATH)