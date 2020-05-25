# groot-bench

`groot-bench` gathers a few programs to benchmark the read/write performances of [groot](https://go-hep.org/x/hep/groot) _wrt_ ROOT/C++.

## toy-data generation

```
$> go build ./cmd/gen-data-scalar
$> ./gen-data-scalar.go -h
Usage of gen-data-scalar:
  -cpu-profile string
    	path to the output CPU profile
  -lvl int
    	compression level to use (if any) (default -1)
  -nevts int
    	number of events to generate (default 10000000)
  -o string
    	path to output ROOT file to generate (default "scalar.root")
  -seed uint
    	seed for random number generation (default 1234)
  -t string
    	name of the output ROOT tree to generate (default "tree")
  -zip string
    	compression to use (if any)
```

### no-compression

```
./gen-data-scalar -zip=none -o ./testdata/scalar-none.root
```

### lz4

```
./gen-data-scalar -zip=lz4 -lvl=0 -o ./testdata/scalar-lz4-0.root
./gen-data-scalar -zip=lz4 -lvl=1 -o ./testdata/scalar-lz4-1.root
./gen-data-scalar -zip=lz4 -lvl=6 -o ./testdata/scalar-lz4-6.root
./gen-data-scalar -zip=lz4 -lvl=9 -o ./testdata/scalar-lz4-9.root
```

### zlib

```
./gen-data-scalar -zip=zlib -lvl=0 -o ./testdata/scalar-zlib-0.root
./gen-data-scalar -zip=zlib -lvl=1 -o ./testdata/scalar-zlib-1.root
./gen-data-scalar -zip=zlib -lvl=2 -o ./testdata/scalar-zlib-2.root
./gen-data-scalar -zip=zlib -lvl=3 -o ./testdata/scalar-zlib-3.root
./gen-data-scalar -zip=zlib -lvl=6 -o ./testdata/scalar-zlib-6.root
./gen-data-scalar -zip=zlib -lvl=9 -o ./testdata/scalar-zlib-9.root
```

### zstd

```
./gen-data-scalar -zip=zstd -o ./testdata/scalar-zstd.root
```

## timings

### scalar

- Go-HEP `v0.26.1-0.20200511085556-0f7b59f24c5e`

```
name             s
ReadScalar/LZ4   1.39 ± 1%
ReadScalar/None  1.28 ± 2%
ReadScalar/Zlib  4.11 ± 1%
```
