/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"filecoin-reward-syner/kv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := kv.NewKvDBConn(badger, false)
		if err != nil {
			log.Error(err)
			return
		}
		filDB := kv.NewFilLedgerInstance(conn)

		filDB.SaveSyncPos(resetHeight)
		log.Infof("reset sync pos : %d", resetHeight)
	},
}
var resetHeight uint64

func init() {
	rootCmd.AddCommand(resetCmd)

	resetCmd.Flags().Uint64Var(&resetHeight, "reset", 0, "reset sync pos")
	resetCmd.Flags().StringVar(&badger, "badger", "/data/sdb/reward", "badger default path")
}
