package rtns_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	ci "gx/ipfs/QmNiJiXwWE3kRhZrC5ej3kSjWHm337pYfhjLGSCDNKJP2s/go-libp2p-crypto"
	peer "gx/ipfs/QmY5Grm8pJdiSSVsYxx4uNRgweY72EmYwuSDbRnbFok3iY/go-libp2p-peer"

	"github.com/RTradeLtd/Temporal/rtns"
)

type contextKey string

const (
	ipnsPublishTTL contextKey = "ipns-publish-ttl"
	testPath                  = "/ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv"
	testSwarmADDR             = "/ip4/0.0.0.0/tcp/4002"
)

func TestPublisher_NoGen(t *testing.T) {
	// create our private key
	pk, _, err := ci.GenerateKeyPair(ci.Ed25519, 256)
	if err != nil {
		t.Fatal(err)
	}
	publisher, err := rtns.NewPublisher(pk, false, testSwarmADDR)
	if err != nil {
		t.Fatal(err)
	}
	// sleep giving time for our node to discover some peers
	time.Sleep(time.Second * 15)
	// create our private key
	pk, _, err = ci.GenerateKeyPair(ci.Ed25519, 256)
	if err != nil {
		t.Fatal(err)
	}
	if pid, err := peer.IDFromPrivateKey(pk); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println("id to check ", pid.Pretty())
	}
	if err := publisher.Publish(context.Background(), pk, testPath); err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), ipnsPublishTTL, time.Minute*10)
	if err := publisher.Publish(ctx, pk, testPath); err != nil {
		t.Fatal(err)
	}
}
