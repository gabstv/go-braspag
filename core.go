package braspag

import (
	"fmt"
	"time"
)

type AuthTxRequest struct {
	RequestId       string
	MerchantId      string
	OrderId         string
	BraspagOrderId  string
	CustomerId      string
	CustomerName    string
	CustomerEmail   string
	CustomerAddress Address
	DeliveryAddress Address
	BoletoPayData   PayDataRequest
	CCPayData       []PayDataRequest
}

type Address struct {
	Street     string
	Number     string
	Complement string
	District   string
	ZIPCode    string
}

type PayDataRequest struct {
	Method    string
	Amount    float64
	Currency  string
	Country   string
	BoletoDef BoletoDef
	CCDef     CCDef
}

type BoletoDef struct {
	Number         string
	Instructions   string
	ExpirationDate TimeMMDDYYYY // MM/dd/yyyy
	SoftDescriptor string       //TODO: view docs (???)
}

type CCDef struct {
	NumberOfPayments int // = erede installments
	PaymentPlan      string
	TransactionType  string
	CardHolder       string
	CardNumber       string
	CardSecurityCode string
	CardExpDate      TimeMMYYYY
	SaveCard         bool
	CardToken        string
	SoftDescriptor   string
}

type TimeMMDDYYYY string

func NewMMDDYYYY(t time.Time) TimeMMDDYYYY {
	return TimeMMDDYYYY(fmt.Sprintf("%02d/%02d/%04d", int(t.Month()), t.Day(), t.Year()))
}

type TimeMMYYYY string

func NewMMYYYY(t time.Time) TimeMMYYYY {
	return TimeMMYYYY(fmt.Sprintf("%02d/%04d", int(t.Month()), t.Year()))
}

func NewMMYYYYi(month, year int) TimeMMYYYY {
	return TimeMMYYYY(fmt.Sprintf("%02d/%04d", month, year))
}

func NewMMYYYYs(month, year string) TimeMMYYYY {
	for len(month) < 2 {
		month = "0" + month
	}
	for len(year) < 4 {
		year = "0" + year
	}
	return TimeMMYYYY(fmt.Sprintf("%s/%s", month, year))
}
