package cli

import (
	"strconv"

	"exchange/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdCreatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-pool [amount] [height]",
		Short: "Create a new pool",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAmount, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			argHeight, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreatePool(clientCtx.GetFromAddress().String(), argAmount, argHeight)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdatePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pool [id] [amount] [height]",
		Short: "Update a pool",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argAmount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}

			argHeight, err := cast.ToInt32E(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdatePool(clientCtx.GetFromAddress().String(), id, argAmount, argHeight)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeletePool() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-pool [id]",
		Short: "Delete a pool by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeletePool(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
