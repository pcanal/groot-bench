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
					cmd:  "./bin/cxx-root-read-br",
				},
				{
					kind: "ROOT-TreeReader",
					cmd:  "./bin/cxx-root-read-rd",
				},
			} {
				b.Run(lc.kind, func(b *testing.B) {
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
