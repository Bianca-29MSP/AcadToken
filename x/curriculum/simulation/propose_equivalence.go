package simulation

import (
	"math/rand"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgProposeEquivalence(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgProposeEquivalence{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the ProposeEquivalence simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "ProposeEquivalence simulation not implemented"), nil, nil
	}
}
