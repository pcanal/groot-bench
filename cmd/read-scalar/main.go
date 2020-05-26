// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"strings"
	"time"

	bench "github.com/go-hep/groot-bench"
	"go-hep.org/x/hep/groot"
	"go-hep.org/x/hep/groot/rtree"
)

func main() {
	var (
		nevts = flag.Int64("nevts", -1, "number of events to read")
		tname = flag.String("t", "tree", "name of the output ROOT tree to generate")
		names = flag.String("b", "", "comma-separated list of branches/leaves to read")

		cpuProf = flag.String("cpu-profile", "", "path to the output CPU profile")
	)

	log.SetPrefix("groot: ")
	log.SetFlags(0)

	flag.Parse()

	bench.Version()

	if *cpuProf != "" {
		prof, err := os.Create(*cpuProf)
		if err != nil {
			log.Fatalf("could not create CPU profile: %+v", err)
		}
		defer prof.Close()
		err = pprof.StartCPUProfile(prof)
		if err != nil {
			log.Fatalf("could not start CPU profile: %+v", err)
		}
		defer pprof.StopCPUProfile()
	}

	var branches []string
	switch *names {
	case "":
		// no op
	default:
		branches = strings.Split(*names, ",")
	}

	read(flag.Arg(0), *tname, *nevts, branches)
}

func read(fname, tname string, nevts int64, names []string) {
	start := time.Now()
	defer func() {
		log.Printf("read time: %v", time.Since(start))
	}()

	log.Printf("reading ROOT file: %s", fname)
	f, err := groot.Open(fname)
	if err != nil {
		log.Fatalf("could not open ROOT file: %+v", err)
	}
	defer f.Close()

	obj, err := f.Get(tname)
	if err != nil {
		log.Fatalf("could not get ROOT tree: %+v", err)
	}
	t := obj.(rtree.Tree)

	rvars := bench.RVarsFrom(t, names)

	if nevts < 0 {
		nevts = t.Entries()
	}

	r, err := rtree.NewReader(t, rvars, rtree.WithRange(0, nevts))
	if err != nil {
		log.Fatalf("could not create tree reader: %+v", err)
	}
	defer r.Close()

	log.Printf("-- created reader: evts=%d", nevts)
	for i, n := range names {
		log.Printf("branch[%d]: name=%q", i, n)
	}

	var (
		n    = 0
		freq = nevts / 10
	)
	err = r.Read(func(rctx rtree.RCtx) error {
		if rctx.Entry%freq == 0 {
			log.Printf("event %d...", rctx.Entry)
		}
		n++
		return nil
	})
	if err != nil {
		log.Fatalf("could not read tree: %+v", err)
	}
	log.Printf("read %d entries", n)
}
