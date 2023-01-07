package cli

import (
    "strconv"
	
	"github.com/spf13/cobra"
    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"ravi-ecommerce-test/x/raviecommercetest/types"
)

var _ = strconv.Itoa(0)

func CmdProductPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "product-price",
		Short: "Query productPrice",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryProductPriceRequest{
				
            }

            

			res, err := queryClient.ProductPrice(cmd.Context(), params)
            if err != nil {
                return err
            }

            return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

    return cmd
}
