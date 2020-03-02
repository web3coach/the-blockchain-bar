package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/web3coach/the-blockchain-bar/fs"
	"os"
)

const flagDataDir = "datadir"
const flagMiner = "miner"
const flagIP = "ip"
const flagPort = "port"
const flagBootstrapAcc = "bootstrap-account"
const flagBootstrapIp = "bootstrap-ip"
const flagBootstrapPort = "bootstrap-port"

func main() {
	var tbbCmd = &cobra.Command{
		Use:   "tbb",
		Short: "The Blockchain Bar CLI",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	tbbCmd.AddCommand(versionCmd)
	tbbCmd.AddCommand(balancesCmd())
	tbbCmd.AddCommand(walletCmd())
	tbbCmd.AddCommand(runCmd())
	tbbCmd.AddCommand(balancesCmd())

	err := tbbCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addDefaultRequiredFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagDataDir, "", "Absolute path to the node data dir where the DB will/is stored")
	cmd.MarkFlagRequired(flagDataDir)
}

func getDataDirFromCmd(cmd *cobra.Command) string {
	dataDir, _ := cmd.Flags().GetString(flagDataDir)

	return fs.ExpandPath(dataDir)
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
