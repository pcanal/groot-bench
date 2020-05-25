// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bench provides scripts to benchmark groot performances.
package bench

import (
	"log"

	"go-hep.org/x/hep"
	"go-hep.org/x/hep/groot/rtree"
)

// Version displays the go-hep version used.
func Version() {
	v, _ := hep.Version()
	log.Printf("bench using go-hep version: %v", v)
}

func RVarsFrom(t rtree.Tree, names []string) []rtree.ReadVar {
	var rvars []rtree.ReadVar
	switch len(names) {
	case 0:
		return rtree.NewReadVars(t)
	default:
		all := rtree.NewReadVars(t)
		set := make(map[string]struct{})
		for _, name := range names {
			set[name] = struct{}{}
		}
		for _, rvar := range all {
			if _, ok := set[rvar.Name]; !ok {
				continue
			}
			rvars = append(rvars, rvar)
		}
	}

	return rvars
}
