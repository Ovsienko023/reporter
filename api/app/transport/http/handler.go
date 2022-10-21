package http

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/transport/http/handlers"
	"net/http"
)

type Transport struct {
	core core.Core
}

func NewTransport(c core.Core) *Transport {
	return &Transport{
		core: c,
	}
}

// AUTH

func (t *Transport) GetToken(w http.ResponseWriter, r *http.Request) {
	handlers.GetToken(&t.core, w, r)
}

// REPORTS

func (t *Transport) GetReport(w http.ResponseWriter, r *http.Request) {
	handlers.GetReport(&t.core, w, r)
}

func (t *Transport) GetReports(w http.ResponseWriter, r *http.Request) {
	handlers.GetReports(&t.core, w, r)
}

func (t *Transport) CreateReport(w http.ResponseWriter, r *http.Request) {
	handlers.CreateReport(&t.core, w, r)
}

func (t *Transport) UpdateReport(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateReport(&t.core, w, r)
}

func (t *Transport) DeleteReport(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteReport(&t.core, w, r)
}
