package braspag

import (
	"bytes"
	"encoding/xml"
	"github.com/gabstv/go-soap"
	"log"
)

type WebService struct {
	merchantid   string
	homologation bool
}

func NewWebService(merchantid string, homologation bool) *WebService {
	return &WebService{merchantid, homologation}
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
	plate, err := getplate("authorize")
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	err = plate.Execute(buffer, req)

	log.Println(buffer.String())

	if err != nil {
		return nil, err
	}

	env, err := soap.Marshal(buffer.String())

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
	plate, err := getplate("capturecc")
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	err = plate.Execute(buffer, req.xmlTpl)

	log.Println(buffer.String())

	if err != nil {
		return nil, err
	}

	env, err := soap.Marshal(buffer.String())

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
