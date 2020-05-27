// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bench_test

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

func BenchmarkReadScalar(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-scalar",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-scalar-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-scalar-rd",
		},
	} {
		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "None",
					fname: "./testdata/scalar-none.root",
				},
				{
					name:  "LZ4",
					fname: "./testdata/scalar-lz4.root",
				},
				{
					name:  "Zlib",
					fname: "./testdata/scalar-zlib.root",
				},
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadSlices(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-slices",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-slices-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-slices-rd",
		},
	} {
		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "None",
					fname: "./testdata/f64s-none.root",
				},
				{
					name:  "LZ4",
					fname: "./testdata/f64s-lz4.root",
				},
				{
					name:  "Zlib",
					fname: "./testdata/f64s-zlib.root",
				},
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadCMS(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
		skip bool
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-cms",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-cms-br",
		},
		//	{
		//		kind: "ROOT-TreeBranchMT",
		//		cmd:  "./bin/cxx-read-cms-br-mt",
		//		skip: true, // takes way too much time (more than single threaded)
		//	},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-cms-rd",
		},
		//	{
		//		kind: "ROOT-TreeReaderMT",
		//		cmd:  "./bin/cxx-read-cms-rd-mt",
		//		skip: true, // TreeReader isn't multi-threaded (RDF is)
		//	},
	} {
		if lc.skip {
			b.Skip()
		}

		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "Zlib",
					fname: "./testdata/Run2012B_DoubleElectron.root",
				},
				// {
				// 	name:  "Zlib-XRD",
				// 	fname: "root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root",
				// },
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}
