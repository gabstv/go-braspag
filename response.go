package braspag

import (
	"encoding/xml"
)

type AuthorizeTransactionResponse struct {
	XMLName       xml.Name `xml:"AuthorizeTransactionResult"`
	CorrelationId string
	Success       bool
	Errors        []ErrorReportDataResponse `xml:"ErrorReportDataCollection>ErrorReportDataResponse"`
	OrderData     OrderDataResponse
	PaymentData   []PaymentDataResponse `xml:"PaymentDataCollection>PaymentDataResponse"`
}

type ErrorReportDataResponse struct {
	ErrorCode    string
	ErrorMessage string
}

type OrderDataResponse struct {
	OrderId        string
	BraspagOrderId string
}

type PaymentDataResponse struct {
	BraspagTransactionId string
	Amount               int64
	PaymentMethod        int
}

type CreditCardDataResponse struct {
	AcquirerTransactionId string
	AuthorizationCode     string
	ReturnCode            string
	ReturnMessage         string
	ProofOfSale           string
	Status                int
	CreditCardToken       string
	ServiceTaxAmount      string
	AuthenticationUrl     string
}

type DebitCardDataResponse struct {
	AcquirerTransactionId string
	ReturnCode            string
	ReturnMessage         string
	Status                int
	AuthenticationUrl     string
}

type BoletoDataResponse struct {
	BoletoNumber         string
	BoletoExpirationDate string
	BoletoUrl            string
	BarCodeNumber        string
	Assigner             string
	Message              string
}

type CaptureCreditCardTransactionResponse struct {
	XMLName       xml.Name `xml:"CaptureCreditCardTransactionResult"`
	CorrelationId string
	Success       bool
	Errors        []ErrorReportDataResponse `xml:"ErrorReportDataCollection>ErrorReportDataResponse"`
	Transactions  []TransactionDataResponse `xml:"TransactionDataCollection>TransactionDataResponse"`
}

type TransactionDataResponse struct {
	BraspagTransactionId  string
	AcquirerTransactionId string
	Amount                int64
	AuthorizationCode     string
	ProofOfSale           string
	ReturnCode            string
	ReturnMessage         string
	Status                int
	ServiceTaxAmount      string
}
