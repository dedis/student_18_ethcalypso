package gocontracts

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

/*func TestAddAWriteRequest(t *testing.T) {
	client, e := ethereum.GetClient()
	require.Nil(t, e)
	privateKey, e := ethereum.GetPrivateKey()
	require.Nil(t, e)
	data := []byte("Bjorn is a cool dude")
	ed := []byte("Bjorn is an extra cool dude")
	ltsid := []byte("Sabrina")
	point := cothority.Suite.Point()
	u, e := point.MarshalBinary()
	require.Nil(t, e)
	cs := make([][]byte, 0)
	cs = append(cs, u)
	policy := ethereum.GetTestListOfAddresses()
	addr, _, _, e := ServiceDeployWriteRequest(privateKey, client, data, ed, ltsid, policy, u, cs)
	require.Nil(t, e)
	fmt.Println("Address of write ", addr)
	//Now we check that it got deployed
	wr, e := ServiceGetWriteRequest(privateKey, client, addr)
	require.Nil(t, e)
	//Check that we still have the same values
	require.Equal(t, "Sabrina", string(wr.LTSID))
	//Check that we still have the same points
	require.True(t, wr.U.Equal(point))
}*/

func TestEmptyCalypso(t *testing.T) {
	a, e := GetStaticCalypso()
	fmt.Println("Address of claypso", (*a).Hex())
	require.Nil(t, e)
}

/*func TestAddReadRequestToCalypso(t *testing.T) {
	client, e := ethereum.GetClient()
	require.Nil(t, e)
	privateKey, e := ethereum.GetPrivateKey()
	require.Nil(t, e)
	calypsoAddr, e := GetStaticCalypso()
	require.Nil(t, e)
	policy := ethereum.GetTestListOfAddresses()
	d, ed, ltsid, u, cs := GetMockWriteRequest(t)
	wr, _, _, e := ServiceDeployWriteRequest(privateKey, client, d, ed, ltsid, policy, u, cs)
	require.Nil(t, e)
	point := cothority.Suite.Point()
	Xc, e := point.MarshalBinary()
	require.Nil(t, e)
	rr, _, _, e := ServiceDeployReadRequest(privateKey, client, wr, Xc)
	require.Nil(t, e)
	_, e = ServiceAddWriteRequest(privateKey, *calypsoAddr, wr, client)
	require.Nil(t, e)
	_, e = ServiceAddReadRequest(privateKey, *calypsoAddr, rr, client)
	require.Nil(t, e)
	//Add fake read request
	fakeAddr, _, _, e := ServiceDeployReadRequest(privateKey, client, policy[0], Xc)
	require.Nil(t, e)
	_, e = ServiceAddReadRequest(privateKey, *calypsoAddr, fakeAddr, client)
	require.NotNil(t, e)
	//Now to make try and make a transaction to the blockchain while not being a registered owner
	fakePrivateKey, e := ethereum.GetAnotherPrivateKey()
	//Create a fake WriteRequest to try and pass to the Calypso contract
	fakeWR, _, _, e := ServiceDeployWriteRequest(fakePrivateKey, client, d, ed, ltsid, policy, u, cs)
	require.Nil(t, e)
	_, e = ServiceAddWriteRequest(fakePrivateKey, *calypsoAddr, fakeWR, client)
	require.NotNil(t, e)
	canRead := ServiceCheckIfCanRead(privateKey, *calypsoAddr, rr, client)
	require.True(t, canRead)
}*/
