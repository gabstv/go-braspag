package braspag

import (
	"code.google.com/p/go-uuid/uuid"
)

type authorizeTransactionRequest struct {
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
	//
	Currency string
	Country  string
}

type AuthorizeTxRequestDef struct {
	xmlTpl *authorizeTransactionRequest
	parent *WebService
}

type AuthorizeTxRequestCardDef struct {
	parent *AuthorizeTxRequestDef
	def    CCDef
	amount int64
	method int
}

func (def *AuthorizeTxRequestCardDef) SetInstallments(val int) *AuthorizeTxRequestCardDef {
	return def.SetNumberOfPayments(val)
}

func (def *AuthorizeTxRequestCardDef) SetNumberOfPayments(val int) *AuthorizeTxRequestCardDef {
	def.def.NumberOfPayments = val
	return def
}

func (def *AuthorizeTxRequestCardDef) SetPaymentPlan(val int) *AuthorizeTxRequestCardDef {
	def.def.PaymentPlan = val
	return def
}

func (def *AuthorizeTxRequestCardDef) SetAVista() *AuthorizeTxRequestCardDef {
	def.def.PaymentPlan = PAYMENTPLAN_AVISTA
	return def
}

func (def *AuthorizeTxRequestCardDef) SetParceladoEstabelecimento() *AuthorizeTxRequestCardDef {
	def.def.PaymentPlan = PAYMENTPLAN_PARCEL_ESTABELECIMENTO
	return def
}

func (def *AuthorizeTxRequestCardDef) SetParceladoEmissor() *AuthorizeTxRequestCardDef {
	def.def.PaymentPlan = PAYMENTPLAN_PARCEL_EMISSOR
	return def
}

func (def *AuthorizeTxRequestCardDef) SetType(transactionType int) *AuthorizeTxRequestCardDef {
	def.def.TransactionType = transactionType
	return def
}

func (def *AuthorizeTxRequestCardDef) SetHolder(holderName string) *AuthorizeTxRequestCardDef {
	def.def.CardHolder = holderName
	return def
}

func (def *AuthorizeTxRequestCardDef) SetNumber(number string) *AuthorizeTxRequestCardDef {
	def.def.CardNumber = number
	return def
}

func (def *AuthorizeTxRequestCardDef) SetCVC2(cvc2 string) *AuthorizeTxRequestCardDef {
	def.def.CardSecurityCode = cvc2
	return def
}

func (def *AuthorizeTxRequestCardDef) SetExpDate(mmyyyy TimeMMYYYY) *AuthorizeTxRequestCardDef {
	def.def.CardExpDate = mmyyyy
	return def
}

func (def *AuthorizeTxRequestCardDef) SetSaveCard(save bool) *AuthorizeTxRequestCardDef {
	def.def.SaveCard = save
	return def
}

func (def *AuthorizeTxRequestCardDef) SetCardToken(token string) *AuthorizeTxRequestCardDef {
	def.def.CardToken = token
	return def
}

func (def *AuthorizeTxRequestCardDef) SetSoftDescriptor(softDescriptor string) *AuthorizeTxRequestCardDef {
	def.def.SoftDescriptor = softDescriptor
	return def
}

func (def *AuthorizeTxRequestCardDef) SetAmount(amount float64) *AuthorizeTxRequestCardDef {
	def.amount = int64((amount * 100.0))
	return def
}

func (def *AuthorizeTxRequestCardDef) SetMethod(method int) *AuthorizeTxRequestCardDef {
	def.method = method
	return def
}

func (def *AuthorizeTxRequestCardDef) Commit() {
	pr := PayDataRequest{}
	pr.Amount = def.amount
	pr.Method = def.method
	pr.Currency = def.parent.xmlTpl.Currency
	pr.Country = def.parent.xmlTpl.Country
	pr.CCDef = def.def
	if def.parent.xmlTpl.CCPayData == nil {
		def.parent.xmlTpl.CCPayData = make([]PayDataRequest, 0)
		def.parent.xmlTpl.CCPayData = append(def.parent.xmlTpl.CCPayData, pr)
	}
}

func (def *AuthorizeTxRequestDef) CardDef() *AuthorizeTxRequestCardDef {
	child := &AuthorizeTxRequestCardDef{}
	child.parent = def
	child.def.TransactionType = TRTYPE_PRE
	child.def.NumberOfPayments = 1
	return child
}

func (ws *WebService) NewAuthorizeRequest(orderid string) *AuthorizeTxRequestDef {
	v := &AuthorizeTxRequestDef{}
	v.xmlTpl = &authorizeTransactionRequest{}
	v.xmlTpl.RequestId = uuid.New()
	v.xmlTpl.MerchantId = ws.merchantid
	v.xmlTpl.OrderId = orderid
	v.xmlTpl.Currency = "BRL"
	v.xmlTpl.Country = "BRA"
	v.parent = ws
	return v
}

func (def *AuthorizeTxRequestDef) SetCustomerId(id string) *AuthorizeTxRequestDef {
	def.xmlTpl.CustomerId = id
	return def
}

func (def *AuthorizeTxRequestDef) SetCustomerName(name string) *AuthorizeTxRequestDef {
	def.xmlTpl.CustomerName = name
	return def
}

func (def *AuthorizeTxRequestDef) SetCustomerEmail(email string) *AuthorizeTxRequestDef {
	def.xmlTpl.CustomerEmail = email
	return def
}

func (def *AuthorizeTxRequestDef) Commit() *AuthorizeTxRequestDef {
	return def
}

func (def *AuthorizeTxRequestDef) Submit() (*AuthorizeTransactionResponse, error) {
	return def.parent.authorize(def.xmlTpl)
}

//func (ws *WebService) NewauthorizeTransactionRequest(orderid string) authorizeTransactionRequest {
//	v := authorizeTransactionRequest{}
//	v.RequestId = uuid.New()
//	v.MerchantId = ws.merchantid
//	v.OrderId = orderid
//	v.Currency = "BRL"
//	v.Country = "BRA"
//	return v
//}

// TODO: replace this function
func (a *authorizeTransactionRequest) SetBoleto(amount float64, method int, def BoletoDef) {
	a.BoletoPayData.Amount = int64(amount * 100)
	a.BoletoPayData.Country = a.Country
	a.BoletoPayData.Currency = a.Currency
	a.BoletoPayData.Method = method
	a.BoletoPayData.BoletoDef = def
}
