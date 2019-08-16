package config

import (
	"fmt"
	"math/big"

	"github.com/threefoldtech/rivine/build"
	"github.com/threefoldtech/rivine/modules"
	"github.com/threefoldtech/rivine/types"
)

var (
	rawVersion = "v0.1"
	// Version of the chain binaries.
	//
	// Value is defined by a private build flag,
	// or hardcoded to the latest released tag as fallback.
	Version build.ProtocolVersion
)

const (
	// TokenUnit defines the unit of one Token.
	TokenUnit = "TFB"
	// TokenChainName defines the name of the chain.
	TokenChainName = "bctempla"
)

// chain names
const (
	NetworkNameStandard = "standard"
	NetworkNameTest     = "testnet"
	NetworkNameDev      = "devnet"
)

// global network config constants
const (
	BlockFrequency types.BlockHeight = 120 // 1 block per 2 minutes on average
)

// GetBlockchainInfo returns the naming and versioning of tfchain.
func GetBlockchainInfo() types.BlockchainInfo {
	return types.BlockchainInfo{
		Name:            TokenChainName,
		NetworkName:     NetworkNameTest,
		CoinUnit:        TokenUnit,
		ChainVersion:    Version,       // use our own blockChain/build version
		ProtocolVersion: build.Version, // use latest available rivine protocol version
	}
}

// GetStandardnetGenesis explicitly sets all the required constants for the genesis block of the standard (prod) net
func GetStandardnetGenesis() types.ChainConstants {
	cfg := types.StandardnetChainConstants()
  
	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersionOne
	cfg.GenesisTransactionVersion = types.TransactionVersionOne

	// 2 minute block time
	cfg.BlockFrequency = 120

	// Payouts take roughly 1 day to mature.
	cfg.MaturityDelay = 144

  // The genesis timestamp
  cfg.GenesisTimestamp = types.Timestamp(1524168391)

	// 1000 block window for difficulty
	cfg.TargetWindow = 1000

	cfg.MaxAdjustmentUp = big.NewRat(10, 25)
	cfg.MaxAdjustmentDown = big.NewRat(10, 25)

	cfg.FutureThreshold = 3600        // 1 hour.
	cfg.ExtremeFutureThreshold = 7200 // 2 hours.

	cfg.StakeModifierDelay = 2000

	// Blockstakes can be used roughly 1 day after receiving
	cfg.BlockStakeAging = 86400

	// Receive 0 coins when you create a block
	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(1)

	// Use 0.001 coins as minimum transaction fee
	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(0.1 * 100)

  // Foundation receives all transactions fees in a single pool address,
  cfg.TransactionFeeCondition = types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("017267221ef1947bb18506e390f1f9446b995acfb6d08d8e39508bb974d9830b8cb8fdca788e34")))

	// no  initial coins, except  1 for initial transaction fee payments
	cfg.GenesisCoinDistribution = []types.CoinOutput{}

  // allocate block stakes
  cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
    {
      Value:     types.NewCurrency64(3000),
      Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("01b5e42056ef394f2ad9b511a61cec874d25bebe2095682dd37455cbafed4bec154e382a23f90e"))),
    },
    
  }

  return cfg
}

// GetTestnetGenesis explicitly sets all the required constants for the genesis block of the testnet
func GetTestnetGenesis() types.ChainConstants {
	cfg := types.TestnetChainConstants()
  
	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersionOne
	cfg.GenesisTransactionVersion = types.TransactionVersionOne

	cfg.BlockFrequency = 120

	cfg.MaturityDelay = 144

  // The genesis timestamp
  cfg.GenesisTimestamp = types.Timestamp(1524168391)

	cfg.TargetWindow = 1000

	cfg.MaxAdjustmentUp = big.NewRat(10, 25)
	cfg.MaxAdjustmentDown = big.NewRat(10, 25)

	cfg.FutureThreshold = 3600
	cfg.ExtremeFutureThreshold = 7200

	cfg.StakeModifierDelay = 2000

	cfg.BlockStakeAging = 86400

	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(1.0)

	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(0.1 * 100)


	// allocate block stakes
  cfg.GenesisCoinDistribution = []types.CoinOutput{
    {
        Value: cfg.CurrencyUnits.OneCoin.Mul64(5000000000000000),
        Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("01434535fd01243c02c277cd58d71423163767a575a8ae44e15807bf545e4a8456a5c4afabad51"))),
      },
    {
        Value: cfg.CurrencyUnits.OneCoin.Mul64(5000000000000000),
        Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("03ca853bbabb8684bfa6ece69fb469b3a686c1bfffe774ab2d0154c1c80128e3398a070296f512"))),
      },
    
  }

  // allocate block stakes
  cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
    {
      Value:     types.NewCurrency64(3000),
      Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("01b5e42056ef394f2ad9b511a61cec874d25bebe2095682dd37455cbafed4bec154e382a23f90e"))),
    },
    
  }

  return cfg
}

