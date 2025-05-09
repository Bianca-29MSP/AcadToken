package simulation

import (
	"math/rand"

	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/keeper"
	"github.com/Bianca-29MSP/AcademicToken/x/academicnft/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgMintCourseNft(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgMintCourseNft{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the MintCourseNft simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "MintCourseNft simulation not implemented"), nil, nil
	}
}
