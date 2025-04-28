package ladokmocks

import (
	"encoding/json"

	"github.com/masv3971/goladok3/ladoktypes"
)

// BehorighetsprofilUID uid for testing behörighetsprofiler
var BehorighetsprofilUID = "3BAD6192-AEBC-4641-9EFD-C740C076E720"

// JSONKataloginformationProfil ladok reply
var JSONKataloginformationProfil = []byte(`{
	"Behorighetsprofiler": [{
		"Benamning": {
			"sv": "Svensk benämning"
		},
		"Dataavgransningar": {
			"LarosateID": 96,
			"Lista": [{
				"DataDimension": "ORGANISATION",
				"DataId": "01234567-1234-abcd-ef01-1234567890abcd",
				"LarosateID": 96,
				"SenastAndradAv": "testMail@example.com",
				"SenastSparad": "2012-01-11T12:45:45",
				"Uid": "11111111-2222-0000-0000-000000000000",
				"link": [{
					"method": "POST",
					"uri": "https://api.mit.ladok.se:443/test",
					"mediaType": "application/vnd.ladok+xml",
					"rel": "http://schemas.ladok.se"
				}]
			}],
			"SenastAndradAv": "testMail@example.com",
			"SenastSparad": "2012-01-11T12:45:45",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		},
		"LarosateID": 96,
		"Rattighetsniva": "rattighetsniva.support",
		"SenastAndradAv": "testMail@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Systemaktiviteter": [{
			"Betafunktion": false,
			"I18nNyckel": "systemaktivitet.resultatrapportering",
			"Id": 2147483647,
			"KlarForProduktion": false,
			"LarosateID": 96,
			"Rattighetsniva": "rattighetsniva.support",
			"link": [{
				"method": "POST",
				"uri": "https://api.mit.ladok.se:443/test",
				"mediaType": "application/vnd.ladok+xml",
				"rel": "http://schemas.ladok.se"
			}]
		}],
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		}]
	}],
	"LarosateID": 96,
	"SenastAndradAv": "testMail@example.com",
	"SenastSparad": "2012-01-11T12:45:45",
	"Uid": "11111111-2222-0000-0000-000000000000",
	"link": [{
		"method": "POST",
		"uri": "https://api.mit.ladok.se:443/test",
		"mediaType": "application/vnd.ladok+xml",
		"rel": "http://schemas.ladok.se"
	}]
}`)

// MockKataloginformationProfil return mock
func MockKataloginformationProfil() *ladoktypes.KataloginformationBehorighetsprofil {
	s := &ladoktypes.KataloginformationBehorighetsprofil{}
	json.Unmarshal(JSONKataloginformationProfil, s)
	return s
}

// JSONKataloginformationBehorighetsprofilNoPermissions ladok reply
var JSONKataloginformationBehorighetsprofilNoPermissions = []byte(`
{
	"Benamning": {
	  "sv": "Intergration-Sunet",
	  "en": "Intergration-Sunet"
	},
	"Dataavgransningar": {
	  "Lista": [],
	  "link": []
	},
	"LarosateID": 27,
	"Rattighetsniva": "rattighetsniva.las",
	"Systemaktiviteter": [],
	"Uid": "09E52B69-5D50-4A62-B65C-636BCA68FAE5",
	"link": [
	  {
		"method": "GET",
		"uri": "https://api.integrationstest.ladok.se:443/kataloginformation/behorighetsprofil/09E52B69-5D50-4A62-B65C-636BCA68FAE5",
		"mediaType": "application/vnd.ladok+xml,application/vnd.ladok-kataloginformation+xml,application/vnd.ladok-kataloginformation+json",
		"rel": "self"
	  }
	]
  }
`)

// JSONKataloginformationBehorighetsprofil ladok reply
var JSONKataloginformationBehorighetsprofil = []byte(`
{
	"Benamning": {
	  "sv": "Intergration-Sunet",
	  "en": "Intergration-Sunet"
	},
	"Dataavgransningar": {
	  "Lista": [],
	  "link": []
	},
	"LarosateID": 27,
	"Rattighetsniva": "rattighetsniva.las",
	"Systemaktiviteter": [
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.uppfoljning.feeds",
		"Id": 90019,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.studiedeltagande.las",
		"Id": 51001,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.studentinformation.lasa",
		"Id": 61001,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.kataloginformation.las",
		"Id": 11004,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  },
	  {
		"Betafunktion": false,
		"I18nNyckel": "systemaktivitet.extintegration.lasa",
		"Id": 860131,
		"KlarForProduktion": true,
		"Rattighetsniva": "rattighetsniva.las",
		"link": []
	  }
	],
	"Uid": "09E52B69-5D50-4A62-B65C-636BCA68FAE5",
	"link": [
	  {
		"method": "GET",
		"uri": "https://api.integrationstest.ladok.se:443/kataloginformation/behorighetsprofil/09E52B69-5D50-4A62-B65C-636BCA68FAE5",
		"mediaType": "application/vnd.ladok+xml,application/vnd.ladok-kataloginformation+xml,application/vnd.ladok-kataloginformation+json",
		"rel": "self"
	  }
	]
  }
`)

