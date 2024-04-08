package types

import (
	"cosmossdk.io/math"
	tmprotocrypto "github.com/cometbft/cometbft/proto/tendermint/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	epochsTypes "github.com/evmos/evmos/v14/x/epochs/types"
)

// EpochsKeeper represents the expected keeper interface for the epochs module.
type EpochsKeeper interface {
	GetEpochInfo(sdk.Context, string) (epochsTypes.EpochInfo, bool)
}

// DogfoodHooks represents the event hooks for dogfood module. Ideally, these should
// match those of the staking module but for now it is only a subset of them. The side effects
// of calling the other hooks are not relevant to running the chain, so they can be skipped.
type DogfoodHooks interface {
	AfterValidatorBonded(
		sdk.Context, sdk.ConsAddress, sdk.ValAddress,
	) error
}

// OperatorHooks is the interface for the operator module's hooks. The functions are called
// whenever an operator opts in to a Cosmos chain, opts out of a Cosmos chain, or replaces their
// public key with another one.
type OperatorHooks interface {
	AfterOperatorOptIn(sdk.Context, sdk.AccAddress, string, tmprotocrypto.PublicKey)
	AfterOperatorKeyReplacement(
		sdk.Context, sdk.AccAddress, tmprotocrypto.PublicKey, tmprotocrypto.PublicKey, string,
	)
	AfterOperatorOptOutInitiated(sdk.Context, sdk.AccAddress, string, tmprotocrypto.PublicKey)
}

// DelegationHooks represent the event hooks for delegation module.
type DelegationHooks interface {
	AfterDelegation(sdk.Context, sdk.AccAddress)
	AfterUndelegationStarted(sdk.Context, sdk.AccAddress, []byte) error
	AfterUndelegationCompleted(sdk.Context, sdk.AccAddress, []byte)
}

// OperatorKeeper represents the expected keeper interface for the operator module.
type OperatorKeeper interface {
	GetOperatorConsKeyForChainID(
		sdk.Context, sdk.AccAddress, string,
	) (bool, tmprotocrypto.PublicKey, error)
	IsOperatorOptingOutFromChainID(
		sdk.Context, sdk.AccAddress, string,
	) bool
	CompleteOperatorOptOutFromChainID(sdk.Context, sdk.AccAddress, string)
	DeleteOperatorAddressForChainIDAndConsAddr(sdk.Context, string, sdk.ConsAddress)
	GetOperatorAddressForChainIDAndConsAddr(
		sdk.Context, string, sdk.ConsAddress,
	) (bool, sdk.AccAddress)
	IsOperatorJailedForChainID(sdk.Context, sdk.AccAddress, string) bool
	Jail(sdk.Context, sdk.ConsAddress, string)
	// GetActiveOperatorsForChainID should return a list of operators and their public keys.
	// These operators should not be in the process of opting our, and should not be jailed
	// whether permanently or temporarily.
	GetActiveOperatorsForChainID(
		sdk.Context, string,
	) ([]sdk.AccAddress, []tmprotocrypto.PublicKey)
}

// DelegationKeeper represents the expected keeper interface for the delegation module.
type DelegationKeeper interface {
	IncrementUndelegationHoldCount(sdk.Context, []byte) error
	DecrementUndelegationHoldCount(sdk.Context, []byte) error
}

// EpochsHooks represents the event hooks for the epochs module.
type EpochsHooks interface {
	AfterEpochEnd(sdk.Context, string, int64)
	BeforeEpochStart(sdk.Context, string, int64)
}

// AssetsKeeper represents the expected keeper interface for the assets module.
type AssetsKeeper interface {
	GetOperatorAssetValue(sdk.Context, sdk.AccAddress) (int64, error)
	IsStakingAsset(sdk.Context, string) bool
	GetAvgDelegatedValue(
		sdk.Context, []sdk.AccAddress, []string, string,
	) ([]int64, error)
}

// SlashingKeeper represents the expected keeper interface for the (exo-)slashing module.
type SlashingKeeper interface {
	SlashWithInfractionReason(
		sdk.Context, sdk.AccAddress, int64,
		int64, sdk.Dec, stakingtypes.Infraction,
	) math.Int
}