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
	Gas   datastore.Datastore // gas 费用
	Trans datastore.Datastore // 流水信息
}

func NewFilLedgerInstance(dstore datastore.Batching) *FilLedger {
	return &FilLedger{
		Gas:   namespace.Wrap(dstore, datastore.NewKey("/ledger/gas")),
		Trans: namespace.Wrap(dstore, datastore.NewKey("/ledger/trans")),
	}
}

func (fil *FilLedger) ReadGasInfo(height uint64) {
	iter, _ := fil.Gas.Query(query.Query{Prefix: fmt.Sprintf("/%d/", height), Orders: []query.Order{query.OrderByKey{}}})
	for r := range iter.Next() {
		if r.Error != nil {
			log.Warn(r.Error)
			continue
		}
		log.Infof("%+v", gconv.String(r.Key))
		var gas Message
		_ = json.Unmarshal(r.Value, &gas)
		log.Infof("%+v", gas)
	}
}

func (fil *FilLedger) ReadTransInfo(height uint64) {
	iter, _ := fil.Trans.Query(query.Query{Prefix: fmt.Sprintf("/%d/", height), Orders: []query.Order{query.OrderByKey{}}})
	for r := range iter.Next() {
		if r.Error != nil {
			log.Warn(r.Error)
			continue
		}
		log.Infof("%+v", gconv.String(r.Key))
		var tran TransferModel
		_ = json.Unmarshal(r.Value, &tran)
		log.Infof("%+v", tran)
	}
}
