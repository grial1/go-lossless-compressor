# go-lossless-compressor
Implementation of a lossless compressor for bitmap (PGM&amp;PPM) images (See: https://en.wikipedia.org/wiki/Netpbm#PPM_example)

## Compilation

```$ make```

## Usage
1. Compression
```$ make run ARGS="test.pgm 5"```
or
```$ ./compressor.bin test.pgm 5```

2. Decompression
```$ bash scripts/ubuntu_compressor.bash -d test.loco```

## Requirements
1. [Go tools](https://golang.org/dl/)
2. [Docker](https://docs.docker.com/get-docker/)
3. [GNU Make](https://www.gnu.org/software/make/)