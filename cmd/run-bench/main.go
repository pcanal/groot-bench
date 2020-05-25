// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var (
		count = flag.Int("count", 20, "number of times to run the benchmark")
		bench = flag.String("bench", "", "benchmark to run")
		typ   = flag.String("type", "go", "type of benchmark to run (go/cxx)")
	)

	log.SetPrefix("bench: ")
	log.SetFlags(0)

	flag.Parse()

	runBench(*bench, *count, *typ)
}

func runBench(bench string, n int, typ string) {
	switch bench {
	case "ReadScalar":
		BenchmarkReadScalar(n, typ, "LZ4")
		BenchmarkReadScalar(n, typ, "None")
		BenchmarkReadScalar(n, typ, "Zlib")
	case "ReadScalar/LZ4":
		BenchmarkReadScalar(n, typ, "LZ4")
	case "ReadScalar/None":
		BenchmarkReadScalar(n, typ, "None")
	case "ReadScalar/Zlib":
		BenchmarkReadScalar(n, typ, "Zlib")
	}
}

func BenchmarkReadScalar(n int, typ string, input string) {
	in := strings.ToLower(input)
	fname := fmt.Sprintf("./testdata/scalar-%s.root", in)
	lname := fmt.Sprintf("./testdata/log.read-scalar-%s.txt", in)
	title := fmt.Sprintf("ReadScalar/%s", input)
	switch typ {
	case "go":
		lname = fmt.Sprintf("./testdata/log.%s.read-scalar-%s.txt", typ, in)
		benchmarkReadScalar(n, lname, title, "./bin/read-scalar", "-f", fname)
	case "cxx":
		lname = fmt.Sprintf("./testdata/log.cxx-br.read-scalar-%s.txt", in)
		benchmarkReadScalar(n, lname, title, "./bin/cxx-root-read-br", fname)
		lname = fmt.Sprintf("./testdata/log.cxx-rd.read-scalar-%s.txt", in)
		benchmarkReadScalar(n, lname, title, "./bin/cxx-root-read-rd", fname)

	default:
		panic("invalid type: " + typ)
	}
}

func benchmarkReadScalar(n int, fname, title, exe string, args ...string) {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < n; i++ {
		run(f, title, exe, args...)
	}
}

func run(w io.Writer, title, exe string, args ...string) {
	cmd := exec.Command(exe, args...)
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = ioutil.Discard

	w = io.MultiWriter(w, os.Stdout)

	start := time.Now()
	defer func() {
		delta := time.Since(start)
		fmt.Fprintf(
			w,
			"Benchmark%s\t%v\t%v s\n",
			strings.Title(title), delta.Nanoseconds(), delta.Seconds(),
		)
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("could not run: %+v", err)
	}
}
