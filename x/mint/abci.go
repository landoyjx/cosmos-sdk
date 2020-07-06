package mint

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/mint/types"
)

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(ctx sdk.Context, k Keeper) {
	params := k.GetParams(ctx)

	if ctx.BlockHeader().Height > int64(params.BlocksPerYear*10) {
		return
	}

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	//params := k.GetParams(ctx)

	// recalculate inflation rate
	// totalStakingSupply := k.StakingTokenSupply(ctx)
	// bondedRatio := k.BondedRatio(ctx)
	// minter.Inflation = minter.NextInflationRate(params, bondedRatio)
	// minter.AnnualProvisions = minter.NextAnnualProvisions(params, totalStakingSupply)
	// k.SetMinter(ctx, minter)

	minter.AnnualProvisions = sdk.NewDec(20000000)

	// mint coins, update supply
	mintedCoin := minter.BlockProvision(params)
	mintedCoins := sdk.NewCoins(mintedCoin)

	err := k.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeMint,
			//sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			//sdk.NewAttribute(types.AttributeKeyInflation, minter.Inflation.String()),
			sdk.NewAttribute(types.AttributeKeyAnnualProvisions, minter.AnnualProvisions.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
		),
	)
}
