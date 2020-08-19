TARGET := GoBCDiceGUI
SRCS := $(shell find . -type f -name '*.go')
PKG_SRCS := $(shell find ../../pkg -type f -name '*.go')

$(TARGET): $(SRCS) $(PKG_SRCS)
	go build

.PHONY: clean

clean:
	$(RM) $(TARGET)
