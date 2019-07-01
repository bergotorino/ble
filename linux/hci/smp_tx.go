package hci

func (c *Conn) smpSendPairingRequest() error {
	p := smp.config
	b := []byte{pairingRequest, p.ioCap, p.oobFlag, p.authReq, p.maxKeySz, p.initKeyDist, p.respKeyDist}
	return c.sendSMP(pdu(b))
}

// func (c *Conn) smpSendPublicKey() error {
// 	kp := smp.keys

// 	var i interface{}
// 	i = kp.public

// 	pub, ok := i.(ecdh.Point)
// 	if !ok {
// 		return fmt.Errorf("key error")
// 	}

// 	b := []byte{pairingPublicKey, pub.X, pub.Y}

// 	return nil
// }
