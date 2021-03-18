// +build bitcoind
// +build !notxindex

package lntest

import (
	"github.com/John-Tonny/vclsuite_vcld/chaincfg"
)

// NewBackend starts a bitcoind node with the txindex enabled and returns a
// BitcoindBackendConfig for that node.
func NewBackend(miner string, netParams *chaincfg.Params) (
	*BitcoindBackendConfig, func() error, error) {

	extraArgs := []string{
		"-debug",
		"-regtest",
		"-txindex",
		"-disablewallet",
	}

	return newBackend(miner, netParams, extraArgs)
}
