package gofattureincloud

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
)

type DocumentType int

const (
	Fatture DocumentType = 1 + iota
	Proforma
	Ordini
	Preventivi
	Ndc
	Ricevute
	Ddt
)

var dTypes = [...]string{
	"Fatture",
	"Proforma",
	"Ordini",
	"Preventivi",
	"Ndc",
	"Ricevute",
	"Ddt",
}

func (m DocumentType) String() string { return dTypes[m-1] }

type DocumentsService struct {
	Creds Credentials
}

func (ds *DocumentsService) List(t DocumentType, params *DocumentListReqParams) ([]Document, error) {
	params.APIKey = ds.Creds.APIKey
	params.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	_, data, errs := request.Post(apiUrl + t.String() + "/lista").Send(params).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}
	res := DocumentListResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New(res.Error)
	}
	return res.ListaDocumenti, nil
}

func (ds *DocumentsService) GetDocument(t DocumentType, params *DocumentDetailsReqParams) (*Document, error) {
	params.APIKey = ds.Creds.APIKey
	params.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	_, data, errs := request.Post(apiUrl + t.String() + "/dettagli").Send(params).EndBytes()
	if errs != nil {
		return nil, errs[0]
	}
	res := DocumentDetailsResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	if !res.Success {
		return nil, errors.New(res.Error)
	}
	return res.DettagliDocumento, nil
}

func (ds *DocumentsService) SaveDocument(t DocumentType, d *Document) error {
	d.APIKey = ds.Creds.APIKey
	d.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	ep := "/nuovo"
	isNew := len(d.Id) > 0
	if !isNew {
		ep = "/modifica"
	}
	_, data, errs := request.Post(apiUrl + t.String() + ep).Send(d).EndBytes()
	if errs != nil {
		return errs[0]
	}
	res := DocumentSaveResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.Error)
	}
	if isNew {
		d.Id = res.NewId
	}
	return nil
}

func (ds *DocumentsService) DeleteDocument(t DocumentType, params *DocumentDetailsReqParams) error {
	params.APIKey = ds.Creds.APIKey
	params.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	_, data, errs := request.Post(apiUrl + t.String() + "/elimina").Send(params).EndBytes()
	if errs != nil {
		return errs[0]
	}
	res := DocumentSaveResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.Error)
	}
	return nil
}
