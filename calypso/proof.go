package calypso

import (
	"crypto/sha256"
	"errors"

	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
)

// CheckProof verifies that the write-request has actually been created with
// somebody having access to the secret key.
func CheckProof(suite suite, wr *gocontracts.Write) error {
	gf := suite.Point().Mul(wr.F, nil)
	ue := suite.Point().Mul(suite.Scalar().Neg(wr.E), wr.U)
	w := suite.Point().Add(gf, ue)

	gBar := suite.Point().Mul(suite.Scalar().SetBytes(wr.LTSID), nil)
	gfBar := suite.Point().Mul(wr.F, gBar)
	ueBar := suite.Point().Mul(suite.Scalar().Neg(wr.E), wr.Ubar)
	wBar := suite.Point().Add(gfBar, ueBar)

	hash := sha256.New()
	for _, c := range wr.Cs {
		c.MarshalTo(hash)
	}
	wr.U.MarshalTo(hash)
	wr.Ubar.MarshalTo(hash)
	w.MarshalTo(hash)
	wBar.MarshalTo(hash)
	//hash.Write(writeID)

	e := suite.Scalar().SetBytes(hash.Sum(nil))
	if e.Equal(wr.E) {
		return nil
	}

	return errors.New("recreated proof is not equal to stored proof")
}
