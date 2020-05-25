.PHONY: all

ROOT_FLAGS=`root-config --cflags --libs`

all: bin/cxx-root-read-br bin/cxx-root-read-rd

bin/cxx-root-read-br: cxx-root/read-scalar-br.cxx
	mkdir -p ./bin
	$(CXX) $(ROOT_FLAGS) -o bin/cxx-root-read-br ./cxx-root/read-scalar-br.cxx

bin/cxx-root-read-rd: cxx-root/read-scalar-rd.cxx
	mkdir -p ./bin
	$(CXX) $(ROOT_FLAGS) -o bin/cxx-root-read-rd ./cxx-root/read-scalar-rd.cxx


