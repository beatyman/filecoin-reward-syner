package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TransferModel struct {
	UUID      string               `json:"uuid" bson:"uuid" `
	Date      time.Time            `json:"date" bson:"date"`
	Height    uint64               `json:"height" bson:"height"`
	MsgCid    string               `json:"cid" bson:"cid"`
	MsgType   string               `json:"msg_type" bson:"msg_type"`
	ActorName string               `json:"actor_name" bson:"actor_name"`
	Method    uint64               `json:"method" bson:"method"`
	From      string               `json:"from" bson:"from"`
	FromId    string               `json:"fromId" bson:"fromId"`
	To        string               `json:"to" bson:"to"`
	ToId      string               `json:"toId" bson:"toId"`
	Type      string               `json:"type" bson:"type"` //burn,transfer,miner-tip
	Value     string               `json:"value" bson:"value"`
	ToFIL     string               `json:"fil" bson:"fil"`
	Amount    primitive.Decimal128 `json:"amount" bson:"amount"` // 维度
}

type Message struct {
	UUID               string               `json:"uuid" bson:"uuid"`
	Date               time.Time            `json:"date" bson:"date"`
	Height             uint64               `json:"height" bson:"height"`
	MsgCid             string               `json:"cid" bson:"cid"`
	Version            uint64               `json:"version" bson:"version"`
	Nonce              uint64               `json:"nonce" bson:"nonce"`
	GasLimit           int64                `json:"gas_limit" bson:"gas_limit"`
	GasFeeCap          primitive.Decimal128 `json:"gas_fee_cap" bson:"gas_fee_cap"`
	GasPremium         primitive.Decimal128 `json:"gas_premium" bson:"gas_premium"`
	BaseFeeBurn        primitive.Decimal128 `json:"base_fee_burn" bson:"base_fee_burn"`
	OverEstimationBurn primitive.Decimal128 `json:"over_estimation_burn" bson:"over_estimation_burn"`
	MinerPenalty       primitive.Decimal128 `json:"miner_penalty" bson:"miner_penalty"`
	MinerTip           primitive.Decimal128 `json:"miner_tip" bson:"miner_tip"`
	Refund             primitive.Decimal128 `json:"refund" bson:"refund"`
	GasRefund          int64                `json:"gas_refund" bson:"gas_refund"`
	GasBurned          int64                `json:"gas_burned" bson:"gas_burned"`
	GasUsed            int64                `json:"gas_used" bson:"gas_used"`
	ExitCode           int64                `json:"exit_code" bson:"exit_code"`
	Transfer           []GasModel           `bson:"transfer"`
}

type GasModel struct {
	UUID   string               `json:"uuid" bson:"uuid"`
	Method uint64               `json:"method" bson:"method"`
	Date   time.Time            `json:"date" bson:"date"`
	Height uint64               `json:"height" bson:"height"`
	MsgCid string               `json:"cid" bson:"cid"`
	From   string               `json:"from" bson:"from"`
	FromId string               `json:"fromId" bson:"fromId"`
	To     string               `json:"to" bson:"to"`
	ToId   string               `json:"toId" bson:"toId"`
	Type   string               `json:"type" bson:"type"` //burn,transfer,miner-tip
	Value  string               `json:"value" bson:"value"`
	ToFIL  string               `json:"fil" bson:"fil"`
	Amount primitive.Decimal128 `json:"amount" bson:"amount" ` // 维度
}
