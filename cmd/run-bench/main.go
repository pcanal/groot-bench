// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	var (
		n = flag.Int("count", 20, "number of times to run the benchmark")
		t = flag.String("title", "", "title of benchmark")
	)

	log.SetPrefix("bench: ")
	log.SetFlags(0)

	flag.Parse()

	runBench(*t, *n, flag.Args())
}

func runBench(title string, n int, args []string) {
	for i := 0; i < n; i++ {
		run(title, args[0], args[1:])
	}
}

func run(title, exe string, args []string) {
	cmd := exec.Command(exe, args...)
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = ioutil.Discard

	start := time.Now()
	defer func() {
		delta := time.Since(start)
		fmt.Printf(
			"Benchmark%s\t%v\t%v s\n",
			strings.Title(title), delta.Nanoseconds(), delta.Seconds(),
		)
	}()
	err := cmd.Run()
	if err != nil {
		log.Fatalf("could not run: %+v", err)
	}
}
