BIN_NAME = cnv
INSTALL_PATH = /usr/local/bin

build:
	go build -o $(BIN_NAME) main.go

install: build
	sudo cp $(BIN_NAME) $(INSTALL_PATH)/$(BIN_NAME)
	sudo chmod +x $(INSTALL_PATH)/$(BIN_NAME)

clean:
	rm -f $(BIN_NAME)