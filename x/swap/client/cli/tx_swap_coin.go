package cli

import (
	"strconv"

	"coswap/x/swap/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSwapCoin() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "swap-coin [amount-in] [amount-out-min]",
		Short: "Broadcast message swap-coin",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmountIn, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argAmountOutMin, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSwapCoin(
				clientCtx.GetFromAddress().String(),
				argAmountIn,
				argAmountOutMin,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
