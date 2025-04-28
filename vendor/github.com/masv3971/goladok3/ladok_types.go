package goladok3

// Permissions is a simplify permissions object
type Permissions map[int64]string

type (
	// FeedID ladok id is an int
	FeedID string
)

type serviceTypes struct {
	client       *Client
	service      string
	acceptHeader string
}

var (
	// ContentTypeStudiedeltagandeJSON server response content type
	ContentTypeStudiedeltagandeJSON = "application/vnd.ladok-studiedeltagande+json;charset=UTF-8"
	// ContentTypeKataloginformationJSON server response content type
	ContentTypeKataloginformationJSON = "application/vnd.ladok-kataloginformation+json;charset=UTF-8"
	// ContentTypeStudentinformationJSON server response content type
	ContentTypeStudentinformationJSON = "application/vnd.ladok-studentinformation+json;charset=UTF-8"
	// ContentTypeAtomXML server response content type
	ContentTypeAtomXML = "application/atom+xml;charset=UTF-8"

	// TypeStudentinformation type
	TypeStudentinformation = "studentinformation"
	//  = "externstudentevent"
)
