package ladoktypes

import (
	"errors"
	"fmt"
)

var (
	// ErrNoValidContentType if no valid content-type is found
	ErrNoValidContentType = errors.New("No valid content-type found")
	// ErrNoEnvFound if no valid environment is found in certificate (ou)
	ErrNoEnvFound = errors.New("No valid ladok environment (OU) found in certificate")
	// ErrNotSufficientPermissions if not all provided permissions are met
	ErrNotSufficientPermissions = PermissionErrors{{Msg: "No permissions found in ladok"}}
	// ErrNoPermissionProvided when input Permission is empty
	ErrNoPermissionProvided = PermissionErrors{{Msg: "No permissions provided"}}
)

// FromLadokError returns error by Ladok
type LadokError struct {
	Detaljkod       string        `json:"Detaljkod"`
	DetaljkodText   string        `json:"DetaljkodText"`
	FelUID          string        `json:"FelUID"`
	Felgrupp        string        `json:"Felgrupp"`
	FelgruppText    string        `json:"FelgruppText"`
	Felkategori     string        `json:"Felkategori"`
	FelkategoriText string        `json:"FelkategoriText"`
	Meddelande      string        `json:"Meddelande"`
	Link            []interface{} `json:"link"`
}

func NewLadokError() LadokError {
	return LadokError{
		Detaljkod:       "",
		DetaljkodText:   "",
		FelUID:          "",
		Felgrupp:        "",
		FelgruppText:    "",
		Felkategori:     "",
		FelkategoriText: "",
		Meddelande:      "",
		Link:            []interface{}{},
	}
}

func (f *LadokError) Error() string {
	if f == nil {
		return ""
	}

	return fmt.Sprintf("felUID: %q, detaljkod_text: %q", f.FelUID, f.DetaljkodText)
}

type PermissionError struct {
	Msg                 string `json:"msg"`
	MissingPermissionID int64  `json:"missing_permission_id"`
	PermissionLevel     string `json:"permission_level"`
}

func (p PermissionError) Error() string {
	return fmt.Sprintf("%s %d %s", p.Msg, p.MissingPermissionID, p.PermissionLevel)
}

type PermissionErrors []PermissionError

func (p PermissionErrors) Error() string {
	if len(p) == 0 {
		return ""
	}

	t := ""
	for _, errors := range p {
		t += fmt.Sprintf("%s,%d,%s\n", errors.Msg, errors.MissingPermissionID, errors.PermissionLevel)

	}
	return t
}
