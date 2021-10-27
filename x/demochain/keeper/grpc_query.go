package keeper

import (
	"github.com/celestiaorg/demochain/x/demochain/types"
)

var _ types.QueryServer = Keeper{}
