#include <iostream>

#include "TFile.h"
#include "TTree.h"

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
	auto t = f->Get<TTree>(tname);

	t->SetBranchStatus("*", 1);

	double var00;
	double var01;
	double var02;
	double var03;

	t->Branch("var00", &var00);
	t->Branch("var01", &var01);
	t->Branch("var02", &var02);
	t->Branch("var03", &var03);

	int n = t->GetEntries();
	auto freq = n/10;
	auto sum = 0.0;

	for (int i=0; i<n; i++) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		t->GetEntry(i);
		sum += var00 + var01 + var02 + var03;
	}
	std::cout << "sum=" << sum << "\n";
}

