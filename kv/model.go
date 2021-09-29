package kv

import (
	"github.com/shopspring/decimal"
	"time"
)

type Message struct {
	UUID               string          `gorm:"column:uuid;unique_index;primary_key;not null" json:"uuid"`
	Date               time.Time       `gorm:"column:date;index" json:"date"`
	Height             uint64          `gorm:"column:height" json:"height"`
	MsgCid             string          `gorm:"column:cid" json:"cid"`
	Version            uint64          `gorm:"column:version" json:"version"`
	Nonce              uint64          `gorm:"column:nonce" json:"nonce"`
	GasLimit           int64           `gorm:"column:gas_limit" json:"gas_limit"`
	GasFeeCap          decimal.Decimal `json:"gas_fee_cap" gorm:"column:gas_fee_cap;type:decimal(32,2)"`
	GasPremium         decimal.Decimal `json:"gas_premium" gorm:"column:gas_premium;type:decimal(32,2)"`
	BaseFeeBurn        decimal.Decimal `json:"base_fee_burn" gorm:"column:base_fee_burn;type:decimal(32,2)"`
	OverEstimationBurn decimal.Decimal `json:"over_estimation_burn" gorm:"column:over_estimation_burn;type:decimal(32,2)"`
	MinerPenalty       decimal.Decimal `json:"miner_penalty" gorm:"column:miner_penalty;type:decimal(32,2)"`
	MinerTip           decimal.Decimal `json:"miner_tip" gorm:"column:miner_tip;type:decimal(32,2)"`
	Refund             decimal.Decimal `json:"refund" gorm:"column:refund;type:decimal(32,2)"`
	GasRefund          int64           `gorm:"column:gas_refund" json:"gas_refund"`
	GasBurned          int64           `gorm:"column:gas_burned" json:"gas_burned"`
	GasUsed            int64           `gorm:"column:gas_used" json:"gas_used"`
	ExitCode           int64           `gorm:"column:exit_code" json:"exit_code"`
	Transfer           []GasModel      `bson:"transfer" gorm:"-"`
}

func (Message) TableName() string {
	return "message_info"
}
type GasModel struct {
	UUID   string          `gorm:"column:uuid;unique_index;primary_key;not null" json:"uuid"`
	Method uint64          `gorm:"column:method" json:"method"`
	Date   time.Time       `gorm:"column:date;index" json:"date"`
	Height uint64          `gorm:"column:height" json:"height"`
	MsgCid string          `gorm:"column:cid" json:"cid"`
	From   string          `gorm:"column:from" json:"from"`
	FromId string          `gorm:"column:fromId" json:"fromId"`
	To     string          `gorm:"column:to" json:"to"`
	ToId   string          `gorm:"column:toId" json:"toId"`
	Type   string          `gorm:"column:type" json:"type"` //burn,transfer,miner-tip
	Value  string          `gorm:"column:value" json:"value"`
	ToFIL  string          `gorm:"column:fil" json:"fil"`
	Amount decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(32,18)"` // 维度

}

func (GasModel) TableName() string {
	return "gas_info"
}

type TransferModel struct {
	UUID      string          `gorm:"column:uuid;unique_index;primary_key;not null" json:"uuid"`
	Date      time.Time       `gorm:"column:date;index" json:"date"`
	Height    uint64          `gorm:"column:height" json:"height"`
	MsgCid    string          `gorm:"column:cid" json:"cid"`
	MsgType   string          `gorm:"column:msg_type" json:"msg_type"`
	ActorName string          `gorm:"column:actor_name" json:"actor_name"`
	Method    uint64          `gorm:"column:method" json:"method"`
	From      string          `gorm:"column:from" json:"from"`
	FromId    string          `gorm:"column:fromId" json:"fromId"`
	To        string          `gorm:"column:to" json:"to"`
	ToId      string          `gorm:"column:toId" json:"toId"`
	Type      string          `gorm:"column:type" json:"type"` //burn,transfer,miner-tip
	Value     string          `gorm:"column:value" json:"value"`
	ToFIL     string          `gorm:"column:fil" json:"fil"`
	Amount    decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(32,18)"` // 维度
}

func (TransferModel) TableName() string {
	return "transfer_info"
}