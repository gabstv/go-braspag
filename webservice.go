package braspag

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

func (ws *WebService) Authorize(req AuthTxRequest) {

}
