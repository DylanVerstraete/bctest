package main

import (
	"github.com/threefoldtech/rivine/extensions/minting"
  mintingcli "github.com/threefoldtech/rivine/extensions/minting/client"

	"github.com/threefoldtech/rivine/pkg/client"
	"github.com/threefoldtech/rivine/types"

  bctemplatypes "github.com/DylanVerstraete/bctempla/pkg/types"
)

// RegisterStandardTransactions registers the bctempla-specific transactions as required for the standard network.
func RegisterStandardTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterTestnetTransactions registers the bctempla-specific transactions as required for the test network.
func RegisterTestnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

// RegisterDevnetTransactions registers the bctempla-specific transactions as required for the dev network.
func RegisterDevnetTransactions(cli *client.CommandLineClient) {
	registerTransactions(cli)
}

func registerTransactions(cli *client.CommandLineClient) {
  // create minting plugin client...
  mintingCLI := mintingcli.NewPluginConsensusClient(cli)
  // ...and register minting types
  types.RegisterTransactionVersion(bctemplatypes.MinterDefinitionTxVersion, minting.MinterDefinitionTransactionController{
    MintConditionGetter: mintingCLI,
    TransactionVersion:  bctemplatypes.MinterDefinitionTxVersion,
  })
  types.RegisterTransactionVersion(bctemplatypes.CoinCreationTxVersion, minting.CoinCreationTransactionController{
    MintConditionGetter: mintingCLI,
    TransactionVersion:  bctemplatypes.CoinCreationTxVersion,
  })
	types.RegisterTransactionVersion(bctemplatypes.CoinDestructionTxVersion, minting.CoinDestructionTransactionController{
		TransactionVersion: bctemplatypes.CoinDestructionTxVersion,
	})
}
