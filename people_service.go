package gofattureincloud

import (
	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
)

type PeopleType int

const (
	Clienti PeopleType = 1 + iota
	Fornitori
)

var pTypes = [...]string{
	"clienti",
	"fornitori",
}

func (m PeopleType) String() string { return pTypes[m-1] }

type PeopleService struct {
	Creds Credentials
}

func (ds *PeopleService) List(t PeopleType, params *PeopleListReqParams) ([]Person, []Person, error) {
	params.APIKey = ds.Creds.APIKey
	params.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	_, data, errs := request.Post(apiUrl + t.String() + "/lista").Send(params).EndBytes()
	if errs != nil {
		return nil, nil, errs[0]
	}
	res := PeopleListResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return nil, nil, err
	}
	if !res.Success {
		return nil, nil, errors.New(res.Error)
	}
	return res.ListaClienti, res.ListaFornitori, nil
}

func (ds *PeopleService) SavePerson(t PeopleType, p *Person) error {
	p.APIKey = ds.Creds.APIKey
	p.APIUID = ds.Creds.APIUID
	request := gorequest.New()
	ep := "/nuovo"
	isNew := len(p.Id) > 0
	if !isNew {
		ep = "/modifica"
	}
	_, data, errs := request.Post(apiUrl + t.String() + ep).Send(p).EndBytes()
	if errs != nil {
		return errs[0]
	}
	res := PeopleSaveResponse{}
	if err := json.Unmarshal(data, res); err != nil {
		return err
	}
	if !res.Success {
		return errors.New(res.Error)
	}
	if isNew {
		p.Id = res.NewId
	}
	return nil
}

func (ds *PeopleService) DeletePerson(t PeopleType, params *PeopleDeleteReqParams) error {
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
