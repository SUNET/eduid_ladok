package ladoktypes

// Link is a general ladok link structure
type Link struct {
	Method    string `json:"method"`
	URI       string `json:"uri"`
	MediaType string `json:"mediaType"`
	Rel       string `json:"rel"`
}

// Benamning is a general ladok Benamning structure
type Benamning struct {
	Sprakkod string     `json:"Sprakkod"`
	Text     string     `json:"Text"`
	Link     []struct{} `json:"link"`
}

const (
	// EnvIntTestAPI ladok integration environment
	EnvIntTestAPI = "Int-test-API"
	// EnvProdAPI ladok production environment
	EnvProdAPI = "Prod-API"
	// EnvTestAPI ladok test environment
	EnvTestAPI = "Test-API"
)
