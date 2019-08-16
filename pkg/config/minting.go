package config

import (
	"github.com/threefoldtech/rivine/types"
)

// GetStandardGenesisMintCondition returns the standard network  minting condition
func GetStandardGenesisMintCondition() types.UnlockConditionProxy {
	// TODO: define final multisig condition
  address := ""
	address = "01434535fd01243c02c277cd58d71423163767a575a8ae44e15807bf545e4a8456a5c4afabad51"
  
	var uh types.UnlockHash
	if err := uh.LoadString(address); err != nil {
		panic(err)
	}
	condition := types.NewCondition(types.NewUnlockHashCondition(uh))
	return condition
}

// GetTestnetGenesisMintCondition returns the testnet network  minting condition
func GetTestnetGenesisMintCondition() types.UnlockConditionProxy {
	// @leesmet
  address := ""
  address = "01434535fd01243c02c277cd58d71423163767a575a8ae44e15807bf545e4a8456a5c4afabad51"
  
	var uh types.UnlockHash
	if err := uh.LoadString(address); err != nil {
		panic(err)
	}
	condition := types.NewCondition(types.NewUnlockHashCondition(uh))
	return condition
}

// GetDevnetGenesisMintCondition returns the devnet network  minting condition
func GetDevnetGenesisMintCondition() types.UnlockConditionProxy {
	// belongs to wallet with mnemonic:
	// carbon boss inject cover mountain fetch fiber fit tornado cloth wing dinosaur proof joy intact fabric thumb rebel borrow poet chair network expire else
  address := ""
  address = "01434535fd01243c02c277cd58d71423163767a575a8ae44e15807bf545e4a8456a5c4afabad51"
	var uh types.UnlockHash
	if err := uh.LoadString(address); err != nil {
		panic(err)
	}
	condition := types.NewCondition(types.NewUnlockHashCondition(uh))
	return condition
}
