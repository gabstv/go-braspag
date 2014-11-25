package braspag

import (
//"encoding/xml"
)

type OrderDataRequest struct {
	MerchantId     string
	OrderId        string
	BraspagOrderId string
}

type PaymentDataRequest struct {
	PaymentMethod            int16
	Amount                   int64
	Currency                 string                  // BRL USD EUR
	Country                  string                  // BRA
	AdditionalDataCollection []AdditionalDataRequest `xml:"AdditionalDataCollection>AdditionalDataRequest"`
}

type AdditionalDataRequest struct {
	Name  string
	Value string
}

type TransactionDataRequest struct {
	BraspagTransactionId string
	Amount               int64
	ServiceTaxAmount     string //TODO: check
}
