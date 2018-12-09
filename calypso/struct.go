package calypso

import (
	"crypto/sha256"
	"fmt"

	"github.com/dedis/kyber"
	"github.com/dedis/kyber/suites"
	"github.com/dedis/onet/log"
	"github.com/dedis/onet/network"
	"github.com/ethereum/go-ethereum/common"
)

func init() {
	fmt.Println("Bjorn")
	network.RegisterMessages(CreateLTS{}, CreateLTSReply{},
		DecryptKey{}, DecryptKeyReply{}, LogAddress{}, LogAddressReply{})
}

type suite interface {
	kyber.Group
	kyber.XOFFactory
}

// NewWrite is used by the writer to ByzCoin to encode his symmetric key
// under the collective public key created by the DKG. As this method uses
// `Embed` to encode the key, depending on the key-length more than one point
// is needed to encode the data.
//
// Input:
//   - suite - the cryptographic suite to use
//   - ltsid - the id of the LTS id - used to create the second generator
//   - writeDarc - the id of the darc where this write will be stored
//   - X - the aggregate public key of the DKG
//   - key - the symmetric key for the document - it will be encrypted in this method
//
// Output:
//   - write - structure containing the encrypted key U, Cs and the NIZKP of
//   it containing the reader-darc.
func NewWrite(suite suites.Suite, ltsid []byte, policy common.Address, X kyber.Point, key []byte) *Write {
	wr := &Write{LTSID: ltsid}
	r := suite.Scalar().Pick(suite.RandomStream())
	C := suite.Point().Mul(r, X)
	wr.U = suite.Point().Mul(r, nil)

	// Create proof
	for len(key) > 0 {
		kp := suite.Point().Embed(key, suite.RandomStream())
		wr.Cs = append(wr.Cs, suite.Point().Add(C, kp))
		key = key[min(len(key), kp.EmbedLen()):]
	}

	gBar := suite.Point().Mul(suite.Scalar().SetBytes(ltsid), nil)
	wr.Ubar = suite.Point().Mul(r, gBar)
	s := suite.Scalar().Pick(suite.RandomStream())
	w := suite.Point().Mul(s, nil)
	wBar := suite.Point().Mul(s, gBar)
	hash := sha256.New()
	for _, c := range wr.Cs {
		c.MarshalTo(hash)
	}
	wr.U.MarshalTo(hash)
	wr.Ubar.MarshalTo(hash)
	w.MarshalTo(hash)
	wBar.MarshalTo(hash)
	//hash.Write(writeDarc)
	wr.E = suite.Scalar().SetBytes(hash.Sum(nil))
	wr.F = suite.Scalar().Add(s, suite.Scalar().Mul(wr.E, r))
	return wr
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// EncodeKey can be used by the writer to ByzCoin to encode his symmetric
// key under the collective public key created by the DKG.
// As this method uses `Pick` to encode the key, depending on the key-length
// more than one point is needed to encode the data.
//
// Input:
//   - suite - the cryptographic suite to use
//   - X - the aggregate public key of the DKG
//   - key - the symmetric key for the document
//
// Output:
//   - U - the schnorr commit
//   - Cs - encrypted key-slices
func EncodeKey(suite suites.Suite, X kyber.Point, key []byte) (U kyber.Point, Cs []kyber.Point) {
	r := suite.Scalar().Pick(suite.RandomStream())
	C := suite.Point().Mul(r, X)
	log.Lvl4("C:", C.String())
	U = suite.Point().Mul(r, nil)
	log.Lvl4("U is:", U.String())

	for len(key) > 0 {
		var kp kyber.Point
		kp = suite.Point().Embed(key, suite.RandomStream())
		log.Lvl4("Keypoint:", kp.String())
		log.Lvl4("X:", X.String())
		Cs = append(Cs, suite.Point().Add(C, kp))
		log.Lvl4("Cs:", C.String())
		key = key[min(len(key), kp.EmbedLen()):]
	}
	return
}

// DecodeKey can be used by the reader of ByzCoin to convert the
// re-encrypted secret back to a symmetric key that can be used later to decode
// the document.
//
// Input:
//   - suite - the cryptographic suite to use
//   - X - the aggregate public key of the DKG
//   - Cs - the encrypted key-slices
//   - XhatEnc - the re-encrypted schnorr-commit
//   - xc - the private key of the reader
//
// Output:
//   - key - the re-assembled key
//   - err - an eventual error when trying to recover the data from the points
func DecodeKey(suite kyber.Group, X kyber.Point, Cs []kyber.Point, XhatEnc kyber.Point,
	xc kyber.Scalar) (key []byte, err error) {
	log.LLvl4("xc:", xc)
	xcInv := suite.Scalar().Neg(xc)
	log.LLvl4("xcInv:", xcInv)
	sum := suite.Scalar().Add(xc, xcInv)
	log.LLvl4("xc + xcInv:", sum, "::", xc)
	log.LLvl4("X:", X)
	XhatDec := suite.Point().Mul(xcInv, X)
	log.LLvl4("XhatDec:", XhatDec)
	log.LLvl4("XhatEnc:", XhatEnc)
	Xhat := suite.Point().Add(XhatEnc, XhatDec)
	log.LLvl4("Xhat:", Xhat)
	XhatInv := suite.Point().Neg(Xhat)
	log.LLvl4("XhatInv:", XhatInv)

	// Decrypt Cs to keyPointHat
	for _, C := range Cs {
		log.LLvl4("C:", C)
		keyPointHat := suite.Point().Add(C, XhatInv)
		log.LLvl4("keyPointHat:", keyPointHat)
		keyPart, err := keyPointHat.Data()
		log.LLvl4("keyPart:", keyPart)
		if err != nil {
			return nil, err
		}
		key = append(key, keyPart...)
	}
	return
}
