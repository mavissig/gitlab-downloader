F_BUILD=downloader

DIR_BUILD=build

all: build

build:
	go build -o $(DIR_BUILD)/$(F_BUILD) ./cmd/loader/main.go

run:
	$(DIR_BUILD)/$(F_BUILD)

clean:
	rm -rf
# developer functions

debug_run:
	go run ./cmd/loader/main.go