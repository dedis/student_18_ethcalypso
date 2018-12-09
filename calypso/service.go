// Package calypso implements the LTS functionality of the Calypso paper. It
// implements both the access-control cothority and the secret management
// cothority. (1) The access-control cothority is implemented using ByzCoin
// with two contracts, `Write` and `Read` (2) The secret-management cothority
// uses an onet service with methods to set up a Long Term Secret (LTS)
// distributed key and to request a re-encryption
//
// For more details, see
// https://github.com/dedis/cothority/tree/master/calypso/README.md
package calypso

import (
	"errors"
	"fmt"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/darc"
	dkgprotocol "github.com/dedis/cothority/dkg"
	"github.com/dedis/kyber"
	"github.com/dedis/kyber/share"
	"github.com/dedis/kyber/util/random"
	"github.com/dedis/onet"
	"github.com/dedis/onet/log"
	"github.com/dedis/onet/network"
	"github.com/dedis/protobuf"
	"github.com/dedis/student_18_ethcalypso/calypso/ethac/gocontracts"
	"github.com/dedis/student_18_ethcalypso/calypso/ethereum"
	"github.com/dedis/student_18_ethcalypso/calypso/protocol"
	"github.com/ethereum/go-ethereum/common"
)

// Used for tests
var calypsoID onet.ServiceID

// ServiceName of the secret-management part of Calypso.
var ServiceName = "Calypso"

// dkgTimeout is how long the system waits for the DKG to finish
const propagationTimeout = 10 * time.Second

func init() {
	var err error
	calypsoID, err = onet.RegisterNewService(ServiceName, newService)
	if err != nil {
		fmt.Println("Error in service student")
	}
	log.ErrFatal(err)
	network.RegisterMessages(&storage1{}, &vData{})
}

// Service is our calypso-service. It stores all created LTSs.
type Service struct {
	*onet.ServiceProcessor
	storage *storage1
}

// pubPoly is a serializable version of share.PubPoly
type pubPoly struct {
	B       kyber.Point
	Commits []kyber.Point
}

// vData is sent to all nodes when re-encryption takes place. If Ephemeral
// is non-nil, Signature needs to hold a valid signature from the reader
// in the Proof.
type vData struct {
	ETHReadAddress common.Address
	Ephemeral      kyber.Point
	Signature      *darc.Signature
}

//LogAddress takes in an address of a deployed contract
//and stores it in the storage of the LTS
func (s *Service) LogAddress(la *LogAddress) (lr *LogAddressReply, err error) {
	fmt.Println("Bjorn this function worked")
	lr = &LogAddressReply{Added: true}
	return lr, nil
}

// CreateLTS takes as input a roster with a list of all nodes that should
// participate in the DKG. Every node will store its private key and wait
// for decryption requests.
// This method will create a random LTSID that can be used to reference
// the LTS group created.
func (s *Service) CreateLTS(cl *CreateLTS) (reply *CreateLTSReply, err error) {
	tree := cl.Roster.GenerateNaryTreeWithRoot(len(cl.Roster.List), s.ServerIdentity())
	pi, err := s.CreateProtocol(dkgprotocol.Name, tree)
	setupDKG := pi.(*dkgprotocol.Setup)
	setupDKG.Wait = true
	reply = &CreateLTSReply{LTSID: make([]byte, 32)}
	random.New().XORKeyStream(reply.LTSID, reply.LTSID)
	setupDKG.SetConfig(&onet.GenericConfig{Data: reply.LTSID})
	log.Lvlf3("%s: reply.LTSID is: %x", s.ServerIdentity(), reply.LTSID)
	if err := pi.Start(); err != nil {
		return nil, err
	}
	log.Lvl3("Started DKG-protocol - waiting for done", len(cl.Roster.List))
	select {
	case <-setupDKG.Finished:
		shared, err := setupDKG.SharedSecret()
		if err != nil {
			return nil, err
		}
		s.storage.Lock()
		s.storage.Shared[string(reply.LTSID)] = shared
		dks, err := setupDKG.DKG.DistKeyShare()
		if err != nil {
			s.storage.Unlock()
			return nil, err
		}
		s.storage.Polys[string(reply.LTSID)] = &pubPoly{s.Suite().Point().Base(), dks.Commits}
		s.storage.Rosters[string(reply.LTSID)] = &cl.Roster
		s.storage.ContractAddresses[string(reply.LTSID)] = cl.CalypsoAddress
		s.storage.Unlock()
		reply.X = shared.X
	case <-time.After(propagationTimeout):
		return nil, errors.New("dkg didn't finish in time")
	}
	return
}

