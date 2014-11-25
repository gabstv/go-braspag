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
	BoletoPayData   *PayDataRequest
	CCPayData       []*PayDataRequest
	//
	Currency string
	Country  string
}

type AuTxReqDef struct {
	xmlTpl *authorizeTransactionRequest
	parent *WebService
}

type AuTxReqDefCC struct {
	parent *AuTxReqDef
	xmlTpl *PayDataRequest
}

type AuTxReqDefBoleto struct {
	parent *AuTxReqDef
	xmlTpl *PayDataRequest
}

func (def *AuTxReqDefCC) SetInstallments(val int) *AuTxReqDefCC {
	return def.SetNumberOfPayments(val)
}

func (def *AuTxReqDefCC) SetNumberOfPayments(val int) *AuTxReqDefCC {
	def.xmlTpl.CCDef.NumberOfPayments = val
	return def
}

func (def *AuTxReqDefCC) SetPaymentPlan(val int) *AuTxReqDefCC {
	def.xmlTpl.CCDef.PaymentPlan = val
	return def
}

func (def *AuTxReqDefCC) SetAVista() *AuTxReqDefCC {
	def.xmlTpl.CCDef.PaymentPlan = PAYMENTPLAN_AVISTA
	return def
}

func (def *AuTxReqDefCC) SetParceladoEstabelecimento() *AuTxReqDefCC {
	def.xmlTpl.CCDef.PaymentPlan = PAYMENTPLAN_PARCEL_ESTABELECIMENTO
	return def
}

func (def *AuTxReqDefCC) SetParceladoEmissor() *AuTxReqDefCC {
	def.xmlTpl.CCDef.PaymentPlan = PAYMENTPLAN_PARCEL_EMISSOR
	return def
}

func (def *AuTxReqDefCC) SetType(transactionType int) *AuTxReqDefCC {
	def.xmlTpl.CCDef.TransactionType = transactionType
	return def
}

func (def *AuTxReqDefCC) SetHolder(holderName string) *AuTxReqDefCC {
	def.xmlTpl.CCDef.CardHolder = holderName
	return def
}

func (def *AuTxReqDefCC) SetNumber(number string) *AuTxReqDefCC {
	def.xmlTpl.CCDef.CardNumber = number
	return def
}

func (def *AuTxReqDefCC) SetCVC2(cvc2 string) *AuTxReqDefCC {
	def.xmlTpl.CCDef.CardSecurityCode = cvc2
	return def
}

func (def *AuTxReqDefCC) SetExpDate(mmyyyy TimeMMYYYY) *AuTxReqDefCC {
	def.xmlTpl.CCDef.CardExpDate = mmyyyy
	return def
}

func (def *AuTxReqDefCC) SetSaveCard(save bool) *AuTxReqDefCC {
	def.xmlTpl.CCDef.SaveCard = save
	return def
}

func (def *AuTxReqDefCC) SetCardToken(token string) *AuTxReqDefCC {
	def.xmlTpl.CCDef.CardToken = token
	return def
}

func (def *AuTxReqDefCC) SetSoftDescriptor(softDescriptor string) *AuTxReqDefCC {
	def.xmlTpl.CCDef.SoftDescriptor = softDescriptor
	return def
}

func (def *AuTxReqDefCC) SetAmount(amount float64) *AuTxReqDefCC {
	def.xmlTpl.Amount = int64((amount * 100.0))
	return def
}

func (def *AuTxReqDefCC) SetMethod(method int) *AuTxReqDefCC {
	def.xmlTpl.Method = method
	return def
}

func (def *AuTxReqDefCC) Parent() *AuTxReqDef {
	return def.parent
}

func (def *AuTxReqDef) NewCard() *AuTxReqDefCC {
	child := &AuTxReqDefCC{}
	child.xmlTpl = &PayDataRequest{}
	if def.xmlTpl.CCPayData == nil {
		def.xmlTpl.CCPayData = make([]*PayDataRequest, 0)
	}
	def.xmlTpl.CCPayData = append(def.xmlTpl.CCPayData, child.xmlTpl)
	//
	child.xmlTpl.Currency = def.xmlTpl.Currency
	child.xmlTpl.Country = def.xmlTpl.Country
	//
	child.parent = def
	child.xmlTpl.CCDef.TransactionType = TRTYPE_PRE
	child.xmlTpl.CCDef.NumberOfPayments = 1
	return child
}

func (def *AuTxReqDef) NewBoleto() *AuTxReqDefBoleto {
	child := &AuTxReqDefBoleto{}
	child.xmlTpl = &PayDataRequest{}
	def.xmlTpl.BoletoPayData = child.xmlTpl
	//
	child.xmlTpl.Currency = def.xmlTpl.Currency
	child.xmlTpl.Country = def.xmlTpl.Country
	//
	child.parent = def
	return child
}

func (def *AuTxReqDefBoleto) SetAmount(amount float64) *AuTxReqDefBoleto {
	def.xmlTpl.Amount = int64((amount * 100.0))
	return def
}

func (def *AuTxReqDefBoleto) SetMethod(method int) *AuTxReqDefBoleto {
	def.xmlTpl.Method = method
	return def
}

// TODO: SetNumber

func (def *AuTxReqDefBoleto) SetInstructions(instructions string) *AuTxReqDefBoleto {
	def.xmlTpl.BoletoDef.Instructions = instructions
	return def
}

func (def *AuTxReqDefBoleto) SetExpirationDate(expdate TimeMMDDYYYY) *AuTxReqDefBoleto {
	def.xmlTpl.BoletoDef.ExpirationDate = expdate
	return def
}

// TODO: SetSoftDescriptor

func (def *AuTxReqDefBoleto) Parent() *AuTxReqDef {
	return def.parent
}

func (ws *WebService) NewAuthorizeRequest(orderid string) *AuTxReqDef {
	v := &AuTxReqDef{}
	v.xmlTpl = &authorizeTransactionRequest{}
	v.xmlTpl.RequestId = uuid.New()
	v.xmlTpl.MerchantId = ws.merchantid
	v.xmlTpl.OrderId = orderid
	v.xmlTpl.Currency = "BRL"
	v.xmlTpl.Country = "BRA"
	v.parent = ws
	return v
}

func (def *AuTxReqDef) SetCustomerId(id string) *AuTxReqDef {
	def.xmlTpl.CustomerId = id
	return def
}

func (def *AuTxReqDef) SetCustomerName(name string) *AuTxReqDef {
	def.xmlTpl.CustomerName = name
	return def
}

func (def *AuTxReqDef) SetCustomerEmail(email string) *AuTxReqDef {
	def.xmlTpl.CustomerEmail = email
	return def
}

func (def *AuTxReqDef) Submit() (*AuthorizeTransactionResponse, error) {
	return def.parent.authorize(def.xmlTpl)
}

// TODO: replace this function
func (a *authorizeTransactionRequest) SetBoleto(amount float64, method int, def BoletoDef) {
	a.BoletoPayData.Amount = int64(amount * 100)
	a.BoletoPayData.Country = a.Country
	a.BoletoPayData.Currency = a.Currency
	a.BoletoPayData.Method = method
	a.BoletoPayData.BoletoDef = def
}
