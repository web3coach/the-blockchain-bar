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
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/console"
	"github.com/spf13/cobra"
	"github.com/web3coach/the-blockchain-bar/wallet"
	"io/ioutil"
	"os"
)

func walletCmd() *cobra.Command {
	var walletCmd = &cobra.Command{
		Use:   "wallet",
		Short: "Manages blockchain accounts and keys.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return incorrectUsageErr()
		},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	walletCmd.AddCommand(walletNewAccountCmd())
	walletCmd.AddCommand(walletPrintPrivKeyCmd())

	return walletCmd
}

func walletNewAccountCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "new-account",
		Short: "Creates a new account with a new set of a elliptic-curve Private + Public keys.",
		Run: func(cmd *cobra.Command, args []string) {
			password := getPassPhrase("Please enter a password to encrypt the new wallet:", true)
			dataDir := getDataDirFromCmd(cmd)

			acc, err := wallet.NewKeystoreAccount(dataDir, password)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			fmt.Printf("New account created: %s\n", acc.Hex())
			fmt.Printf("Saved in: %s\n", wallet.GetKeystoreDirPath(dataDir))
		},
	}

	addDefaultRequiredFlags(cmd)

	return cmd
}

func walletPrintPrivKeyCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "pk-print",
		Short: "Unlocks keystore file and prints the Private + Public keys.",
		Run: func(cmd *cobra.Command, args []string) {
			ksFile, _ := cmd.Flags().GetString(flagKeystoreFile)
			password := getPassPhrase("Please enter a password to decrypt the wallet:", false)

			keyJson, err := ioutil.ReadFile(ksFile)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			key, err := keystore.DecryptKey(keyJson, password)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			spew.Dump(key)
		},
	}

	addKeystoreFlag(cmd)

	return cmd
}

func getPassPhrase(prompt string, confirmation bool) string {
	password, err := console.Stdin.PromptPassword(prompt)
	if err != nil {
		utils.Fatalf("Failed to read password: %v", err)
	}

	if confirmation {
		confirm, err := console.Stdin.PromptPassword("Repeat password: ")
		if err != nil {
			utils.Fatalf("Failed to read password confirmation: %v", err)
		}
		if password != confirm {
			utils.Fatalf("Passwords do not match")
		}
	}

	return password
}
