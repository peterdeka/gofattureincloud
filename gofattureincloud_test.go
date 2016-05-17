package gofattureincloud

import (
	"env"
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

var tp Person
var creds Credentials

//tests setup
func setup() {
	rand.Seed(time.Now().Unix())
	creds = Credentials{
		APIKey: os.Getenv("FATTURE_KEY"),
		APIUID: os.Getenv("FATTURE_UID"),
	}
	tp = Person{
		Cf:                 randN(16),
		Piva:               randN(13),
		Mail:               randS(5) + "@" + randS(5) + ".com",
		Nome:               "APITEST " + randS(5) + " " + randS(3),
		IndirizzoCap:       "20100",
		IndirizzoCitta:     "Milano",
		IndirizzoProvincia: "MI",
		IndirizzoVia:       randS(18),
		Paese:              "Italia",
		TerminiPagamento:   "0",
	}
}

func TestPeopleService(t *testing.T) {
	require := require.New(t)
	ps := PeopleService{Creds: creds}
	err, newId := ps.SavePerson(Clienti, &tp)
	require.Nil(err)
	require.Equal(tp.Id, *newId)
	lc, lf, err := ps.List(Clienti, &PeopleListReqParams{ID: *newId, Pagina: 1})
	require.Nil(err)
	require.Equal(1, len(lc))
	require.Equal(0, len(lf))
	require.Equal(lc[0].Nome, tp.Nome)
	err = ps.DeletePerson(Clienti, &PeopleDeleteReqParams{Id: *newId})
	require.Nil(err)
}

func TestDocumentsService(t *testing.T) {
	require := require.New(t)
	ps := PeopleService{Creds: creds}
	tp.Id = ""
	err, newId := ps.SavePerson(Clienti, &tp)
	require.Nil(err)
	fmt.Println(tp.Id)

	ta := []Article{
		Article{
			Nome:        "Test article",
			PrezzoNetto: 99,
			Tassabile:   true,
			Quantita:    "1",
		},
	}
	td := Document{
		AnnoCompetenza:     2016,
		Cf:                 tp.Cf,
		Data:               "02/05/2016",
		IDCliente:          tp.Id,
		Nome:               tp.Nome,
		IndirizzoCap:       tp.IndirizzoCap,
		IndirizzoCitta:     tp.IndirizzoCitta,
		IndirizzoProvincia: tp.IndirizzoProvincia,
		IndirizzoVia:       tp.IndirizzoVia,
		Lingua:             "it",
		Paese:              tp.Paese,
		Piva:               tp.Piva,
		Numero:             "6-TESTAPI",
		Valuta:             "EUR",
		ListaArticoli:      ta,
		ListaPagamenti: []Payment{
			Payment{
				DataSaldo:    "02/05/2016",
				DataScadenza: "02/05/2016",
				Importo:      99 + 99*0.22,
				Metodo:       "Bonifico",
			},
		},
	}
	ds := DocumentsService{Creds: creds}
	err, newId = ds.SaveDocument(Fatture, &td)
	require.Nil(err)
	require.Equal(td.Id, *newId)
	d, err := ds.GetDocument(Fatture, &DocumentDetailsReqParams{Id: *newId})
	require.Nil(err)
	require.Equal(d.ListaPagamenti[0].Importo, td.ListaPagamenti[0].Importo)
	err = ds.DeleteDocument(Fatture, &DocumentDetailsReqParams{Id: td.Id})
	require.Nil(err)
	err = ps.DeletePerson(Clienti, &PeopleDeleteReqParams{Id: *newId})
	require.Nil(err)
}

//util funcs
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "0123456789"

func randS(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randN(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = digitBytes[rand.Intn(len(digitBytes))]
	}
	return string(b)
}
