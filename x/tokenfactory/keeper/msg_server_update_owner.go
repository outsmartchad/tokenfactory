package keeper

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denomFound, found := k.GetDenom(ctx, msg.Denom)
	if !found{
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}
	if msg.Owner != denomFound.Owner{
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	denom := types.Denom{
		Denom: denomFound.Denom,
		Owner: msg.NewOwner,
		Description: denomFound.Description,
		Supply: denomFound.Supply,
		Ticker: denomFound.Ticker,
		MaxSupply: denomFound.MaxSupply,
		Precision: denomFound.Precision,
		Url: denomFound.Url,
		CanChangeMaxSupply: denomFound.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateOwnerResponse{}, nil
}
