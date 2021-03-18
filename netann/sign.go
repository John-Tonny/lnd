package netann

import (
	"fmt"

	"github.com/John-Tonny/lnd/input"
	"github.com/John-Tonny/lnd/lnwallet"
	"github.com/John-Tonny/lnd/lnwire"
	"github.com/John-Tonny/vclsuite_vcld/btcec"
)

// SignAnnouncement signs any type of gossip message that is announced on the
// network.
func SignAnnouncement(signer lnwallet.MessageSigner, pubKey *btcec.PublicKey,
	msg lnwire.Message) (input.Signature, error) {

	var (
		data []byte
		err  error
	)

	switch m := msg.(type) {
	case *lnwire.ChannelAnnouncement:
		data, err = m.DataToSign()
	case *lnwire.ChannelUpdate:
		data, err = m.DataToSign()
	case *lnwire.NodeAnnouncement:
		data, err = m.DataToSign()
	default:
		return nil, fmt.Errorf("can't sign %T message", m)
	}
	if err != nil {
		return nil, fmt.Errorf("unable to get data to sign: %v", err)
	}

	return signer.SignMessage(pubKey, data)
}
