CC = go
CFLAGS = build -o
OUT_DIR = .env
PROJ_NAME = yup
TEST_FS = ubuntu-20.10-rootfs.tar.xz

build: clean deps test build-local
end-to-end: build pull run
all: build run
enter: build enter
break: build invalid-cmd
help: build help
release: build compress

deps:
	$(CC) get ./... && $(CC) mod tidy

build-local:
	$(CC) $(CFLAGS) $(OUT_DIR)/$(PROJ_NAME)

help:
	$(OUT_DIR)/$(PROJ_NAME) -h

test:
	$(CC) test -v ./...

compress:
	tar --zstd -cf $(OUT_DIR)/$(PROJ_NAME).tar.zst $(OUT_DIR)/$(PROJ_NAME)

pull:
	$(OUT_DIR)/$(PROJ_NAME) -k pull

run:
	$(OUT_DIR)/$(PROJ_NAME) -k run -c echo -a hello

enter:
	$(OUT_DIR)/$(PROJ_NAME) -k run -c bash
	
invalid-cmd:
	$(OUT_DIR)/$(PROJ_NAME) INVALID_CMD echo hello

clean:
	$(RM) $(OUT_DIR)/$(PROJ_NAME) && $(RM) $(OUT_DIR)/$(PROJ_NAME).tar.* && $(RM) $(PROJ_NAME).tar.*

.PHONY: deps, build, test, compress, run, clean, all, release, dry-run, enter, build-local, pull
