package client

import (
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"

	treasuryCli "terra/x/treasury/client/cli"

	"github.com/cosmos/cosmos-sdk/client"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group treasury queries under a subcommand
	treasuryQueryCmd := &cobra.Command{
		Use:   "treasury",
		Short: "Querying commands for the treasury module",
	}

	treasuryQueryCmd.AddCommand(client.GetCommands(
		treasuryCli.GetCmdQueryTaxRate(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryTaxCap(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryMiningRewardWeight(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryIssuance(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryTaxProceeds(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQuerySeigniorageProceeds(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryActiveClaims(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryCurrentEpoch(mc.storeKey, mc.cdc),
		treasuryCli.GetCmdQueryParams(mc.storeKey, mc.cdc),
	)...)

	return treasuryQueryCmd
}

// GetTxCmd The treasury module returns no TX commands.
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	return &cobra.Command{Hidden: true}
}
