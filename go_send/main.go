package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/config"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

func main() {
	command := os.Args[1]
	from := os.Args[2]
	to := os.Args[3]
	amount := os.Args[4]
	if command != "send" {
		return
	}

	embending_account := []string{"Alice", "Bob", "Dave", "Eve", "Charlie"}

	var from_address, to_address signature.KeyringPair
	var err error
	for _, v := range embending_account {
		if strings.EqualFold(v, from) {
			from_address, err = signature.KeyringPairFromSecret("//"+v, 42)
			if err != nil {
				log.Fatal(err)
			}
		} else if strings.EqualFold(v, to) {
			to_address, err = signature.KeyringPairFromSecret("//"+v, 42)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println(from_address, to_address)

	api, err := gsrpc.NewSubstrateAPI(config.Default().RPCURL)
	if err != nil {
		log.Fatal(err)
	}
	meta, err := api.RPC.State.GetMetadataLatest()
	if err != nil {
		log.Fatal(err)
	}

	bal, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		log.Fatal(fmt.Errorf("failed to convert balance"))
	}

	to_multi_address, err := types.NewMultiAddressFromAccountID(to_address.PublicKey)
	if err != nil {
		panic(err)
	}
	c, err := types.NewCall(meta, "Balances.transfer", to_multi_address, types.NewUCompact(bal))
	if err != nil {
		log.Fatal(err)
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(c)

	genesisHash, err := api.RPC.Chain.GetBlockHash(0)
	if err != nil {
		log.Fatal(err)
	}

	rv, err := api.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		log.Fatal(err)
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", from_address.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	var accountInfo types.AccountInfo
	ok, err = api.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		log.Fatal(err)
	}

	nonce := uint32(accountInfo.Nonce)
	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(100),
		TransactionVersion: rv.TransactionVersion,
	}

	// Sign the transaction using Alice's default account
	err = ext.Sign(from_address, o)
	if err != nil {
		log.Fatal(err)
	}

	// Send the extrinsic
	hash, err := api.RPC.Author.SubmitExtrinsic(ext)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Balance transferred hash ", hash.Hex())

}