// GetDevnetGenesis explicitly sets all the required constants for the genesis block of the devnet
func GetDevnetGenesis() types.ChainConstants {
	cfg := types.DevnetChainConstants()
  
	// set transaction versions
	cfg.DefaultTransactionVersion = types.TransactionVersionOne
	cfg.GenesisTransactionVersion = types.TransactionVersionOne

	cfg.BlockFrequency = 12

	cfg.MaturityDelay = 10

  // The genesis timestamp
  cfg.GenesisTimestamp = types.Timestamp(1524168391)

	cfg.TargetWindow = 20

	cfg.MaxAdjustmentUp = big.NewRat(10, 25)
	cfg.MaxAdjustmentDown = big.NewRat(10, 25)

	cfg.FutureThreshold = 120
	cfg.ExtremeFutureThreshold = 240

	cfg.StakeModifierDelay = 2000

	cfg.BlockStakeAging = 1024

	cfg.BlockCreatorFee = cfg.CurrencyUnits.OneCoin.Mul64(1.0)

	cfg.MinimumTransactionFee = cfg.CurrencyUnits.OneCoin.Div64(0.1 * 100)


	// allocate block stakes
  cfg.GenesisCoinDistribution = []types.CoinOutput{
    {
        Value: cfg.CurrencyUnits.OneCoin.Mul64(5000000000000000),
        Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
      },
    
  }

  // allocate block stakes
  cfg.GenesisBlockStakeAllocation = []types.BlockStakeOutput{
    {
      Value:     types.NewCurrency64(3000),
      Condition: types.NewCondition(types.NewUnlockHashCondition(unlockHashFromHex("015a080a9259b9d4aaa550e2156f49b1a79a64c7ea463d810d4493e8242e6791584fbdac553e6f"))),
    },
    
  }

  return cfg
}

// GetStandardnetBootstrapPeers sets the standard bootstrap node addresses
func GetStandardnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
    "bootstrap1.threefoldtoken.com:23112",
    "bootstrap2.threefoldtoken.com:23112",
    "bootstrap3.threefoldtoken.com:23112",
    "bootstrap4.threefoldtoken.com:23112",
    "bootstrap5.threefoldtoken.com:23112",
    
	}
}

// GetTestnetBootstrapPeers sets the testnet bootstrap node addresses
func GetTestnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
    "bootstrap1.testnet.threefoldtoken.com:23112",
    "bootstrap2.testnet.threefoldtoken.com:23112",
    "bootstrap3.testnet.threefoldtoken.com:23112",
    "bootstrap4.testnet.threefoldtoken.com:24112",
    "bootstrap5.testnet.threefoldtoken.com:24112",
    
	}
}

// GetDevnetBootstrapPeers sets the default devnet bootstrap node addresses
func GetDevnetBootstrapPeers() []modules.NetAddress {
	return []modules.NetAddress{
		"localhost:23112", // TODO: add port
	}
}

func unlockHashFromHex(hstr string) (uh types.UnlockHash) {
	err := uh.LoadString(hstr)
	if err != nil {
		panic(fmt.Sprintf("func unlockHashFromHex(%s) failed: %v", hstr, err))
	}
	return
}

func init() {
	Version = build.MustParse(rawVersion)
}
