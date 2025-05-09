package keeper

import (
	"context"
	"fmt"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) RegisterInstitution(goCtx context.Context, msg *types.MsgRegisterInstitution) (*types.MsgRegisterInstitutionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Verificar se a instituição já existe
	_, found := k.Keeper.GetInstitution(goCtx, msg.Address)
	if found {
		return nil, fmt.Errorf("instituição já registrada: %w", sdkerrors.ErrInvalidRequest)
	}

	// Criar nova instituição
	institution := types.Institution{
		Index:        msg.Address, // Usar Address como o índice conforme gerado pelo Ignite
		Address:      msg.Address,
		Name:         msg.Name,
		IsAuthorized: true, // Autorizada por padrão quando registrada
	}

	// Salvar a instituição no estado
	k.Keeper.SetInstitution(goCtx, institution)

	// Emitir evento para auditoria
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"institution_registered",
			sdk.NewAttribute("address", msg.Address),
			sdk.NewAttribute("name", msg.Name),
			sdk.NewAttribute("creator", msg.Creator),
		),
	)

	return &types.MsgRegisterInstitutionResponse{}, nil
}
