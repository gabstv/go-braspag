package braspag

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/xml"
	"github.com/gabstv/go-soap"
	"log"
)

func (ws *WebService) query_getorderdata(req *QGetOrderDataDef) (*GetOrderDataResult, error) {
	if len(req.xmlTpl.MerchantId) < 1 {
		req.xmlTpl.MerchantId = ws.merchantid
	}

	env, err := soap_tpl_env("query_getorderdata", req.xmlTpl)

	if err != nil {
		return nil, err
	}

	env, err = env.PostAdv(ws.url(SERVICE_QUERY), soap.M{"SOAPAction": SOAPACTION_QUERY_GETORDERDATA})

	if err != nil {
		return nil, err
	}

	outp := GetOrderDataResponse{}

	log.Println(env.Body.Data)

	err = env.Unmarshal(&outp)

	if err != nil {
		return nil, err
	}

	return &outp.GetOrderDataResult, nil
}

func (wsq *wSQuery) NewGetOrderData(braspagtransactionid string) *QGetOrderDataDef {
	def := &QGetOrderDataDef{}
	def.parent = wsq
	def.xmlTpl.MerchantId = wsq.parent.merchantid
	def.xmlTpl.RequestId = uuid.New()
	def.xmlTpl.BraspagTransactionId = braspagtransactionid
	return def
}

type QGetOrderDataDef struct {
	xmlTpl struct {
		RequestId            string
		MerchantId           string
		BraspagTransactionId string
	}
	parent *wSQuery
}

func (def *QGetOrderDataDef) Submit() (*GetOrderDataResult, error) {
	return def.parent.parent.query_getorderdata(def)
}

type GetOrderDataResponse struct {
	XMLName            xml.Name `xml:"GetOrderDataResponse"`
	GetOrderDataResult GetOrderDataResult
}

type GetOrderDataResult struct {
	CorrelationId string
	Success       bool
	Errors        []ErrorReportDataResponse      `xml:"ErrorReportDataCollection>ErrorReportDataResponse"`
	Transactions  []OrderTransactionDataResponse `xml:"TransactionDataCollection>OrderTransactionDataResponse"`
}
