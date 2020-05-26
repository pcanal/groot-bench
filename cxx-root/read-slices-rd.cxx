// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <iostream>
#include <stdint.h>
#include <vector>

#include "TFile.h"
#include "TTree.h"
#include "TTreeReader.h"
#include "TTreeReaderArray.h"

int main(int argc, char **argv) {
	auto fname = "./scalar.root";
	auto tname = "tree";

	if (argc > 1) {
		fname = argv[1];
	}
	if (argc > 2) {
		tname = argv[2];
	}
	auto f = TFile::Open(fname);
	auto r = TTreeReader(tname, f);

	TTreeReaderArray<double> var00(r, "var00");
	TTreeReaderArray<double> var01(r, "var01");
	TTreeReaderArray<double> var02(r, "var02");
	TTreeReaderArray<double> var03(r, "var03");

	int n = r.GetEntries();
	auto freq = n/10;
	int i = 0;
	auto sum = 0.0;

	while (r.Next()) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		i++;
		sum += var00[0] + var01[0] + var02[0] + var03[0];
	}
	std::cout << "sum=" << sum << "\n";
}

