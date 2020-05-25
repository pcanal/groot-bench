// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"compress/flate"
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"

	"go-hep.org/x/hep/groot"
	"go-hep.org/x/hep/groot/riofs"
	"go-hep.org/x/hep/groot/rtree"
	"golang.org/x/exp/rand"
)

func main() {
	var (
		nevts = flag.Int64("nevts", 1e7, "number of events to generate")
		zip   = flag.String("zip", "", "compression to use (if any)")
		lvl   = flag.Int("lvl", flate.DefaultCompression, "compression level to use (if any)")
		fname = flag.String("f", "scalar.root", "path to output ROOT file to generate")
		tname = flag.String("t", "tree", "name of the output ROOT tree to generate")
		seed  = flag.Uint64("seed", 1234, "seed for random number generation")

		cpuProf = flag.String("cpu-profile", "", "path to the output CPU profile")
	)

	log.SetPrefix("gen-data: ")
	log.SetFlags(0)

	flag.Parse()

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
		// use default
		*zip = "default"
	default:
		log.Fatalf("invalid compression flag %q", *zip)
	}

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

	log.Printf(
		"creating ROOT file with compr=%q, level=%d: %s",
		*zip, *lvl, *fname,
	)
	f, err := groot.Create(*fname, opts...)
	if err != nil {
		log.Fatalf("could not create output ROOT file: %+v", err)
	}
	defer f.Close()

	var evt struct {
		I32 int32
		I64 int64
		F32 float32
		F64 float64
		Str string
	}

	w, err := rtree.NewWriter(f, *tname, rtree.WriteVarsFromStruct(&evt))
	if err != nil {
		log.Fatalf("could not create output ROOT tree: %+v", err)
	}
	defer w.Close()

	log.Printf("-- created tree %q:", w.Name())
	for i, b := range w.Branches() {
		log.Printf("branch[%d]: name=%q, title=%q", i, b.Name(), b.Title())
	}

	rnd := rand.New(rand.NewSource(*seed))
	for i := int64(0); i < *nevts; i++ {
		if i%(*nevts/10) == 0 {
			log.Printf("event %d...", i)
		}
		evt.I32 = rnd.Int31()
		evt.I64 = rnd.Int63()
		evt.F32 = rnd.Float32() * float32(rnd.Int())
		evt.F64 = rnd.Float64() * float64(rnd.Int())
		evt.Str = "evt-" + strconv.Itoa(int(i))
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
