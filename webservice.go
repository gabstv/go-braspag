package braspag

import (
	"bytes"
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

func (ws *WebService) Authorize(req AuthTxRequest) (*AuthorizeTransactionResponse, error) {
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

	txr := &AuthorizeTransactionResponse{}

	log.Println(env.Body.Data)
	err = env.Unmarshal(txr)

	if err != nil {
		return nil, err
	}

	return txr, nil
}
