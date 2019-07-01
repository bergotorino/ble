package hci

import (
	"bytes"
	"crypto/elliptic"
	"testing"

	"github.com/aead/ecdh"
)

func Test_SMP_Key(t *testing.T) {
	k1, err := GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}
	k2, err := GenerateKeys()
	if err != nil {
		t.Fatal(err)
	}
	p256 := ecdh.Generic(elliptic.P256())
	s1 := p256.ComputeSecret(k1.private, k2.public)
	s2 := p256.ComputeSecret(k2.private, k1.public)

	if !bytes.Equal(s1, s2) {
		t.Fatal()
	}
}