//MockKataloginformationBehorighetsprofil return mock ladok KataloginformationBehorighetsprofil
func MockKataloginformationBehorighetsprofil() *ladoktypes.KataloginformationBehorighetsprofil {
	s := &ladoktypes.KataloginformationBehorighetsprofil{}
	json.Unmarshal(JSONKataloginformationBehorighetsprofil, s)
	return s
}

// JSONKataloginformationAutentiserad ladok reply
var JSONKataloginformationAutentiserad = []byte(`
{
		"Anvandarnamn": "mail@school.se",
		"Efternamn": "testEfternamn",
		"Fornamn": "testFornamn",
		"LarosateID": 96,
		"SenastAndradAv": "name@ladok3.ladok.se",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [ {
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		} ]
	  }
`)

// MockKataloginformationAutentiserad return ladok mock
func MockKataloginformationAutentiserad() *ladoktypes.KataloginformationAnvandareAutentiserad {
	s := &ladoktypes.KataloginformationAnvandareAutentiserad{}
	json.Unmarshal(JSONKataloginformationAutentiserad, s)
	return s
}

// JSONKataloginformationEgna ladok reply
var JSONKataloginformationEgna = []byte(`{
		"Anvandarbehorighet": [{
		  "AnvandareRef": {
			"Anvandarnamn": "testEppn@example.com",
			"Efternamn": "TestEfternamn",
			"Fornamn": "TestFornamn",
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BehorighetsprofilRef":{
			"Benamning":[{
			  "Sprakkod":"sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "BestalldTidpunkt": "2013-10-14T12:45:45",
		  "LarosateID": 96,
		  "OrganisationRef": {
			"Benamning": [{
			  "Sprakkod": "sv",
			  "Text": "Svenska",
			  "link": [ ]
			}, {
			  "Sprakkod": "en",
			  "Text": "English",
			  "link": [ ]
			}],
			"Uid": "11111111-2222-0000-0000-000000000000",
			"link": {
			  "method": "POST",
			  "uri": "https://api.mit.ladok.se:443/test",
			  "mediaType": "application/vnd.ladok+xml",
			  "rel": "http://schemas.ladok.se"
			}
		  },
		  "SenastAndradAv": "testEppn@example.com",
		  "SenastSparad": "2012-01-11T12:45:45",
		  "Status": "AKTIV",
		  "Uid": "11111111-2222-0000-0000-000000000000",
		  "link": [{
			"method": "POST",
			"uri": "https://api.mit.ladok.se:443/test",
			"mediaType": "application/vnd.ladok+xml",
			"rel": "http://schemas.ladok.se"
		  }]
		}],
		"LarosateID": 96,
		"SenastAndradAv": "testEppn@example.com",
		"SenastSparad": "2012-01-11T12:45:45",
		"Uid": "11111111-2222-0000-0000-000000000000",
		"link": [{
		  "method": "POST",
		  "uri": "https://api.mit.ladok.se:443/test",
		  "mediaType": "application/vnd.ladok+xml",
		  "rel": "http://schemas.ladok.se"
		}]
	  }`)

// MockKataloginformationEgna return mock
func MockKataloginformationEgna() *ladoktypes.KataloginformationAnvandarbehorighetEgna {
	s := &ladoktypes.KataloginformationAnvandarbehorighetEgna{}
	json.Unmarshal(JSONKataloginformationEgna, s)
	return s
}

var JSONKataloginformationGrunddataLarosateinformation = []byte(`
{
	"LarosateID" : 96,
	"Larosatesinformation" : [ {
	  "Benamning" : {
		"sv" : "Svensk benämning"
	  },
	  "Beskrivning" : {
		"sv" : "Svensk benämning"
	  },
	  "EpostadressForAdmingranssnitt" : "info@hsk.se",
	  "EpostadressForStudentgranssnitt" : "info-student@hsk.se",
	  "Giltighetsperiod" : {
		"LarosateID" : 96,
		"Slutdatum" : "2014-06-25",
		"Startdatum" : "2014-01-25",
		"link" : [ {
		  "method" : "POST",
		  "uri" : "https://api.mit.ladok.se:443/test",
		  "mediaType" : "application/vnd.ladok+xml",
		  "rel" : "http://schemas.ladok.se"
		} ]
	  },
	  "ID" : "123",
	  "Kod" : "KOD",
	  "LankTillWebbplats" : {
		"Lanktext" : "10100",
		"Url" : "https://hsk.se"
	  },
	  "LankTillWebbplatsEngelskSida" : {
		"Lanktext" : "10100",
		"Url" : "https://hsk.se"
	  },
	  "LarosateID" : 96,
	  "OrtID" : 1,
	  "Postadress" : {
		"Postnummer" : "10100",
		"Postort" : "Storstan",
		"Utdelningsadress" : "Storgatan 1"
	  },
	  "Telefonnummer" : "010-778899",
	  "link" : [ {
		"method" : "POST",
		"uri" : "https://api.mit.ladok.se:443/test",
		"mediaType" : "application/vnd.ladok+xml",
		"rel" : "http://schemas.ladok.se"
	  } ]
	} ],
	"link" : [ {
	  "method" : "POST",
	  "uri" : "https://api.mit.ladok.se:443/test",
	  "mediaType" : "application/vnd.ladok+xml",
	  "rel" : "http://schemas.ladok.se"
	} ]
  }
`)
