# groot-bench

`groot-bench` gathers a few programs to benchmark the read/write performances of [groot](https://go-hep.org/x/hep/groot) _wrt_ ROOT/C++.


## input data

- `root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root` (1.8Gb, 23571931 entries)
- toy-data (`float64` and/or `[]float64`)

## results

- Go-HEP `v0.27.0` (Go-1.14), `ROOT-6.20/04` (g++-10.1.0) (both on local file)

```
name                               time/op
ReadCMS/GoHEP/Zlib-8               19.2s ± 1%
ReadCMS/ROOT-TreeBranch/Zlib-8     37.5s ± 1%
ReadCMS/ROOT-TreeReader/Zlib-8     26.1s ± 3%
ReadCMS/ROOT-TreeReaderMT/Zlib-8   25.6s ± 5%  (ROOT::EnableImplicitMT())

ReadScalar/GoHEP/None-8            737ms ± 3%
ReadScalar/GoHEP/LZ4-8             769ms ± 3%
ReadScalar/GoHEP/Zlib-8            1.33s ± 1%
ReadScalar/ROOT-TreeBranch/None-8  1.22s ± 3%
ReadScalar/ROOT-TreeBranch/LZ4-8   1.35s ± 3%
ReadScalar/ROOT-TreeBranch/Zlib-8  2.47s ± 1%
ReadScalar/ROOT-TreeReader/None-8  1.43s ± 5%
ReadScalar/ROOT-TreeReader/LZ4-8   1.57s ± 2%
ReadScalar/ROOT-TreeReader/Zlib-8  2.69s ± 1%
```

- Go-HEP `v0.26.1-0.20200511085556-0f7b59f24c5e`

```
name                     time/op
ReadScalar/GoHEP/None-8  1.27s ± 3%
ReadScalar/GoHEP/LZ4-8   1.39s ± 1%
ReadScalar/GoHEP/Zlib-8  4.10s ± 1%
```

## toy-data

### generation

```
$> make binaries
$> ./bin/gen-data-scalar.go -h
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

#### no-compression

```
./bin/gen-data-scalar -zip=none -o ./testdata/scalar-none.root
```

#### lz4

```
./bin/gen-data-scalar -zip=lz4 -lvl=0 -o ./testdata/scalar-lz4-0.root
./bin/gen-data-scalar -zip=lz4 -lvl=1 -o ./testdata/scalar-lz4-1.root
./bin/gen-data-scalar -zip=lz4 -lvl=6 -o ./testdata/scalar-lz4-6.root
./bin/gen-data-scalar -zip=lz4 -lvl=9 -o ./testdata/scalar-lz4-9.root
```

#### zlib

```
./bin/gen-data-scalar -zip=zlib -lvl=0 -o ./testdata/scalar-zlib-0.root
./bin/gen-data-scalar -zip=zlib -lvl=1 -o ./testdata/scalar-zlib-1.root
./bin/gen-data-scalar -zip=zlib -lvl=2 -o ./testdata/scalar-zlib-2.root
./bin/gen-data-scalar -zip=zlib -lvl=3 -o ./testdata/scalar-zlib-3.root
./bin/gen-data-scalar -zip=zlib -lvl=6 -o ./testdata/scalar-zlib-6.root
./bin/gen-data-scalar -zip=zlib -lvl=9 -o ./testdata/scalar-zlib-9.root
```

#### zstd

```
./bin/gen-data-scalar -zip=zstd -o ./testdata/scalar-zstd.root
```
