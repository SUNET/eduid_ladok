package ladokmocks

import "github.com/masv3971/goladok3/ladoktypes"

// JSONErrors500 ladok error
var JSONErrors500 = []byte(`
	  {
		"FelUID": "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
		"Felkategori": "commons.fel.kategori.applikationsfel",
		"FelkategoriText": "Generellt fel i applikationen",
		"Meddelande": "java.lang.NullPointerException null",
		"link": []
	  } 
	  `)

// JSONErrorsValideringsFel ladok error
var JSONErrorsValideringsFel = []byte(`
	  {
		"Detaljkod": "commons.domain.uid",
		"DetaljkodText": "Unik identifierare",
		"FelUID": "14c837fd-3a60-11ec-aa00-acd34b504da7",
		"Felgrupp": "commons.fel.grupp.felaktigt_format",
		"FelgruppText": "Felaktigt format",
		"Felkategori": "commons.fel.kategori.valideringsfel",
		"FelkategoriText": "Valideringsfel",
		"Meddelande": "Uid: 6daf0d1e-114f-11ec-95ca-f52940734df",
		"link": []
	  } 
	  `)

// Errors500 return struct containing 500 error
var Errors500 = &ladoktypes.LadokError{
	FelUID:          "c0f52d2c-3a5f-11ec-aa00-acd34b504da7",
	Felkategori:     "commons.fel.kategori.applikationsfel",
	FelkategoriText: "Generellt fel i applikationen",
	Meddelande:      "java.lang.NullPointerException null",
	Link:            []interface{}{},
}
