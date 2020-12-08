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
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/web3coach/the-blockchain-bar/database"
	"github.com/web3coach/the-blockchain-bar/node"
	"os"
)

func runCmd() *cobra.Command {
	var runCmd = &cobra.Command{
		Use:   "run",
		Short: "Launches the TBB node and its HTTP API.",
		Run: func(cmd *cobra.Command, args []string) {
			miner, _ := cmd.Flags().GetString(flagMiner)
			sslEmail, _ := cmd.Flags().GetString(flagSSLEmail)
			isSSLDisabled, _ := cmd.Flags().GetBool(flagDisableSSL)
			ip, _ := cmd.Flags().GetString(flagIP)
			port, _ := cmd.Flags().GetUint64(flagPort)
			bootstrapIp, _ := cmd.Flags().GetString(flagBootstrapIp)
			bootstrapPort, _ := cmd.Flags().GetUint64(flagBootstrapPort)
			bootstrapAcc, _ := cmd.Flags().GetString(flagBootstrapAcc)

			fmt.Println("Launching TBB node and its HTTP API...")

			bootstrap := node.NewPeerNode(
				bootstrapIp,
				bootstrapPort,
				true,
				database.NewAccount(bootstrapAcc),
				false,
			)

			if !isSSLDisabled {
				port = node.HttpSSLPort
			}

			n := node.New(getDataDirFromCmd(cmd), ip, port, database.NewAccount(miner), bootstrap)
			err := n.Run(context.Background(), isSSLDisabled, sslEmail)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	addDefaultRequiredFlags(runCmd)
	runCmd.Flags().Bool(flagDisableSSL, false, "should the HTTP API SSL certificate be disabled? (default false)")
	runCmd.Flags().String(flagSSLEmail, "", "your node's HTTP SSL certificate email")
	runCmd.Flags().String(flagMiner, node.DefaultMiner, "your node's miner account to receive the block rewards")
	runCmd.Flags().String(flagIP, node.DefaultIP, "your node's public IP to communication with other peers")
	runCmd.Flags().Uint64(flagPort, node.HttpSSLPort, "your node's public HTTP port for communication with other peers (configurable if SSL is disabled)")
	runCmd.Flags().String(flagBootstrapIp, node.DefaultBootstrapIp, "default bootstrap Web3Coach's server to interconnect peers")
	runCmd.Flags().Uint64(flagBootstrapPort, node.HttpSSLPort, "default bootstrap Web3Coach's server port to interconnect peers")
	runCmd.Flags().String(flagBootstrapAcc, node.DefaultBootstrapAcc, "default bootstrap Web3Coach's Genesis account with 1M TBB tokens")

	return runCmd
}
