package ladoktypes

// Student is ladok reply from /studentinformation/student/{studentuid}
type Student struct {
	Avliden                           bool   `json:"Avliden"`
	Efternamn                         string `json:"Efternamn"`
	ExterntUID                        string `json:"ExterntUID"`
	FelVidEtableringExternt           bool   `json:"FelVidEtableringExternt"`
	Fodelsedata                       string `json:"Fodelsedata"`
	FolkbokforingsbevakningTillOchMed string `json:"FolkbokforingsbevakningTillOchMed"`
	Fornamn                           string `json:"Fornamn"`
	KonID                             int    `json:"KonID"`
	LarosateID                        int    `json:"LarosateID"`
	Personnummer                      string `json:"Personnummer"`
	SenastAndradAv                    string `json:"SenastAndradAv"`
	SenastSparad                      string `json:"SenastSparad"`
	UID                               string `json:"Uid"`
	Link                              []Link `json:"link"`
}

// AktivPaLarosate is ladok reply from /studentinformation/student/{uid}/aktivpalarosate
type AktivPaLarosate struct {
	Studentkopplingar []struct {
		LarosateID int    `json:"larosateID"`
		Link       []Link `json:"link"`
		StudentUID string `json:"studentUID"`
	} `json:"Studentkopplingar"`
	Link []Link `json:"link"`
}

// GenderString translate from KonID to the equal string value
func (s *Student) GenderString() string {
	switch s.KonID {
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "n/a"
	}
}