func compare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if b[i] != val {
			return false
		}
	}
	return true
}

// DecryptKey takes as an input a Read- and a Write-proof. Proofs contain
// everything necessary to verify that a given instance is correct and
// stored in ByzCoin.
// Using the Read and the Write-instance, this method verifies that the
// requests match and then re-encrypts the secret to the public key given
// in the Read-instance.
// TODO: support ephemeral keys.
func (s *Service) DecryptKey(dkr *DecryptKey) (reply *DecryptKeyReply, err error) {
	privateKey, e := ethereum.GetPrivateKey()
	client, e := ethereum.GetClient()
	if e != nil {
		return nil, e
	}
	reply = &DecryptKeyReply{}
	wAddress := dkr.Write
	rAddress := dkr.Read
	readHash := dkr.ReadHash
	isReadMined, e := gocontracts.IsMined(client, readHash)
	if !isReadMined || e != nil {
		return nil, errors.New("Read was not mined")
	}
	write, e := gocontracts.ServiceGetWriteRequest(privateKey, client, wAddress)
	if e != nil {
		return nil, e
	}
	read, e := gocontracts.ServiceGetRead(rAddress, client, privateKey)
	if e != nil {
		return nil, e
	}
	writeHash := common.BytesToHash(read.WriteHash)
	isWriteMined, e := gocontracts.IsMined(client, writeHash)
	if !isWriteMined || e != nil {
		return nil, errors.New("The write was not mined")
	}
	if !compare(read.Write.Bytes(), wAddress.Bytes()) {
		return nil, errors.New("This read did not point to this write")
	}
	cal := s.storage.ContractAddresses[string(write.LTSID)]
	canRead := gocontracts.ServiceCheckIfCanRead(privateKey, cal, rAddress, client)
	if !canRead {
		return nil, errors.New("This is not a valid read/write pair")
	}
	//I'll have to rewrite until here
	s.storage.Lock()
	roster := s.storage.Rosters[string(write.LTSID)]
	if roster == nil {
		s.storage.Unlock()
		return nil, errors.New("don't know the LTSID stored in write")
	}

	//I don't think I'll use this
	scID := make([]byte, 32)
	copy(scID, s.storage.OLIDs[string(write.LTSID)])
	s.storage.Unlock()

	// Start ocs-protocol to re-encrypt the file's symmetric key under the
	// reader's public key.
	nodes := len(roster.List)
	threshold := nodes - (nodes-1)/3
	tree := roster.GenerateNaryTreeWithRoot(nodes, s.ServerIdentity())
	pi, err := s.CreateProtocol(protocol.NameOCS, tree)
	if err != nil {
		return nil, err
	}
	ocsProto := pi.(*protocol.OCS)
	//I'll need U. I have U.
	ocsProto.U = write.U

	verificationData := &vData{
		ETHReadAddress: dkr.Read,
	}

	ocsProto.Xc = read.Xc
	log.Lvlf2("Public key is: %s", ocsProto.Xc)
	ocsProto.VerificationData, err = protobuf.Encode(verificationData)
	if err != nil {
		return nil, errors.New("couldn't marshal verification data: " + err.Error())
	}

	// Make sure everything used from the s.Storage structure is copied, so
	// there will be no races.
	s.storage.Lock()
	ocsProto.Shared = s.storage.Shared[string(write.LTSID)]
	pp := s.storage.Polys[string(write.LTSID)]
	reply.X = s.storage.Shared[string(write.LTSID)].X.Clone()
	var commits []kyber.Point
	for _, c := range pp.Commits {
		commits = append(commits, c.Clone())
	}
	ocsProto.Poly = share.NewPubPoly(s.Suite(), pp.B.Clone(), commits)
	s.storage.Unlock()

	log.Lvl3("Starting reencryption protocol")
	ocsProto.SetConfig(&onet.GenericConfig{Data: write.LTSID})
	err = ocsProto.Start()
	if err != nil {
		return nil, err
	}
	if !<-ocsProto.Reencrypted {
		return nil, errors.New("reencryption got refused")
	}
	log.Lvl3("Reencryption protocol is done.")
	reply.XhatEnc, _ = share.RecoverCommit(cothority.Suite, ocsProto.Uis, threshold, nodes)
	if err != nil {
		return nil, err
	}
	reply.Cs = write.Cs
	log.Lvl3("Successfully reencrypted the key")
	return
}

