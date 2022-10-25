package swaprouter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	gammtypes "github.com/osmosis-labs/osmosis/v12/x/gamm/types"
	"github.com/osmosis-labs/osmosis/v12/x/swaprouter/types"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	// TODO: remove nolint
	// nolint: unused
	gammKeeper types.SimulationExtension
	// TODO: remove nolint
	// nolint: unused
	concentratedKeeper types.SwapI

	paramSpace paramtypes.Subspace
}

func NewKeeper(storeKey sdk.StoreKey, paramSpace paramtypes.Subspace, gammKeeper types.SimulationExtension, concentratedKeeper types.SwapI) *Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{storeKey: storeKey, paramSpace: paramSpace, gammKeeper: gammKeeper, concentratedKeeper: concentratedKeeper}
}

// GetParams returns the total set of swaprouter parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of swaprouter parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// InitGenesis initializes the swaprouter module's state from a provided genesis
// state.
func (k Keeper) InitGenesis(ctx sdk.Context, genState *types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}

	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the swaprouter module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}

// GetPoolAndPoke gets a pool with the given pool id.
// This method is used for simulation only.
// TODO: remove it after refactoring simulation logic.
func (k Keeper) GetPoolAndPoke(ctx sdk.Context, poolId uint64) (gammtypes.TraditionalAmmInterface, error) {
	return k.gammKeeper.GetPoolAndPoke(ctx, poolId)
}

// GetNextPoolId returns the next pool id.
// This method is used for simulation only.
// TODO: remove it after refactoring simulation logic.
func (k Keeper) GetNextPoolId(ctx sdk.Context) uint64 {
	return k.gammKeeper.GetNextPoolId(ctx)
}