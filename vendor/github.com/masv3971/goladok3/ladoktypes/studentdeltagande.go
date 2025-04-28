package ladoktypes

type TillfallesdeltagandePagaendeStudent struct {
	LarosateID            int `json:"LarosateID"`
	Tillfallesdeltaganden []struct {
		Aterbud                           bool   `json:"Aterbud"`
		Avklarad                          bool   `json:"Avklarad"`
		ForegaendeTillfallesdeltagandeUID string `json:"ForegaendeTillfallesdeltagandeUID"`
		Godkannandedatum                  string `json:"Godkannandedatum"`
		HarPagandeUppehall                bool   `json:"HarPagandeUppehall"`
		HarTillgodoraknande               bool   `json:"HarTillgodoraknande"`
		LarosateID                        int    `json:"LarosateID"`
		Nuvarande                         bool   `json:"Nuvarande"`
		Paborjad                          bool   `json:"Paborjad"`
		Perioddeltaganden                 []struct {
			Anpassat          bool   `json:"Anpassat"`
			AterkalladDatum   string `json:"AterkalladDatum"`
			LarosateID        int    `json:"LarosateID"`
			Omfattningsvarde  string `json:"Omfattningsvarde"`
			Periodbenamningar struct {
				Sv string `json:"sv"`
			} `json:"Periodbenamningar"`
			PerioddeltagandeUID string `json:"PerioddeltagandeUID"`
			Periodindex         int    `json:"Periodindex"`
			Registreringsperiod struct {
				LarosateID int    `json:"LarosateID"`
				Slutdatum  string `json:"Slutdatum"`
				Startdatum string `json:"Startdatum"`
				Link       []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Registreringsperiod"`
			SenastAndradAv        string `json:"SenastAndradAv"`
			SenastSparad          string `json:"SenastSparad"`
			Tillstand             string `json:"Tillstand"`
			UID                   string `json:"Uid"`
			UtbildningsPeriodFran string `json:"UtbildningsPeriodFran"`
			UtbildningsPeriodTill string `json:"UtbildningsPeriodTill"`
			Link                  []struct {
				Method    string `json:"method"`
				URI       string `json:"uri"`
				MediaType string `json:"mediaType"`
				Rel       string `json:"rel"`
			} `json:"link"`
		} `json:"Perioddeltaganden"`
		Registreringsperiod struct {
			LarosateID int    `json:"LarosateID"`
			Slutdatum  string `json:"Slutdatum"`
			Startdatum string `json:"Startdatum"`
			Link       []struct {
				Method    string `json:"method"`
				URI       string `json:"uri"`
				MediaType string `json:"mediaType"`
				Rel       string `json:"rel"`
			} `json:"link"`
		} `json:"Registreringsperiod"`
		SenareDel                            bool   `json:"SenareDel"`
		SenastAndradAv                       string `json:"SenastAndradAv"`
		SenastSparad                         string `json:"SenastSparad"`
		Sparrad                              bool   `json:"Sparrad"`
		Studiestrukturreferens               string `json:"Studiestrukturreferens"`
		SummeradGodkandOmfattning            string `json:"SummeradGodkandOmfattning"`
		SummeradHeltTillgodoraknadOmfattning string `json:"SummeradHeltTillgodoraknadOmfattning"`
		SummeradTillgodoraknadOmfattning     string `json:"SummeradTillgodoraknadOmfattning"`
		Tillganglighet                       struct {
			RegistreringEjTillgangligtFranOchMed string `json:"RegistreringEjTillgangligtFranOchMed"`
			RegistreringTillgangligtFranOchMed   string `json:"RegistreringTillgangligtFranOchMed"`
			ValEjTillgangligtFranOchMed          string `json:"ValEjTillgangligtFranOchMed"`
			ValTillgangligtFranOchMed            string `json:"ValTillgangligtFranOchMed"`
		} `json:"Tillganglighet"`
		TillstandKurs struct {
			Sammanfattat          string `json:"Sammanfattat"`
			Sammanfattattillstand string `json:"Sammanfattattillstand"`
			Tillfallesdeltagande  string `json:"Tillfallesdeltagande"`
			Utbildning            string `json:"Utbildning"`
		} `json:"TillstandKurs"`
		TillstandKurspaketering struct {
			Sammanfattat         string `json:"Sammanfattat"`
			Tillfallesdeltagande string `json:"Tillfallesdeltagande"`
			Utbildning           string `json:"Utbildning"`
		} `json:"TillstandKurspaketering"`
		UID                    string `json:"Uid"`
		Utbildningsinformation struct {
			AntalPerioder int    `json:"AntalPerioder"`
			AvsesLedaTill string `json:"AvsesLedaTill"`
			Benamning     struct {
				Sv string `json:"sv"`
			} `json:"Benamning"`
			Enhetskod             string `json:"Enhetskod"`
			FinansieringsformID   int    `json:"FinansieringsformID"`
			Installt              bool   `json:"Installt"`
			LarosateID            int    `json:"LarosateID"`
			Omfattningsvarde      string `json:"Omfattningsvarde"`
			Organisationbenamning struct {
				Sv string `json:"sv"`
			} `json:"Organisationbenamning"`
			Organisationskod string `json:"Organisationskod"`
			PeriodIOrdning   string `json:"PeriodIOrdning"`
			Perioder         []struct {
				Index            int    `json:"Index"`
				LarosateID       int    `json:"LarosateID"`
				Omfattningsvarde string `json:"Omfattningsvarde"`
				SenastAndradAv   string `json:"SenastAndradAv"`
				SenastSparad     string `json:"SenastSparad"`
				Slutdatum        string `json:"Slutdatum"`
				Startdatum       string `json:"Startdatum"`
				UID              string `json:"Uid"`
				Link             []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Perioder"`
			SenareDel                    bool   `json:"SenareDel"`
			SpecificeratOmfattningsvarde string `json:"SpecificeratOmfattningsvarde"`
			Studielokalisering           struct {
				Sv string `json:"sv"`
			} `json:"Studielokalisering"`
			Studielokaliseringrepresentation struct {
				Benamningar struct {
					Sv string `json:"sv"`
				} `json:"Benamningar"`
				Giltighetsperiod struct {
					LarosateID int    `json:"LarosateID"`
					Slutdatum  string `json:"Slutdatum"`
					Startdatum string `json:"Startdatum"`
					Link       []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"Giltighetsperiod"`
				ID         int    `json:"ID"`
				Kod        string `json:"Kod"`
				LarosateID int    `json:"LarosateID"`
				Link       []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Studielokaliseringrepresentation"`
			StudieordningID string `json:"StudieordningID"`
			Studieperiod    struct {
				LarosateID int    `json:"LarosateID"`
				Slutdatum  string `json:"Slutdatum"`
				Startdatum string `json:"Startdatum"`
				Link       []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Studieperiod"`
			Studietakt struct {
				Benamning struct {
					Sv string `json:"sv"`
				} `json:"Benamning"`
				Takt int64 `json:"Takt"`
			} `json:"Studietakt"`
			Undervisningsform struct {
				Benamningar struct {
					Sv string `json:"sv"`
				} `json:"Benamningar"`
				Giltighetsperiod struct {
					LarosateID int    `json:"LarosateID"`
					Slutdatum  string `json:"Slutdatum"`
					Startdatum string `json:"Startdatum"`
					Link       []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"Giltighetsperiod"`
				ID         int `json:"ID"`
				LarosateID int `json:"LarosateID"`
				Link       []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Undervisningsform"`
			UtbildningUID                 string `json:"UtbildningUID"`
			UtbildningensOmfattningsvarde string `json:"UtbildningensOmfattningsvarde"`
			UtbildningsinstansUID         string `json:"UtbildningsinstansUID"`
			Utbildningskod                string `json:"Utbildningskod"`
			UtbildningssamarbeteID        int    `json:"UtbildningssamarbeteID"`
			UtbildningstillfalleUID       string `json:"UtbildningstillfalleUID"`
			Utbildningstillfalleskod      string `json:"Utbildningstillfalleskod"`
			Utbildningstillfallestyp      struct {
				Benamningar struct {
					Sv string `json:"sv"`
				} `json:"Benamningar"`
				Giltighetsperiod struct {
					LarosateID int    `json:"LarosateID"`
					Slutdatum  string `json:"Slutdatum"`
					Startdatum string `json:"Startdatum"`
					Link       []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"Giltighetsperiod"`
				Grundtyp                   string `json:"Grundtyp"`
				ID                         int    `json:"ID"`
				Kod                        string `json:"Kod"`
				LarosateID                 int    `json:"LarosateID"`
				RegelverkForUtbildningstyp struct {
					LarosateID  int `json:"LarosateID"`
					Regelvarden []struct {
						LarosateID     int    `json:"LarosateID"`
						Regelnamn      string `json:"Regelnamn"`
						SenastAndradAv string `json:"SenastAndradAv"`
						SenastSparad   string `json:"SenastSparad"`
						UID            string `json:"Uid"`
						Varde          string `json:"Varde"`
						Link           []struct {
							Method    string `json:"method"`
							URI       string `json:"uri"`
							MediaType string `json:"mediaType"`
							Rel       string `json:"rel"`
						} `json:"link"`
					} `json:"Regelvarden"`
					SenastAndradAv string `json:"SenastAndradAv"`
					SenastSparad   string `json:"SenastSparad"`
					UID            string `json:"Uid"`
					Link           []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"RegelverkForUtbildningstyp"`
				Sjalvstandig bool `json:"Sjalvstandig"`
				Link         []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Utbildningstillfallestyp"`
			Utbildningstyp struct {
				Benamningar struct {
					Sv string `json:"sv"`
				} `json:"Benamningar"`
				Giltighetsperiod struct {
					LarosateID int    `json:"LarosateID"`
					Slutdatum  string `json:"Slutdatum"`
					Startdatum string `json:"Startdatum"`
					Link       []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"Giltighetsperiod"`
				Grundtyp                   string `json:"Grundtyp"`
				ID                         int    `json:"ID"`
				Kod                        string `json:"Kod"`
				LarosateID                 int    `json:"LarosateID"`
				RegelverkForUtbildningstyp struct {
					LarosateID  int `json:"LarosateID"`
					Regelvarden []struct {
						LarosateID     int    `json:"LarosateID"`
						Regelnamn      string `json:"Regelnamn"`
						SenastAndradAv string `json:"SenastAndradAv"`
						SenastSparad   string `json:"SenastSparad"`
						UID            string `json:"Uid"`
						Varde          string `json:"Varde"`
						Link           []struct {
							Method    string `json:"method"`
							URI       string `json:"uri"`
							MediaType string `json:"mediaType"`
							Rel       string `json:"rel"`
						} `json:"link"`
					} `json:"Regelvarden"`
					SenastAndradAv string `json:"SenastAndradAv"`
					SenastSparad   string `json:"SenastSparad"`
					UID            string `json:"Uid"`
					Link           []struct {
						Method    string `json:"method"`
						URI       string `json:"uri"`
						MediaType string `json:"mediaType"`
						Rel       string `json:"rel"`
					} `json:"link"`
				} `json:"RegelverkForUtbildningstyp"`
				Sjalvstandig bool `json:"Sjalvstandig"`
				Link         []struct {
					Method    string `json:"method"`
					URI       string `json:"uri"`
					MediaType string `json:"mediaType"`
					Rel       string `json:"rel"`
				} `json:"link"`
			} `json:"Utbildningstyp"`
			Utbildningsversion int `json:"Utbildningsversion"`
			Link               []struct {
				Method    string `json:"method"`
				URI       string `json:"uri"`
				MediaType string `json:"mediaType"`
				Rel       string `json:"rel"`
			} `json:"link"`
			OrganisationUID string `json:"organisationUID"`
		} `json:"Utbildningsinformation"`
		YtterstaPaketeringen bool `json:"YtterstaPaketeringen"`
		Link                 []struct {
			Method    string `json:"method"`
			URI       string `json:"uri"`
			MediaType string `json:"mediaType"`
			Rel       string `json:"rel"`
		} `json:"link"`
	} `json:"Tillfallesdeltaganden"`
	Link []struct {
		Method    string `json:"method"`
		URI       string `json:"uri"`
		MediaType string `json:"mediaType"`
		Rel       string `json:"rel"`
	} `json:"link"`
}
