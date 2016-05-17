package gofattureincloud

type Credentials struct {
	APIKey string `json:"api_key"`
	APIUID string `json:"api_uid"`
}

//******DOCUMENTS******
type Document struct {
	AnnoCompetenza      int       `json:"anno_competenza"`
	APIKey              string    `json:"api_key"`
	APIUID              string    `json:"api_uid"`
	Cassa               int       `json:"cassa"`
	CentroRicavo        string    `json:"centro_ricavo"`
	Cf                  string    `json:"cf"`
	Data                string    `json:"data"`
	Ddt                 bool      `json:"ddt"`
	DdtAnnotazioni      string    `json:"ddt_annotazioni"`
	DdtCausale          string    `json:"ddt_causale"`
	DdtColli            string    `json:"ddt_colli"`
	DdtData             string    `json:"ddt_data"`
	DdtIDTemplate       string    `json:"ddt_id_template"`
	DdtLuogo            string    `json:"ddt_luogo"`
	DdtNumero           string    `json:"ddt_numero"`
	DdtPeso             string    `json:"ddt_peso"`
	DdtTrasportatore    string    `json:"ddt_trasportatore"`
	Ftacc               bool      `json:"ftacc"`
	FtaccIDTemplate     string    `json:"ftacc_id_template"`
	Id                  string    `json:"id,omitempty"`
	IDCliente           string    `json:"id_cliente"`
	IDTemplate          string    `json:"id_template"`
	ImponibileRitenuta  int       `json:"imponibile_ritenuta"`
	IndirizzoCap        string    `json:"indirizzo_cap"`
	IndirizzoCitta      string    `json:"indirizzo_citta"`
	IndirizzoExtra      string    `json:"indirizzo_extra"`
	IndirizzoProvincia  string    `json:"indirizzo_provincia"`
	IndirizzoVia        string    `json:"indirizzo_via"`
	Lingua              string    `json:"lingua"`
	ListaArticoli       []Article `json:"lista_articoli"`
	ListaPagamenti      []Payment `json:"lista_pagamenti"`
	MarcaBollo          int       `json:"marca_bollo"`
	MetodoDescN         string    `json:"metodo_descN"`
	MetodoPagamento     string    `json:"metodo_pagamento"`
	MetodoTitoloN       string    `json:"metodo_titoloN"`
	MostraInfoPagamento bool      `json:"mostra_info_pagamento"`
	NascondiScadenza    bool      `json:"nascondi_scadenza"`
	Nome                string    `json:"nome"`
	Note                string    `json:"note"`
	Numero              string    `json:"numero"`
	OggettoFattura      string    `json:"oggetto_fattura"`
	OggettoInterno      string    `json:"oggetto_interno"`
	Paese               string    `json:"paese"`
	Piva                string    `json:"piva"`
	PrezziIvati         bool      `json:"prezzi_ivati"`
	RitAcconto          int       `json:"rit_acconto"`
	RitAltra            int       `json:"rit_altra"`
	Rivalsa             int       `json:"rivalsa"`
	Valuta              string    `json:"valuta"`
	ValutaCambio        int       `json:"valuta_cambio"`
}

type Article struct {
	ApplicaRaContributi bool    `json:"applica_ra_contributi"`
	Categoria           string  `json:"categoria"`
	CodIva              int     `json:"cod_iva"`
	Codice              string  `json:"codice"`
	Descrizione         string  `json:"descrizione"`
	ID                  string  `json:"id"`
	InDdt               bool    `json:"in_ddt"`
	Magazzino           bool    `json:"magazzino"`
	Nome                string  `json:"nome"`
	Ordine              int     `json:"ordine"`
	PrezzoLordo         float32 `json:"prezzo_lordo"`
	PrezzoNetto         float32 `json:"prezzo_netto"`
	Quantita            int     `json:"quantita"`
	Sconto              int     `json:"sconto"`
	ScontoRosso         int     `json:"sconto_rosso"`
	Tassabile           bool    `json:"tassabile"`
	Um                  string  `json:"um"`
}

type Payment struct {
	DataSaldo    string  `json:"data_saldo"`
	DataScadenza string  `json:"data_scadenza"`
	Importo      float32 `json:"importo"`
	Metodo       string  `json:"metodo"`
}

//******PEOPLE******
type Person struct {
	PA                 bool   `json:"PA"`
	PACodice           string `json:"PA_codice"`
	APIKey             string `json:"api_key"`
	APIUID             string `json:"api_uid"`
	Cf                 string `json:"cf"`
	CodIvaDefault      int    `json:"cod_iva_default"`
	Extra              string `json:"extra"`
	Fax                string `json:"fax"`
	Id                 string `json:"id,omitempty"`
	IndirizzoCap       string `json:"indirizzo_cap"`
	IndirizzoCitta     string `json:"indirizzo_citta"`
	IndirizzoExtra     string `json:"indirizzo_extra"`
	IndirizzoProvincia string `json:"indirizzo_provincia"`
	IndirizzoVia       string `json:"indirizzo_via"`
	Mail               string `json:"mail"`
	Nome               string `json:"nome"`
	Paese              string `json:"paese"`
	PagamentoFineMese  bool   `json:"pagamento_fine_mese"`
	Piva               string `json:"piva"`
	Referente          string `json:"referente"`
	Tel                string `json:"tel"`
	TerminiPagamento   string `json:"termini_pagamento"`
}
