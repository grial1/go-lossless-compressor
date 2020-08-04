BUILD=compressor

.PHONY: $(BUILD).bin

all: clean $(BUILD)

$(BUILD): $(BUILD).bin
	go build -o $<

clean:
	rm -rf *.bin

run: $(BUILD).bin
	./$< $(ARGS)