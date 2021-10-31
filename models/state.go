package models

type RiskPrediction struct {
	Claims0    uint64
	Supply0    uint64
	Claims1    uint64
	Supply1    uint64
	IInvariant int // Invariant
	Kappa      int
	Price      uint64
	Alpha      int
}

type AgentRiskPrediction struct {
	Address    string
	Claims0    uint64
	Supply0    uint64
	Claims1    uint64
	Supply1    uint64
	SuppleFree uint64
}

type BondingCurve struct {
	BondingCurve string
	PubKey       string
	Kappa        int
	Price        uint64
	Alpha        int
}

type GlobalState struct {
}
