package keeper

import (
	"context"

	"github.com/cdbo/brain/x/membership/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateStatus(goCtx context.Context, msg *types.MsgUpdateStatus) (*types.MsgUpdateStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Target member must have a valid address
	target, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, err
	}

	// Execute the status update
	k.UpdateMemberStatus(ctx, target, msg.Status)

	return &types.MsgUpdateStatusResponse{}, nil
}
