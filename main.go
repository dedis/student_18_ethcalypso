package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/dedis/cothority"
	"github.com/dedis/cothority/byzcoin/bcadmin/lib"
	"github.com/dedis/cothority/darc"
	"github.com/dedis/onet/log"
	"github.com/dedis/onet/network"
	"github.com/dedis/student_18_ethcalypso/calypso"
	"github.com/ethereum/go-ethereum/common"
	cli "gopkg.in/urfave/cli.v1"
)

var cliApp = cli.NewApp()

var cmds = cli.Commands{
	{
		Name:    "LTSID",
		Usage:   "create a ledger",
		Aliases: []string{"c"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
			cli.DurationFlag{
				Name:  "interval, i",
				Usage: "the block interval for this ledger",
				Value: 5 * time.Second,
			},
		},
		Action: LTSID,
	},
	{
		Name:    "AddWrite",
		Usage:   "Add write request and returns the address",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "LTSID",
				Usage: "the ByzCoin config to use",
			},
			cli.StringFlag{
				Name:  "X",
				Usage: "The X variable of a writeRequest",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
		},
		Action: AddWrite,
	},
	{
		Name:    "AddRead",
		Usage:   "Add Read request and returns the address and the secret",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "wr",
				Usage: "the ByzCoin config to use",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
		},
		Action: AddRead,
	},
	{
		Name:    "DecryptKey",
		Usage:   "Decrypts the key",
		Aliases: []string{"s"},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "wr",
				Usage: "The address of the write request",
			},
			cli.StringFlag{
				Name:  "rr",
				Usage: "The address of the read request",
			},
			cli.StringFlag{
				Name:  "roster, r",
				Usage: "the roster of the cothority that will host the ledger",
			},
		},
		Action: DecryptKey,
	},
}

func init() {
	network.RegisterMessages(&darc.Darc{}, &darc.Identity{}, &darc.Signer{})
}

func show(c *cli.Context) {
	fmt.Println(c.String("bc"))
}

func LTSID(c *cli.Context) error {
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		return err
	}
	roster := r.NewRosterWithRoot(r.List[0])
	fmt.Println("Roster ", roster)
	calypsoClient := calypso.NewClient(*roster)
	LTS, e := calypsoClient.CreateLTS()
	if e != nil {
		fmt.Println("Can't get LTS")
		return e
	}
	hexLTSID := hex.EncodeToString(LTS.LTSID)
	fmt.Println("LTSID: ", hexLTSID)
	fmt.Println("X: ", LTS.X.String())
	return nil

}

func init() {
	cliApp.Name = "bcadmin"
	cliApp.Usage = "Create eth whatever"
	cliApp.Version = "0.1"
	cliApp.Commands = cmds
	cliApp.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "debug, d",
			Value: 0,
			Usage: "debug-level: 1 for terse, 5 for maximal",
		},
	}
	cliApp.Before = func(c *cli.Context) error {
		log.SetDebugVisible(c.Int("debug"))
		lib.ConfigPath = c.String("config")
		return nil
	}
}

func DecryptKey(c *cli.Context) {
	wr := c.String("wr")
	rr := c.String("rr")
	wrAddr := common.HexToAddress(wr)
	rrAddr := common.HexToAddress(rr)
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		log.Fatal(err)
	}
	roster := r.NewRosterWithRoot(r.List[0])
	calypsoClient := calypso.NewClient(*roster)
	dk := &calypso.DecryptKey{
		Write: wrAddr,
		Read:  rrAddr,
	}
	dkr, e := calypsoClient.DecryptKey(dk)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println("Xhat: ", dkr.XhatEnc.String())
	fmt.Println("Cs", dkr.Cs)
	fmt.Println("X is ", dkr.X.String())
}

func AddRead(c *cli.Context) error {
	wr := c.String("wr")
	wrAddr := common.HexToAddress(wr)
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		return err
	}
	roster := r.NewRosterWithRoot(r.List[0])
	fmt.Println("Roster ", roster)
	calypsoClient := calypso.NewClient(*roster)
	secret, rAddr, e := calypsoClient.AddRead(wrAddr)
	if e != nil {
		return e
	}
	fmt.Println("Your secret is: ", secret)
	fmt.Println("Read address is ", rAddr.Hex())
	return nil
}

func AddWrite(c *cli.Context) error {
	ltsid := c.String("LTSID")
	X := c.String("X")
	xHex, e := hex.DecodeString(X)
	if e != nil {
		log.Fatal(e)
	}
	point := cothority.Suite.Point()
	e = point.UnmarshalBinary(xHex)
	if e != nil {
		return e
	}
	fmt.Println("LTSID: ", ltsid)
	fmt.Println("Value of X: ", X)
	fn := c.String("roster")
	r, err := lib.ReadRoster(fn)
	if err != nil {
		return err
	}
	roster := r.NewRosterWithRoot(r.List[0])
	fmt.Println("Roster ", roster)
	calypsoClient := calypso.NewClient(*roster)
	id, e := hex.DecodeString(ltsid)
	if e != nil {
		return e
	}
	wrAddr, e := calypsoClient.AddWrite(id, point, []byte("Sabrina"))
	fmt.Println("Address of write is: ", wrAddr.Hex())
	return e
}

func main() {
	log.ErrFatal(cliApp.Run(os.Args))
}