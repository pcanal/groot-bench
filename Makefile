.PHONY: all

ROOT_FLAGS=`root-config --cflags --libs`
OPT=-O2

all: bin/cxx-root-read-br bin/cxx-root-read-rd

bin/cxx-root-read-br: cxx-root/read-scalar-br.cxx
	mkdir -p ./bin
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-root-read-br ./cxx-root/read-scalar-br.cxx

bin/cxx-root-read-rd: cxx-root/read-scalar-rd.cxx
	mkdir -p ./bin
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-root-read-rd ./cxx-root/read-scalar-rd.cxx


