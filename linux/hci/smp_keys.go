package hci

import (
	"crypto"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/aead/ecdh"
)

type Keys struct {
	public  crypto.PublicKey
	private crypto.PublicKey
}

func GenerateKeys() (*Keys, error) {
	var err error
	kp := Keys{}
	p256 := ecdh.Generic(elliptic.P256())
	kp.private, kp.public, err = p256.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	return &kp, nil
}
