package kv

import (
	"encoding/json"
	"fmt"
	dgbadger "github.com/dgraph-io/badger/v2"
	"github.com/gogf/gf/util/gconv"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	"github.com/ipfs/go-datastore/query"
	badger "github.com/ipfs/go-ds-badger2"
	log "github.com/sirupsen/logrus"
	"os"
)

func NewKvDBConn(path string, readOnly bool) (datastore.Batching, error) {
	if _, err := os.Stat(path); err != nil {
		log.Error(err)
		return nil, err
	}
	opts := badger.DefaultOptions
	opts.ReadOnly = readOnly
	opts.Options = dgbadger.DefaultOptions("").WithTruncate(true).WithValueThreshold(1 << 10)
	ds, err := badger.NewDatastore(path, &opts)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return ds, nil
}

type FilLedger struct {
	Gas     datastore.Datastore // gas 费用
	Trans   datastore.Datastore // 流水信息
	SyncPos datastore.Datastore // gas索引 存储同步位置
}

func NewFilLedgerInstance(dstore datastore.Batching) *FilLedger {
	return &FilLedger{
		Gas:     namespace.Wrap(dstore, datastore.NewKey("/ledger/gas")),
		Trans:   namespace.Wrap(dstore, datastore.NewKey("/ledger/trans")),
		SyncPos: namespace.Wrap(dstore, datastore.NewKey("/ledger/syncPos")),
	}
}

func (fil *FilLedger) ReadGasInfo(height uint64) []Message {
	result := make([]Message, 0)
	iter, _ := fil.Gas.Query(query.Query{Prefix: fmt.Sprintf("/%d/", height), Orders: []query.Order{query.OrderByKey{}}})
	for r := range iter.Next() {
		if r.Error != nil {
			log.Warn(r.Error)
			continue
		}
		var gas Message
		_ = json.Unmarshal(r.Value, &gas)
		result = append(result, gas)
	}
	return result
}

func (fil *FilLedger) ReadTransInfo(height uint64) []TransferModel {
	result := make([]TransferModel, 0)
	iter, _ := fil.Trans.Query(query.Query{Prefix: fmt.Sprintf("/%d/", height), Orders: []query.Order{query.OrderByKey{}}})
	for r := range iter.Next() {
		if r.Error != nil {
			log.Warn(r.Error)
			continue
		}
		var tran TransferModel
		_ = json.Unmarshal(r.Value, &tran)
		result = append(result, tran)
	}
	return result
}

func (fil *FilLedger)GetSyncPos()uint64  {
	gasHeight, err := fil.SyncPos.Get(datastore.NewKey("gasHeight"))
	if err!=nil{
		return 0
	}
	syncHeight := gconv.Uint64(gasHeight)
	return syncHeight
}
func (fil *FilLedger)SaveSyncPos(height uint64) error {
	err:=fil.SyncPos.Put(datastore.NewKey("gasHeight"), gconv.Bytes(height))
	if err!=nil{
		log.Error(err)
		return err
	}
	return nil
}