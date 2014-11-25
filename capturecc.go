package braspag

import (
	"code.google.com/p/go-uuid/uuid"
)

func (ws *WebService) NewCaptureCCRequest() *CaptureCCReqDef {
	v := &CaptureCCReqDef{}
	v.xmlTpl.MerchantId = ws.merchantid
	v.xmlTpl.TransactionDataCollection = make([]TransactionDataRequest, 0)
	v.xmlTpl.RequestId = uuid.New()
	v.parent = ws
	return v
}

type CaptureCCReqDef struct {
	xmlTpl struct {
		MerchantId                string
		RequestId                 string
		TransactionDataCollection []TransactionDataRequest
	}
	parent *WebService
}

func (c *CaptureCCReqDef) AddTransaction(braspagid string, amount float64) *CaptureCCReqDef {
	c.xmlTpl.TransactionDataCollection = append(c.xmlTpl.TransactionDataCollection, TransactionDataRequest{
		BraspagTransactionId: braspagid,
		Amount:               int64(amount * 100),
	})
	return c
}

func (c *CaptureCCReqDef) Submit() (*CaptureCreditCardTransactionResponse, error) {
	return c.parent.capturecc(c)
}
