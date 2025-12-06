PROGRAM=grompt

SILENT=go
BUILD_DIR := bin

.PHONY: run build build-static

run:
	$(SILENT) run .


dev:
	env CONFIG_PATH=./config.json $(SILENT) run .

build:
	$(SILENT) build -ldflags "-s -w" -o $(BUILD_DIR)/$(PROGRAM)_light .

build-static:
	env CGO_ENABLED=0 $(SILENT) build -o $(BUILD_DIR)/$(PROGRAM)_static -a -ldflags '-extldflags "-static"' .

release:
	make build build-static
