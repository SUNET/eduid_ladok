package ladokmocks

import (
	"encoding/json"

	"github.com/masv3971/goladok3/ladoktypes"
)

// JSONStudentinformationStudent mock ladok reply
var JSONStudentinformationStudent = []byte(`
{
	"Avliden": false,
	"Efternamn": "TestEfternamn",
	"ExterntUID": "11111111-2222-0000-0000-000000000000",
	"FelVidEtableringExternt": false,
	"Fodelsedata": "1996-11-05",
	"FolkbokforingsbevakningTillOchMed": "2020-03-13",
	"Fornamn": "TestFornamn",
	"KonID": 1,
	"LarosateID": 96,
	"Personnummer": "199611052383",
	"SenastAndradAv": "testEppn@ladok3.ladok.se",
	"SenastSparad": "2012-01-11T12:45:45",
	"Uid": "11111111-2222-0000-0000-000000000000"
}
`)

// JSONAktivPaLarosate mock ladok reply
var JSONAktivPaLarosate = []byte(` 
		{
			"Studentkopplingar": [
			  {
				"LarosateID": 27,
				"link": [],
				"studentUID": "339A47C0-426D-4012-B83A-6427E9587352"
			  },
			  {
				"LarosateID": 39,
				"link": [],
				"studentUID": "339A47C0-426D-4012-B83A-6427E9587352"
			  }
			],
			"link": []
		  }
		`)

// MockStudentinformationStudent return mock
func MockStudentinformationStudent() *ladoktypes.Student {
	s := &ladoktypes.Student{}
	json.Unmarshal(JSONStudentinformationStudent, s)
	return s
}

// StudentJSON return JSON object of a student
func StudentJSON(studentData StudentData) []byte {
	s := &ladoktypes.Student{}

	if err := json.Unmarshal(JSONStudentinformationStudent, s); err != nil {
		return nil
	}

	s.Personnummer = studentData.Personnummer
	s.UID = studentData.StudentUID
	s.ExterntUID = studentData.ExterntUID
	s.Fodelsedata = studentData.DateOfBirth

	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	return b
}

// StudentData keeps a mock record of a student
type StudentData struct {
	Personnummer string
	StudentUID   string
	ExterntUID   string
	DateOfBirth  string
}

// Students mocks a student with personnummer, studentUID, externtUID and birth date
var Students = map[int]StudentData{
	0: {
		Personnummer: "198601049995",
		StudentUID:   "44889B47-C78B-440B-BA98-A16C2C27BE7C",
		ExterntUID:   "72A06BD3-A7A0-44A3-A3AA-51B9E3208015",
		DateOfBirth:  "1986-01-04",
	},
	1: {
		Personnummer: "198602179882",
		StudentUID:   "339A47C0-426D-4012-B83A-6427E9587352",
		ExterntUID:   "72460B4B-8F15-442C-A464-0743BDFB1429",
		DateOfBirth:  "1986-02-17",
	},
	2: {
		Personnummer: "198603139885",
		StudentUID:   "82E208E7-FCDC-407E-9EE4-D2708CD609CC",
		ExterntUID:   "0BDF38F5-30A3-4F1E-B851-D538E8A83FBB",
		DateOfBirth:  "1986-03-13",
	},
	3: {
		Personnummer: "198603249999",
		StudentUID:   "9711A40B-2C40-414D-ACC9-FAC4C4D35C50",
		ExterntUID:   "9C0E5285-41E1-4190-BA49-6DC134A014D4",
		DateOfBirth:  "1986-03-24",
	},
}
