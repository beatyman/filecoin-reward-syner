package util

import (
	"filecoin-reward-syner/kv"
	"filecoin-reward-syner/mongo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertGasKvToMongo(messages []kv.Message) []mongo.Message {
	result := make([]mongo.Message, 0)
	for i := range messages {
		details := make([]mongo.GasModel, 0)
		for j := range messages[i].Transfer {
			detail := mongo.GasModel{
				UUID:   messages[i].Transfer[j].UUID,
				Method: messages[i].Transfer[j].Method,
				Date:   messages[i].Transfer[j].Date,
				Height: messages[i].Transfer[j].Height,
				MsgCid: messages[i].Transfer[j].MsgCid,
				From:   messages[i].Transfer[j].From,
				FromId: messages[i].Transfer[j].FromId,
				To:     messages[i].Transfer[j].To,
				ToId:   messages[i].Transfer[j].ToId,
				Type:   messages[i].Transfer[j].Type,
				Value:  messages[i].Transfer[j].Value,
				ToFIL:  messages[i].Transfer[j].ToFIL,
				Amount: convert(messages[i].Transfer[j].Amount),
			}
			details = append(details, detail)
		}
		msg := mongo.Message{
			UUID:               messages[i].UUID,
			Date:               messages[i].Date,
			Height:             messages[i].Height,
			MsgCid:             messages[i].MsgCid,
			Version:            messages[i].Version,
			Nonce:              messages[i].Nonce,
			GasLimit:           messages[i].GasLimit,
			GasFeeCap:          convert(messages[i].GasFeeCap),
			GasPremium:         convert(messages[i].GasPremium),
			BaseFeeBurn:        convert(messages[i].BaseFeeBurn),
			OverEstimationBurn: convert(messages[i].OverEstimationBurn),
			MinerPenalty:       convert(messages[i].MinerPenalty),
			MinerTip:           convert(messages[i].MinerTip),
			Refund:             convert(messages[i].Refund),
			GasRefund:          messages[i].GasRefund,
			GasBurned:          messages[i].GasBurned,
			GasUsed:            messages[i].GasUsed,
			ExitCode:           messages[i].ExitCode,
			Transfer:           details,
		}
		result = append(result, msg)
	}
	return result
}

func ConvertTransferKvToMongo(transfers []kv.TransferModel) []mongo.TransferModel {
	result := make([]mongo.TransferModel, 0)
	for i := range transfers {
		transfer := mongo.TransferModel{
			UUID:      transfers[i].UUID,
			Date:      transfers[i].Date,
			Height:    transfers[i].Height,
			MsgCid:    transfers[i].MsgCid,
			MsgType:   transfers[i].MsgType,
			ActorName: transfers[i].ActorName,
			Method:    transfers[i].Method,
			From:      transfers[i].From,
			FromId:    transfers[i].FromId,
			To:        transfers[i].To,
			ToId:      transfers[i].ToId,
			Type:      transfers[i].Type,
			Value:     transfers[i].Value,
			ToFIL:     transfers[i].ToFIL,
			Amount:    convert(transfers[i].Amount),
		}
		result = append(result, transfer)
	}
	return result
}

func convert(origin decimal.Decimal) primitive.Decimal128 {
	newValue, err := primitive.ParseDecimal128(origin.String())
	if err != nil {
		log.Error(err)
	}
	return newValue
}
