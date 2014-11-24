package braspag

const (
	URLDEV  = "https://homologacao.pagador.com.br"
	URLPROD = "https://www.pagador.com.br"
	//
	SERVICE_QUERY       = "/services/pagadorQuery.asmx"
	SERVICE_TRANSACTION = "/webservice/pagadorTransaction.asmx"
	//
	SOAPACTION_AUTHORIZE_TRANSACTION = "https://www.pagador.com.br/webservice/pagador/AuthorizeTransaction"
	//
	PM_CIELO_VISAELECTRON           = 123
	PM_CIELO_VISA                   = 500
	PM_CIELO_MASTERCARD             = 501
	PM_CIELO_AMEX                   = 502
	PM_CIELO_DINERS                 = 503
	PM_CIELO_ELO                    = 504
	PM_BANORTE_VISA                 = 505
	PM_BANORTE_MASTERCARD           = 506
	PM_BANORTE_DINERS               = 507
	PM_BANORTE_AMEX                 = 508
	PM_REDECARD_VISA                = 509
	PM_REDECARD_MASTERCARD          = 510
	PM_REDECARD_DINERS              = 511
	PM_PAGOSONLINE_VISA             = 512
	PM_PAGOSONLINE_MASTERCARD       = 513
	PM_PAGOSONLINE_AMEX             = 514
	PM_PAGOSONLINE_DINERS           = 515
	PM_PAYVISION_VISA               = 516
	PM_PAYVISION_MASTERCARD         = 517
	PM_PAYVISION_DINERS             = 518
	PM_PAYVISION_AMEX               = 519
	PM_BANORTECARGOSAUTO_VISA       = 520
	PM_BANORTECARGOSAUTO_MASTERCARD = 521
	PM_BANORTECARGOSAUTO_DINERS     = 522
	PM_AMEX_2P                      = 523
	PM_SITEF_VISA                   = 524
	PM_SITEF_MASTERCARD             = 525
	PM_SITEF_AMEX                   = 526
	PM_SITEF_DINERS                 = 527
	PM_SITEF_HIPERCARD              = 528
	PM_SITEF_LEADER                 = 529
	PM_SITEF_AURA                   = 530
	PM_SITEF_SANTANDERVISA          = 531
	PM_SITEF_SANTANDERMASTERCARD    = 532
	PM_SIMULATED_USD                = 995
	PM_SIMULATED_EUR                = 996
	PM_SIMULATED_BRL                = 997
	//
	CCDRSTAT_CAPTURED       = 0 // byte? // Transação Capturada
	CCDRSTAT_AUTHORIZED     = 1 // byte? // Transação Autorizada, pendente de captura.
	CCDRSTAT_NOT_AUTHORIZED = 2 // byte? // Transação não Autorizada, pela Adquirente.
	CCDRSTAT_DEQUAL_ERROR   = 3 // byte? // Transação com erro Desqualificante.
	CCDRSTAT_WAITING        = 4 // byte? // Transação aguardando resposta.
)
