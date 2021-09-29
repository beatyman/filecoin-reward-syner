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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("badger path: %+v",badger)
		conn,err:=kv.NewKvDBConn(badger,true)
		if err!=nil{
			log.Error(err)
			return
		}
		filDB:=kv.NewFilLedgerInstance(conn)
		log.Infof("sync pos: ",filDB.GetSyncPos())
		filDB.ReadGasInfo(fromHeight)
		filDB.ReadTransInfo(fromHeight)
	},
}
var badger string
var fromHeight uint64
func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().StringVar(&badger, "badger", "/data/sdb/reward", "badger default path")
	listCmd.Flags().Uint64Var(&fromHeight, "from", 0, "list from height")
}