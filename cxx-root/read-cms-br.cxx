// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <iostream>
#include <stdint.h>
#include <vector>

#include "TROOT.h"
#include "TFile.h"
#include "TTree.h"

struct Event {
   Int_t           run;
   UInt_t          luminosityBlock;
   ULong64_t       event;
   Int_t           PV_npvs;
   Float_t         PV_x;
   Float_t         PV_y;
   Float_t         PV_z;
   UInt_t          nMuon;
   Float_t         Muon_pt[22];   //[nMuon]
   Float_t         Muon_eta[22];   //[nMuon]
   Float_t         Muon_phi[22];   //[nMuon]
   Float_t         Muon_mass[22];   //[nMuon]
   Int_t           Muon_charge[22];   //[nMuon]
   Float_t         Muon_pfRelIso03_all[22];   //[nMuon]
   Float_t         Muon_pfRelIso04_all[22];   //[nMuon]
   Bool_t          Muon_tightId[22];   //[nMuon]
   Bool_t          Muon_softId[22];   //[nMuon]
   Float_t         Muon_dxy[22];   //[nMuon]
   Float_t         Muon_dxyErr[22];   //[nMuon]
   Float_t         Muon_dz[22];   //[nMuon]
   Float_t         Muon_dzErr[22];   //[nMuon]
   UInt_t          nElectron;
   Float_t         Electron_pt[9];   //[nElectron]
   Float_t         Electron_eta[9];   //[nElectron]
   Float_t         Electron_phi[9];   //[nElectron]
   Float_t         Electron_mass[9];   //[nElectron]
   Int_t           Electron_charge[9];   //[nElectron]
   Float_t         Electron_pfRelIso03_all[9];   //[nElectron]
   Float_t         Electron_dxy[9];   //[nElectron]
   Float_t         Electron_dxyErr[9];   //[nElectron]
   Float_t         Electron_dz[9];   //[nElectron]
   Float_t         Electron_dzErr[9];   //[nElectron]
};

int main(int argc, char **argv) {
	auto fname = "./testdata/Run2012B_DoubleElectron.root";
	auto tname = "Events";

	if (argc > 1) {
		fname = argv[1];
	}
	if (argc > 2) {
		tname = argv[2];
	}

#ifdef GROOT_ENABLE_IMT
	ROOT::EnableImplicitMT();
#endif // GROOT_ENABLE_IMT

	auto f = TFile::Open(fname);
	auto t = f->Get<TTree>(tname);

	t->SetBranchStatus("*", 1);

	Event evt;

	t->SetBranchAddress("run", &evt.run);
	t->SetBranchAddress("luminosityBlock", &evt.luminosityBlock);
	t->SetBranchAddress("event", &evt.event);
	t->SetBranchAddress("PV_npvs", &evt.PV_npvs);
	t->SetBranchAddress("PV_x", &evt.PV_x);
	t->SetBranchAddress("PV_y", &evt.PV_y);
	t->SetBranchAddress("PV_z", &evt.PV_z);
	t->SetBranchAddress("nMuon", &evt.nMuon);
	t->SetBranchAddress("Muon_pt", evt.Muon_pt);
	t->SetBranchAddress("Muon_eta", evt.Muon_eta);
	t->SetBranchAddress("Muon_phi", evt.Muon_phi);
	t->SetBranchAddress("Muon_mass", evt.Muon_mass);
	t->SetBranchAddress("Muon_charge", evt.Muon_charge);
	t->SetBranchAddress("Muon_pfRelIso03_all", evt.Muon_pfRelIso03_all);
	t->SetBranchAddress("Muon_pfRelIso04_all", evt.Muon_pfRelIso04_all);
	t->SetBranchAddress("Muon_tightId", evt.Muon_tightId);
	t->SetBranchAddress("Muon_softId", evt.Muon_softId);
	t->SetBranchAddress("Muon_dxy", evt.Muon_dxy);
	t->SetBranchAddress("Muon_dxyErr", evt.Muon_dxyErr);
	t->SetBranchAddress("Muon_dz", evt.Muon_dz);
	t->SetBranchAddress("Muon_dzErr", evt.Muon_dzErr);
	t->SetBranchAddress("nElectron", &evt.nElectron);
	t->SetBranchAddress("Electron_pt", evt.Electron_pt);
	t->SetBranchAddress("Electron_eta", evt.Electron_eta);
	t->SetBranchAddress("Electron_phi", evt.Electron_phi);
	t->SetBranchAddress("Electron_mass", evt.Electron_mass);
	t->SetBranchAddress("Electron_charge", evt.Electron_charge);
	t->SetBranchAddress("Electron_pfRelIso03_all", evt.Electron_pfRelIso03_all);
	t->SetBranchAddress("Electron_dxy", evt.Electron_dxy);
	t->SetBranchAddress("Electron_dxyErr", evt.Electron_dxyErr);
	t->SetBranchAddress("Electron_dz", evt.Electron_dz);
	t->SetBranchAddress("Electron_dzErr", evt.Electron_dzErr);

	int n = t->GetEntries();
	auto freq = n/10;
	auto sum = 0;

	for (int i=0; i<n; i++) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		t->GetEntry(i);
		sum += evt.event;
	}
	std::cout << "sum=" << sum << "\n";
}

