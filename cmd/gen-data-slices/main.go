// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"compress/flate"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"strings"

	bench "github.com/go-hep/groot-bench"
	"go-hep.org/x/hep/groot"
	"go-hep.org/x/hep/groot/riofs"
	"go-hep.org/x/hep/groot/rtree"
)

func main() {
	var (
		nevts = flag.Int64("nevts", 1e6, "number of events to generate")
		zip   = flag.String("zip", "", "compression to use (if any)")
		lvl   = flag.Int("lvl", flate.DefaultCompression, "compression level to use (if any)")
		fname = flag.String("o", "f64s.root", "path to output ROOT file to generate")
		tname = flag.String("t", "tree", "name of the output ROOT tree to generate")
		cols  = flag.Int("ncols", 4, "number of columns to generate")
		seed  = flag.Int64("seed", 1234, "seed for random number generation")

		cpuProf = flag.String("cpu-profile", "", "path to the output CPU profile")
	)

	log.SetPrefix("gen-data: ")
	log.SetFlags(0)

	flag.Parse()

	bench.Version()

	var opts []riofs.FileOption
	switch strings.ToLower(*zip) {
	case "lz4":
		opts = append(opts, riofs.WithLZ4(*lvl))
	case "lzma":
		opts = append(opts, riofs.WithLZMA(*lvl))
	case "zlib":
		opts = append(opts, riofs.WithZlib(*lvl))
	case "zstd":
		opts = append(opts, riofs.WithZstd(*lvl))
	case "none":
		opts = append(opts, riofs.WithoutCompression())
	case "", "default":
		*zip = "default"
	default:
		log.Fatalf("invalid compression flag %q", *zip)
	}

	log.Printf(
		"creating ROOT file with compr=%q, level=%d: %s",
		*zip, *lvl, *fname,
	)

	gen(*fname, *tname, opts, *nevts, *seed, *cols, *cpuProf)
}

func gen(fname, tname string, opts []riofs.FileOption, nevts, seed int64, ncols int, cpu string) {
	if cpu != "" {
		prof, err := os.Create(cpu)
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

	f, err := groot.Create(fname, opts...)
	if err != nil {
		log.Fatalf("could not create output ROOT file: %+v", err)
	}
	defer f.Close()

	var (
		nels  int32
		data  = make([][]float64, ncols)
		wvars = make([]rtree.WriteVar, 1, 1+ncols)
	)
	wvars[0] = rtree.WriteVar{
		Name:  "n",
		Value: &nels,
	}
	for i := range data {
		wvars = append(wvars, rtree.WriteVar{
			Name:  fmt.Sprintf("var%02d", i),
			Count: "n",
			Value: &data[i],
		})
	}

	w, err := rtree.NewWriter(f, tname, wvars)
	if err != nil {
		log.Fatalf("could not create output ROOT tree: %+v", err)
	}
	defer w.Close()

	log.Printf("-- created tree %q:", w.Name())
	for i, b := range w.Branches() {
		log.Printf("branch[%d]: name=%q, title=%q", i, b.Name(), b.Title())
	}

	rnd := rand.New(rand.NewSource(seed))
	freq := nevts / 10

	for i := int64(0); i < nevts; i++ {
		if i%freq == 0 {
			log.Printf("event %d...", i)
		}
		nels = rnd.Int31n(10) + 1
		for i := range data {
			vs := data[i][:0]
			for j := 0; j < int(nels); j++ {
				vs = append(vs, rnd.Float64())
			}
			data[i] = vs
		}
		_, err = w.Write()
		if err != nil {
			log.Fatalf("could not write event %d: %+v", i, err)
		}
	}

	err = w.Close()
	if err != nil {
		log.Fatalf("could not close tree writer: %+v", err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalf("could not close ROOT file: %+v", err)
	}
}
