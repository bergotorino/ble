package hci

import (
	"encoding/hex"
	"fmt"
)

type smpConfig struct {
	ioCap, oobFlag, authReq, maxKeySz, initKeyDist, respKeyDist byte
}

type smpState struct {
	config smpConfig
	keys   *Keys
}

var smp *smpState

func SmpInit() error {
	c := smpConfig{
		ioCap:       0x03,          //no input/output
		oobFlag:     0,             //no oob
		authReq:     (1<<0 | 1<<3), //bond+sc
		maxKeySz:    16,
		respKeyDist: 1,
	}

	k, err := GenerateKeys()
	if err != nil {
		return err
	}

	smp = &smpState{c, k}
	return nil
}

func smpOnPairingRequest(in pdu) ([]byte, error) {
	if len(in) < 6 {
		return nil, fmt.Errorf("%v, invalid length %v", hex.EncodeToString(in), len(in))
	}

	rx := smpConfig{}
	rx.ioCap = in[0]
	rx.oobFlag = in[1]
	rx.authReq = in[2]
	rx.maxKeySz = in[3]
	rx.initKeyDist = in[4]
	rx.respKeyDist = in[5]

	fmt.Printf("pair req: %+v\n", rx)

	//reply with pairing resp

	return nil, nil
}

func smpOnPairingResponse(in pdu) ([]byte, error) {
	if len(in) < 6 {
		return nil, fmt.Errorf("%v, invalid length %v", hex.EncodeToString(in), len(in))
	}

	rx := smpConfig{}
	rx.ioCap = in[0]
	rx.oobFlag = in[1]
	rx.authReq = in[2]
	rx.maxKeySz = in[3]
	rx.initKeyDist = in[4]
	rx.respKeyDist = in[5]
	fmt.Printf("pair rsp: %+v\n", rx)

	//send pub key

	return nil, nil
}

func smpOnPairingRandom(in pdu) ([]byte, error) {
	return nil, nil
}

func smpOnPairingFailed(in pdu) ([]byte, error) {
	return nil, nil
}

func smpOnSecurityRequest(in pdu) ([]byte, error) {
	if len(in) < 1 {
		return nil, fmt.Errorf("%v, invalid length %v", hex.EncodeToString(in), len(in))
	}

	// do something...
	rx := smpConfig{}
	rx.authReq = in[0]

	fmt.Printf("sec req: %+v\n", rx)

	// if known, encrypt, otherwise pair

	return nil, nil
}
