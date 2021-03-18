package lnwallet

import (
	"github.com/John-Tonny/lnd/input"
	vclutil "github.com/John-Tonny/vclsuite_vclutil"
	"github.com/John-Tonny/vclsuite_vclwallet/wallet/txrules"
)

// DefaultDustLimit is used to calculate the dust HTLC amount which will be
// send to other node during funding process.
func DefaultDustLimit() vclutil.Amount {
	return txrules.GetDustThreshold(input.P2WSHSize, txrules.DefaultRelayFeePerKb)
}