// SharedPublic returns the shared public key of an LTSID group.
func (s *Service) SharedPublic(req *SharedPublic) (reply *SharedPublicReply, err error) {
	log.Lvl2("Getting shared public key")
	s.storage.Lock()
	shared, ok := s.storage.Shared[string(req.LTSID)]
	s.storage.Unlock()
	if !ok {
		return nil, errors.New("didn't find this Long Term Secret")
	}
	return &SharedPublicReply{X: shared.X}, nil
}

// NewProtocol intercepts the DKG and OCS protocols to retrieve the values
func (s *Service) NewProtocol(tn *onet.TreeNodeInstance, conf *onet.GenericConfig) (onet.ProtocolInstance, error) {
	log.Lvl3(s.ServerIdentity(), tn.ProtocolName(), conf)
	switch tn.ProtocolName() {
	case dkgprotocol.Name:
		pi, err := dkgprotocol.NewSetup(tn)
		if err != nil {
			return nil, err
		}
		setupDKG := pi.(*dkgprotocol.Setup)
		go func(conf *onet.GenericConfig) {
			<-setupDKG.Finished
			shared, err := setupDKG.SharedSecret()
			if err != nil {
				log.Error(err)
				return
			}
			log.Lvl3(s.ServerIdentity(), "Got shared", shared)
			s.storage.Lock()
			s.storage.Shared[string(conf.Data)] = shared
			s.storage.Unlock()
			s.save()
		}(conf)
		return pi, nil
	case protocol.NameOCS:
		s.storage.Lock()
		shared, ok := s.storage.Shared[string(conf.Data)]
		s.storage.Unlock()
		if !ok {
			return nil, errors.New("didn't find skipchain")
		}
		pi, err := protocol.NewOCS(tn)
		if err != nil {
			return nil, err
		}
		ocs := pi.(*protocol.OCS)
		ocs.Shared = shared
		ocs.Verify = s.verifyReencryption
		return ocs, nil
	}
	return nil, nil
}

// verifyReencryption checks that the read and the write instances match.
func (s *Service) verifyReencryption(rc *protocol.Reencrypt) bool {
	err := func() error {
		privateKey, e := ethereum.GetPrivateKey()
		if e != nil {
			return e
		}

		client, e := ethereum.GetClient()
		if e != nil {
			return nil
		}
		var verificationData vData
		e = protobuf.DecodeWithConstructors(*rc.VerificationData, &verificationData, network.DefaultConstructors(cothority.Suite))
		if e != nil {
			return e
		}
		rr := verificationData.ETHReadAddress
		r, e := gocontracts.ServiceGetRead(rr, client, privateKey)
		if e != nil {
			return e
		}

		if verificationData.Ephemeral != nil {
			return errors.New("ephemeral keys not supported yet")
		}
		if !r.Xc.Equal(rc.Xc) {
			return errors.New("wrong reader")
		}
		return nil
	}()
	if err != nil {
		log.Lvl2(s.ServerIdentity(), "wrong reencryption:", err)
		return false
	}
	return true
}

// newService receives the context that holds information about the node it's
// running on. Saving and loading can be done using the context. The data will
// be stored in memory for tests and simulations, and on disk for real deployments.
func newService(c *onet.Context) (onet.Service, error) {
	s := &Service{
		ServiceProcessor: onet.NewServiceProcessor(c),
	}
	if err := s.RegisterHandlers(s.CreateLTS, s.DecryptKey, s.SharedPublic, s.LogAddress); err != nil {
		return nil, errors.New("couldn't register messages")
	}
	//byzcoin.RegisterContract(c, ContractWriteID, s.ContractWrite)
	//byzcoin.RegisterContract(c, ContractReadID, s.ContractRead)
	if err := s.tryLoad(); err != nil {
		log.Error(err)
		return nil, err
	}
	s.storage.ContractAddresses = make(map[string]common.Address)
	s.storage.ContractTx = make(map[string][]byte)
	return s, nil
}
