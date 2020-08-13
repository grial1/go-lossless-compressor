# go-lossless-compressor

[![Build Status](http://34.67.15.139:8080/buildStatus/icon?job=go-loseless-compressor%2Fgo-loseless-compressor-build&subject=Build&color=green)](http://34.72.229.215:8080/job/go-loseless-compressor/job/go-loseless-compressor-build/)

[![Build Status](http://34.67.15.139:8080/buildStatus/icon?job=go-loseless-compressor%2Fgo-loseless-compressor-test&subject=Test&color=blue)](http://34.72.229.215:8080/job/go-loseless-compressor/job/go-loseless-compressor-test/)

* Implementation of a lossless compressor for bitmap (PGM&amp;PPM) images (See: https://en.wikipedia.org/wiki/Netpbm#PPM_example)

## Compilation

```$ make```

## Testing

```$ make test```

## Usage
1. Compression
```$ make run ARGS="test.pgm 5"```
or
```$ make run ARGS="test.ppm 16"```
or
```$ ./compressor.bin test.pgm 5```
or
```$ ./compressor.bin test.ppm 16```

2. Decompression
```$ bash scripts/ubuntu_compressor.bash -d test.loco```

## Requirements
1. [Go tools](https://golang.org/dl/)
2. [Docker](https://docs.docker.com/get-docker/)
3. [GNU Make](https://www.gnu.org/software/make/)
