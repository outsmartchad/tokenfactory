package keeper

import (
	"context"

	"tokenfactory/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denomFound, found := k.GetDenom(ctx, msg.Denom)
	if !found{
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}
	if denomFound.Owner != msg.Owner{
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	if msg.Amount + denomFound.Supply > denomFound.MaxSupply{
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "Cannot mint more than Max Supply")
	}
	moduleAcct := k.accountKeeper.GetModuleAddress(types.ModuleName) 
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil{
		return nil, err
	}
	var mintCoins sdk.Coins
	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Amount))))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins)
	if err != nil{
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, moduleAcct, recipientAddress, mintCoins)
	if err != nil{
		return nil, err
	}
	denom := types.Denom{
		Denom: msg.Denom,
		Owner: denomFound.Owner,
		Supply: msg.Amount + denomFound.Supply,
		Description: denomFound.Description,
		Precision: denomFound.Precision,
		Ticker: denomFound.Ticker,
		Url: denomFound.Url,
		MaxSupply: denomFound.MaxSupply,
		CanChangeMaxSupply: denomFound.CanChangeMaxSupply,
	}
	k.SetDenom(ctx, denom)
	return &types.MsgMintAndSendTokensResponse{}, nil
}
