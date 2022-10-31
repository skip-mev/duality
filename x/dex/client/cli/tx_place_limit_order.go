package cli

import (
	"strconv"

	"github.com/NicholasDotSol/duality/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdPlaceLimitOrder() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "place-limit-order [receiver] [token-a] [token-b] [tick-index] [token-in] [amount-in]",
		Short: "Broadcast message PlaceLimitOrder",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argReceiver := args[0]
			argTokenA := args[1]
			argTokenB := args[2]
			argTickIndex := args[3]

			argTickIndexInt, err := strconv.Atoi(argTickIndex)

			if err != nil {
				return err
			}

			argTokenIn := args[4]
			argAmountIn := args[5]

			argAmountInDec, err := sdk.NewDecFromStr(argAmountIn)

			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgPlaceLimitOrder(
				clientCtx.GetFromAddress().String(),
				argReceiver,
				argTokenA,
				argTokenB,
				int64(argTickIndexInt),
				argTokenIn,
				argAmountInDec,
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
