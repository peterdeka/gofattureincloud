package gofattureincloud

const apiUrl = "https://api.fattureincloud.it:443/v1/"

type DocumentListReqParams struct {
	AnnoCompetenza int    `json:"anno_competenza"`
	APIKey         string `json:"api_key"`
	APIUID         string `json:"api_uid"`
	Cliente        string `json:"cliente"`
	DataFine       string `json:"data_fine"`
	DataInizio     string `json:"data_inizio"`
	IDCliente      string `json:"id_cliente"`
	Oggetto        string `json:"oggetto"`
	Saldato        string `json:"saldato"`
}

type DocumentListResponse struct {
	Success        bool       `json:"success"`
	Error          string     `json:"error"`
	ErrorCode      int        `json:"error_code"`
	ListaDocumenti []Document `json:"lista_documenti"`
}

type DocumentDetailsReqParams struct {
	APIKey string `json:"api_key"`
	APIUID string `json:"api_uid"`
	Id     string `json:"id"`
}

type DocumentDetailsResponse struct {
	Success           bool      `json:"success"`
	Error             string    `json:"error"`
	ErrorCode         int       `json:"error_code"`
	DettagliDocumento *Document `json:"dettagli_documento"`
}

type DocumentSaveResponse struct {
	Success   bool   `json:"success"`
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
	NewId     string `json:"new_id"`
}
