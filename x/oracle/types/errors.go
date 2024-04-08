package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/oracle module sentinel errors
var (
	ErrSample               = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidMsg           = sdkerrors.Register(ModuleName, 1, "invalid input create price")
	ErrPriceProposalIgnored = sdkerrors.Register(ModuleName, 2, "price proposal ignored")
)
