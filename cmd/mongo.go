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
	"errors"
	"filecoin-reward-syner/kv"
	"filecoin-reward-syner/mongo"
	"filecoin-reward-syner/util"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

// mongoCmd represents the mongo command
var mongoCmd = &cobra.Command{
	Use:   "mongo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Infof("badger path: %+v", badger)
		conn, err := kv.NewKvDBConn(badger, true)
		if err != nil {
			log.Error(err)
			return
		}
		filDB := kv.NewFilLedgerInstance(conn)
		mongoConn, err := mongo.NewMongoConn(mongoUri, database)
		if err != nil {
			log.Error(err)
			return
		}
		if fromHeight == 0 {
			fromHeight = filDB.GetSyncPos()
		}
		if endHeight <= fromHeight {
			log.Error(errors.New("end height is less than from height"))
			return
		}
		sigCh := make(chan os.Signal, 2)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		for endHeight > fromHeight {
			select {
			case <-cmd.Context().Done():
				log.Info("context done")
				time.Sleep(time.Second*5)
				return
			case sig := <-sigCh:
				log.Infof("signal %s captured", sig)
				time.Sleep(time.Second*5)
				return
			default:
				gas := filDB.ReadGasInfo(fromHeight)
				if err := mongoConn.SaveGasInfoBatch(cmd.Context(), util.ConvertGasKvToMongo(gas)); err != nil {
					log.Error(err)
				}
				transfers := filDB.ReadTransInfo(fromHeight)
				if err = mongoConn.SaveTransInfoBatch(cmd.Context(), util.ConvertTransferKvToMongo(transfers)); err != nil {
					log.Error(err)
				}
				_ = filDB.SaveSyncPos(fromHeight)
				fromHeight++
				log.Infof("sync progress height: %d", fromHeight)
			}
		}
		log.Info("gracefull down")
	},
}
var mongoUri string
var database string
var endHeight uint64

func init() {
	rootCmd.AddCommand(mongoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mongoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mongoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mongoCmd.Flags().StringVar(&badger, "badger", "/data/sdb/reward", "badger default path")
	mongoCmd.Flags().Uint64Var(&fromHeight, "from", 0, "list from height")
	mongoCmd.Flags().Uint64Var(&endHeight, "end", 950000, "list from height")
	mongoCmd.Flags().StringVar(&mongoUri, "mongoUri", "mongodb://admin:k1LxehGHCR8Ws@10.41.1.13:27017/?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false", "mongo uri")
	mongoCmd.Flags().StringVar(&database, "database", "filecoin", "database name")
}
