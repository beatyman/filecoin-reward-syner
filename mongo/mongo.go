package mongo

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/qiniu/qmgo/options"
	log "github.com/sirupsen/logrus"
)

func NewMongoConn(uri string, database string) (*FilMongoSink, error) {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri:      uri,
		Database: database,
	})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	gasColl := client.Database(database).Collection("gas")
	transColl := client.Database(database).Collection("transfer")
	if err := gasColl.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"height,cid"}, Unique: true}); err != nil {
		log.Error(err)
	}
	if err := transColl.CreateOneIndex(context.Background(), options.IndexModel{Key: []string{"height,cid"}, Unique: true}); err != nil {
		log.Error(err)
	}
	return &FilMongoSink{
		DB:        client.Database(database),
		gasColl:   gasColl,
		transColl: transColl,
	}, nil
}

type FilMongoSink struct {
	DB        *qmgo.Database
	gasColl   *qmgo.Collection
	transColl *qmgo.Collection
}

func (fil *FilMongoSink) SaveGasInfoBatch(ctx context.Context, list []Message) error {
	result, err := fil.gasColl.InsertMany(ctx, list)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("%+v", result)
	return nil
}

func (fil *FilMongoSink) SaveTransInfoBatch(ctx context.Context, list []TransferModel) error {
	result, err := fil.transColl.InsertMany(ctx, list)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("%+v", result)
	return nil
}