#include <iostream>

#include "TFile.h"
#include "TTreeReader.h"

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

	TTreeReaderValue<double> var00(r, "var00");
	TTreeReaderValue<double> var01(r, "var01");
	TTreeReaderValue<double> var02(r, "var02");
	TTreeReaderValue<double> var03(r, "var03");

	int n = r.GetEntries();
	auto freq = n/10;
	int i = 0;
	auto sum = 0.0;

	while (r.Next()) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		i++;
		sum += *var00 + *var01 + *var02 + *var03;
	}
	std::cout << "sum=" << sum << "\n";
}

