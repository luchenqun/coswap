package swap

import (
	"math/rand"

	"coswap/testutil/sample"
	swapsimulation "coswap/x/swap/simulation"
	"coswap/x/swap/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = swapsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgMintCoins = "op_weight_msg_mint_coins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgMintCoins int = 100

	opWeightMsgDistributeCoins = "op_weight_msg_distribute_coins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDistributeCoins int = 100

	opWeightMsgAddLiquidity = "op_weight_msg_add_liquidity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddLiquidity int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	swapGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&swapGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	swapParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyAdmin), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(swapParams.Admin))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyCommissionRate), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(swapParams.CommissionRate))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgMintCoins int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgMintCoins, &weightMsgMintCoins, nil,
		func(_ *rand.Rand) {
			weightMsgMintCoins = defaultWeightMsgMintCoins
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgMintCoins,
		swapsimulation.SimulateMsgMintCoins(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDistributeCoins int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDistributeCoins, &weightMsgDistributeCoins, nil,
		func(_ *rand.Rand) {
			weightMsgDistributeCoins = defaultWeightMsgDistributeCoins
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDistributeCoins,
		swapsimulation.SimulateMsgDistributeCoins(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddLiquidity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddLiquidity, &weightMsgAddLiquidity, nil,
		func(_ *rand.Rand) {
			weightMsgAddLiquidity = defaultWeightMsgAddLiquidity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddLiquidity,
		swapsimulation.SimulateMsgAddLiquidity(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
