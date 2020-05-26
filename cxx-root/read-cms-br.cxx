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

	t->Branch("run", &evt.run);
	t->Branch("luminosityBlock", &evt.luminosityBlock);
	t->Branch("event", &evt.event);
	t->Branch("PV_npvs", &evt.PV_npvs);
	t->Branch("PV_x", &evt.PV_x);
	t->Branch("PV_y", &evt.PV_y);
	t->Branch("PV_z", &evt.PV_z);
	t->Branch("nMuon", &evt.nMuon);
	t->Branch("Muon_pt", evt.Muon_pt);
	t->Branch("Muon_eta", evt.Muon_eta);
	t->Branch("Muon_phi", evt.Muon_phi);
	t->Branch("Muon_mass", evt.Muon_mass);
	t->Branch("Muon_charge", evt.Muon_charge);
	t->Branch("Muon_pfRelIso03_all", evt.Muon_pfRelIso03_all);
	t->Branch("Muon_pfRelIso04_all", evt.Muon_pfRelIso04_all);
	t->Branch("Muon_tightId", evt.Muon_tightId);
	t->Branch("Muon_softId", evt.Muon_softId);
	t->Branch("Muon_dxy", evt.Muon_dxy);
	t->Branch("Muon_dxyErr", evt.Muon_dxyErr);
	t->Branch("Muon_dz", evt.Muon_dz);
	t->Branch("Muon_dzErr", evt.Muon_dzErr);
	t->Branch("nElectron", &evt.nElectron);
	t->Branch("Electron_pt", evt.Electron_pt);
	t->Branch("Electron_eta", evt.Electron_eta);
	t->Branch("Electron_phi", evt.Electron_phi);
	t->Branch("Electron_mass", evt.Electron_mass);
	t->Branch("Electron_charge", evt.Electron_charge);
	t->Branch("Electron_pfRelIso03_all", evt.Electron_pfRelIso03_all);
	t->Branch("Electron_dxy", evt.Electron_dxy);
	t->Branch("Electron_dxyErr", evt.Electron_dxyErr);
	t->Branch("Electron_dz", evt.Electron_dz);
	t->Branch("Electron_dzErr", evt.Electron_dzErr);

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

