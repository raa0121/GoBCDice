TARGET := GoBCDiceAPI
SRCS := $(shell find . -type f -name '*.go')

$(TARGET): $(SRCS)
	go build

.PHONY: clean

clean:
	$(RM) $(TARGET)
