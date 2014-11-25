package braspag

import (
	"encoding/xml"
	"github.com/gabstv/go-soap"
	"log"
)

type WebService struct {
	merchantid   string
	homologation bool
	Query        wSQuery
}

type wSQuery struct {
	parent *WebService
}

func NewWebService(merchantid string, homologation bool) *WebService {
	ws := &WebService{}
	ws.merchantid = merchantid
	ws.homologation = homologation
	ws.Query.parent = ws
	return ws
}

func (ws *WebService) url(service string) string {
	if ws.homologation {
		return URLDEV + service
	}
	return URLPROD + service
}

func (ws *WebService) authorize(req *authorizeTransactionRequest) (*AuthorizeTransactionResponse, error) {
	if len(req.MerchantId) < 1 {
		req.MerchantId = ws.merchantid
	}
	env, err := soap_tpl_env("authorize", req)

	if err != nil {
		return nil, err
	}

	env, err = env.PostAdv(ws.url(SERVICE_TRANSACTION), soap.M{"SOAPAction": SOAPACTION_AUTHORIZE_TRANSACTION})

	if err != nil {
		return nil, err
	}

	txr := struct {
		XMLName                    xml.Name `xml:"AuthorizeTransactionResponse"`
		AuthorizeTransactionResult AuthorizeTransactionResponse
	}{}

	log.Println(env.Body.Data)
	err = env.Unmarshal(&txr)

	if err != nil {
		return nil, err
	}

	return &txr.AuthorizeTransactionResult, nil
}

func (ws *WebService) capturecc(req *CaptureCCReqDef) (*CaptureCreditCardTransactionResponse, error) {
	if len(req.xmlTpl.MerchantId) < 1 {
		req.xmlTpl.MerchantId = ws.merchantid
	}
	env, err := soap_tpl_env("capturecc", req.xmlTpl)

	if err != nil {
		return nil, err
	}

	env, err = env.PostAdv(ws.url(SERVICE_TRANSACTION), soap.M{"SOAPAction": SOAPACTION_CAPTURE_CC_TRANSACTION})

	if err != nil {
		return nil, err
	}

	txr := struct {
		XMLName                            xml.Name `xml:"CaptureCreditCardTransactionResponse"`
		CaptureCreditCardTransactionResult CaptureCreditCardTransactionResponse
	}{}

	log.Println(env.Body.Data)
	err = env.Unmarshal(&txr)

	if err != nil {
		return nil, err
	}

	return &txr.CaptureCreditCardTransactionResult, nil
}
