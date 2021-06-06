// Copyright 2020 The the-blockchain-bar Authors
// This file is part of the the-blockchain-bar library.
//
// The the-blockchain-bar library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The the-blockchain-bar library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/node"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interacts with balances (list...).",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	balancesCmd.AddCommand(balancesListCmd())

	return balancesCmd
}

func balancesListCmd() *cobra.Command {
	var balancesListCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all balances.",
		Run: func(cmd *cobra.Command, args []string) {
			state, err := database.NewStateFromDisk(getDataDirFromCmd(cmd), node.DefaultMiningDifficulty)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			defer state.Close()

			fmt.Printf("Accounts balances at %x:\n", state.LatestBlockHash())
			fmt.Println("__________________")
			fmt.Println("")
			for account, balance := range state.Balances {
				fmt.Println(fmt.Sprintf("%s: %d", account.String(), balance))
			}
			fmt.Println("")
			fmt.Printf("Accounts nonces:")
			fmt.Println("")
			fmt.Println("__________________")
			fmt.Println("")
			for account, nonce := range state.Account2Nonce {
				fmt.Println(fmt.Sprintf("%s: %d", account.String(), nonce))
			}
		},
	}

	addDefaultRequiredFlags(balancesListCmd)

	return balancesListCmd
}
