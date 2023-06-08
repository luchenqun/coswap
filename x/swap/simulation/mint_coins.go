package simulation

import (
	"math/rand"

	"coswap/x/swap/keeper"
	"coswap/x/swap/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintCoins(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintCoins{
			Admin: simAccount.Address.String(),
		}

		// TODO: Handling the MintCoins simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "MintCoins simulation not implemented"), nil, nil
	}
}
