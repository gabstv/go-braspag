package braspag

import (
	"bytes"
	"fmt"
	"github.com/gabstv/go-soap"
	"log"
	"text/template"
	"time"
)

var (
	tpls = make(map[string]*template.Template)
)

func getplate(tpl string) (*template.Template, error) {
	tpl0 := tpls[tpl]
	if tpl0 == nil {
		var f string
		var err error
		switch tpl {
		case "authorize":
			f = string(templates_authorize_xml)
		case "capturecc":
			f = string(templates_capturecc_xml)
		case "query_getboletodata":
			f = string(templates_query_getboletodata_xml)
		}
		tpl0, err = template.New(tpl).Parse(f)
		if err != nil {
			return nil, err
		}
		tpls[tpl] = tpl0
	}
	return tpl0, nil
}

func soap_tpl_env(tpl string, content interface{}) (*soap.Envelope, error) {
	plate, err := getplate(tpl)
	if err != nil {
		return nil, err
	}
	buffer := new(bytes.Buffer)
	err = plate.Execute(buffer, content)

	log.Println(buffer.String())

	if err != nil {
		return nil, err
	}

	env, err := soap.Marshal(buffer.String())

	return env, err
}

type Address struct {
	Street     string
	Number     string
	Complement string
	District   string
	ZIPCode    string
}

type PayDataRequest struct {
	Method    int
	Amount    int64
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
	PaymentPlan      int
	TransactionType  int
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
