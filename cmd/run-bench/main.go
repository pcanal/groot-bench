// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"golang.org/x/perf/benchstat"
)

func main() {
	var (
		count   = flag.Int("count", 20, "number of times to run the benchmark")
		bench   = flag.String("bench", ".", "benchmark to run")
		fname   = flag.String("log", "", "path to log file")
		timeout = flag.Duration("timeout", 2*time.Hour, "benchmark timeout")
	)

	log.SetPrefix("bench: ")
	log.SetFlags(0)

	flag.Parse()

	if *fname == "" {
		*fname = fmt.Sprintf("./testdata/log-%s.txt", time.Now().Format("2006-01-02-150405"))
	}

	log.Printf("storing bench data into: %s", *fname)
	f, err := os.Create(*fname)
	if err != nil {
		log.Fatalf("could not create output log file: %+v", err)
	}
	defer f.Close()

	o := new(strings.Builder)
	err = run(io.MultiWriter(o, f), *bench, *count, *timeout)
	if err != nil {
		log.Fatalf("could not run bench: %+v", err)
	}

	fmt.Printf("\n\n")
	log.Printf("===== benchstat =====")
	bstat(*fname, strings.NewReader(o.String()))
}

func bstat(fname string, r io.Reader) {
	c := benchstat.Collection{
		Alpha:      0.05,
		AddGeoMean: false,
		DeltaTest:  benchstat.UTest,
	}
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c.AddFile(fname, f)
	benchstat.FormatText(os.Stdout, c.Tables())
}

func run(w io.Writer, bench string, count int, timeout time.Duration) error {
	cmd := exec.Command("go", "test", "-run=NONE",
		"-bench="+bench,
		fmt.Sprintf("-count=%d", count),
		fmt.Sprintf("-timeout=%v", timeout),
	)
	cmd.Stdout = io.MultiWriter(w, os.Stdout)
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not run benchmarks: %w", err)
	}

	return nil
}
