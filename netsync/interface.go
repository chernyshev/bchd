// Copyright (c) 2017 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netsync

import (
	"github.com/chernyshev/bchd/blockchain"
	"github.com/chernyshev/bchd/chaincfg"
	"github.com/chernyshev/bchd/chaincfg/chainhash"
	"github.com/chernyshev/bchd/mempool"
	"github.com/chernyshev/bchd/peer"
	"github.com/chernyshev/bchd/wire"
	"github.com/gcash/bchutil"
)

// PeerNotifier exposes methods to notify peers of status changes to
// transactions, blocks, etc. Currently server (in the main package) implements
// this interface.
type PeerNotifier interface {
	AnnounceNewTransactions(newTxs []*mempool.TxDesc)

	UpdatePeerHeights(latestBlkHash *chainhash.Hash, latestHeight int32, updateSource *peer.Peer)

	RelayInventory(invVect *wire.InvVect, data interface{})

	TransactionConfirmed(tx *bchutil.Tx)
}

// Config is a configuration struct used to initialize a new SyncManager.
type Config struct {
	PeerNotifier PeerNotifier
	Chain        *blockchain.BlockChain
	TxMemPool    *mempool.TxPool
	ChainParams  *chaincfg.Params

	DisableCheckpoints bool
	MaxPeers           int

	FeeEstimator *mempool.FeeEstimator

	MinSyncPeerNetworkSpeed uint64

	FastSyncMode bool

	RegTestSyncAnyHost bool
}
