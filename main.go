package main

import (
	"context"
	"fmt"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-rendezvous"
	rdbsqlite "github.com/libp2p/go-libp2p-rendezvous/db/sqlite"
)

func main() {
	opt := libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/7654")

	h, err := libp2p.New(context.Background(), opt)
	if err != nil {
		panic(err)
	}

	fmt.Println("My addresses")
	for _, a := range h.Addrs() {
		fmt.Printf("%s/ipfs/%s\n", a, h.ID().Pretty())
	}

	db, err := rdbsqlite.OpenDB(context.Background(), "rzsqldq")
	if err != nil {
		panic(err)
	}

	_ = rendezvous.NewRendezvousService(h, db)

	select {}
}
