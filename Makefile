INSTALL = /usr/bin/install
INSTALL_DATA = ${INSTALL} -m 644

PROGRAM_NAME = luci

PREFIX ?= /usr/local
BIN_DIR = ${PREFIX}/bin

all: ${PROGRAM_NAME}

${PROGRAM_NAME}:
	@go build -o ${PROGRAM_NAME}

clean:
	@rm ${PROGRAM_NAME}

.PHONY: all clean
