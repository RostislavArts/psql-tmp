CODE = psql-tmp.go

TARGET = psql-tmp

build:
	go build $(CODE)

install: build
	sudo mv $(TARGET) /usr/bin/$(TARGET)
