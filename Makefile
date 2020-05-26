.PHONY: all binaries clean bench

ROOT_FLAGS=`root-config --cflags --libs`
OPT=-O2

all: binaries

clean:
	/bin/rm -fr ./bin

binaries:
	mkdir -p ./bin
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-root-read-br ./cxx-root/read-scalar-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-root-read-rd ./cxx-root/read-scalar-rd.cxx
	go build -o ./bin/gen-data-scalar ./cmd/gen-data-scalar
	go build -o ./bin/run-bench ./cmd/run-bench
	go build -o ./bin/read-scalar ./cmd/read-scalar

bench:
	./bin/run-bench -count=20
