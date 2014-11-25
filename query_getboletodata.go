package braspag

import (
	"code.google.com/p/go-uuid/uuid"
	"encoding/xml"
	"github.com/gabstv/go-soap"
	"log"
)

func (ws *WebService) query_getboletodata(req *QGetBoletoDataDef) (*GetBoletoDataResult, error) {
	if len(req.xmlTpl.MerchantId) < 1 {
		req.xmlTpl.MerchantId = ws.merchantid
	}

	env, err := soap_tpl_env("query_getboletodata", req.xmlTpl)

	if err != nil {
		return nil, err
	}

	env, err = env.PostAdv(ws.url(SERVICE_QUERY), soap.M{"SOAPAction": SOAPACTION_QUERY_GETBOLETODATA})

	if err != nil {
		return nil, err
	}

	outp := GetBoletoDataResponse{}

	log.Println(env.Body.Data)

	err = env.Unmarshal(&outp)

	if err != nil {
		return nil, err
	}

	return &outp.GetBoletoDataResult, nil
}

func (wsq *wSQuery) NewGetBoletoData(braspagtransactionid string) *QGetBoletoDataDef {
	def := &QGetBoletoDataDef{}
	def.parent = wsq
	def.xmlTpl.MerchantId = wsq.parent.merchantid
	def.xmlTpl.RequestId = uuid.New()
	return def
}

type QGetBoletoDataDef struct {
	xmlTpl struct {
		RequestId            string
		MerchantId           string
		BraspagTransactionId string
	}
	parent *wSQuery
}

func (def *QGetBoletoDataDef) Submit() (*GetBoletoDataResult, error) {
	return def.parent.parent.query_getboletodata(def)
}

type GetBoletoDataResponse struct {
	XMLName             xml.Name `xml:"GetBoletoDataResponse"`
	GetBoletoDataResult GetBoletoDataResult
}

type GetBoletoDataResult struct {
	CorrelationId  string
	Success        bool
	Errors         []ErrorReportDataResponse `xml:"ErrorReportDataCollection>ErrorReportDataResponse"`
	DocumentNumber string
	Instructions   string
	Customer       string
	BoletoNumber   int64  // TODO: check type
	BarCodeNumber  string //
	DocumentDate   string // DATETIME
	ExpirationDate string // DATETIME
	BankNumber     string
	Agency         string
	Account        string
	BoletoType     string // carteira
	Amount         int64  // dividir por 100
	PaidAmount     int64  // dividir por 100
	PaymentDate    string // DATETIME
	Assignor       string // campo cedente
	BoletoUrl      string // url para montar o boleto
}
