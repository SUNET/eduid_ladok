package goeduidiam

import (
	"fmt"
	"time"
)

// Meta type
type Meta struct {
	Location     string    `json:"location"`
	LastModified time.Time `json:"lastModified"`
	ResourceType string    `json:"resourceType"`
	Created      time.Time `json:"created"`
	Version      string    `json:"version"`
}

// Member type
type Member struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

// HTTPSScimEduidSeSchemaNutidGroupV1  type
type HTTPSScimEduidSeSchemaNutidGroupV1 struct {
	Data struct {
	} `json:"data"`
}

// HTTPSScimEduidSeSchemaNutidUserV1 type
type HTTPSScimEduidSeSchemaNutidUserV1 struct {
	Profiles       Profiles        `json:"profiles"`
	LinkedAccounts []LinkedAccount `json:"linked_accounts"`
}

// HTTPSScimEduidSeSchemaNutidEventV1 type
type HTTPSScimEduidSeSchemaNutidEventV1 struct {
	Resource Resource `json:"resource"`
	Level    string   `json:"level"`
	Data     struct {
	} `json:"data"`
	ExpiresAt time.Time `json:"expiresAt"`
	Timestamp time.Time `json:"timestamp"`
	Source    string    `json:"source"`
}

// Profiles type
type Profiles struct {
	AdditionalProp1 Prop `json:"additionalProp1"`
	AdditionalProp2 Prop `json:"additionalProp2"`
	AdditionalProp3 Prop `json:"additionalProp3"`
}

// LinkedAccount type
type LinkedAccount struct {
	Issuer     string `json:"issuer"`
	Value      string `json:"value"`
	Parameters struct {
	} `json:"parameters"`
}

// Prop type
type Prop struct {
	Attributes struct {
	} `json:"attributes"`
	Data struct {
	} `json:"data"`
}

// SearchRequest type
type SearchRequest struct {
	Schemas    []string `json:"schemas"`
	Filter     string   `json:"filter"`
	StartIndex int      `json:"startIndex"`
	Count      int      `json:"count"`
	Attributes []string `json:"attributes"`
}

// SearchReply type
type SearchReply struct {
	Schemas   []string `json:"schemas"`
	Resources []struct {
	} `json:"Resources"`
	TotalResults int `json:"totalResults"`
}

// GroupsRequest type
type GroupsRequest struct {
	Schemas                            []string                           `validate:"required" json:"schemas"`
	ExternalID                         string                             `validate:"required" json:"externalId"`
	DisplayName                        string                             `validate:"required" json:"displayName"`
	Members                            []Member                           `validate:"required" json:"members"`
	HTTPSScimEduidSeSchemaNutidGroupV1 HTTPSScimEduidSeSchemaNutidGroupV1 `validate:"required" json:"https://scim.eduid.se/schema/nutid/group/v1"`
}

// GroupsReply type
type GroupsReply struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	GroupsRequest
}

// EventsReply type
type EventsReply struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	EventsRequest
}

// EventsRequest type
type EventsRequest struct {
	Schemas                            []string                           `json:"schemas"`
	ExternalID                         string                             `json:"externalId"`
	HTTPSScimEduidSeSchemaNutidEventV1 HTTPSScimEduidSeSchemaNutidEventV1 `json:"https://scim.eduid.se/schema/nutid/event/v1"`
}

// LoginRequest type
type LoginRequest struct {
	DataOwner string `json:"data_owner"`
}

// Resource type
type Resource struct {
	ResourceType string    `json:"resourceType"`
	ID           string    `json:"id"`
	LastModified time.Time `json:"lastModified"`
	Version      string    `json:"version"`
	ExternalID   string    `json:"externalId"`
	Location     string    `json:"location"`
}

// UsersReply type
type UsersReply struct {
	ID   string `json:"id"`
	Meta Meta   `json:"meta"`
	UsersRequest
}

// UsersRequest type
type UsersRequest struct {
	Schemas                           []string                          `json:"schemas"`
	ExternalID                        string                            `json:"externalId"`
	Name                              Name                              `json:"name"`
	Emails                            []ContactInfo                     `json:"emails"`
	PhoneNumbers                      []ContactInfo                     `json:"phoneNumbers"`
	PreferredLanguage                 string                            `json:"preferredLanguage"`
	Groups                            []Group                           `json:"groups"`
	HTTPSScimEduidSeSchemaNutidUserV1 HTTPSScimEduidSeSchemaNutidUserV1 `json:"https://scim.eduid.se/schema/nutid/user/v1"`
}

// InvitesRequest type
type InvitesRequest struct{}

// Name type
type Name struct {
	FamilyName      string `json:"familyName"`
	GivenName       string `json:"givenName"`
	Formatted       string `json:"formatted"`
	MiddleName      string `json:"middleName"`
	HonorificPrefix string `json:"honorificPrefix"`
	HonorificSuffix string `json:"honorificSuffix"`
}

// ContactInfo type
type ContactInfo struct {
	Value   string `json:"value"`
	Display string `json:"display"`
	Type    string `json:"type"`
	Primary bool   `json:"primary"`
}

// Group type
type Group struct {
	Value   string `json:"value"`
	Ref     string `json:"$ref"`
	Display string `json:"display"`
}

// HealthyReply type
type HealthyReply struct {
	Status   string `json:"status"`
	Hostname string `json:"hostname"`
	Reason   string `json:"reason"`
}

// EmptyStruct type
type EmptyStruct struct{}

// Errors is a general error reply
type Errors struct {
	Detail []struct {
		Loc  []string `json:"loc"`
		Msg  string   `json:"msg"`
		Type string   `json:"type"`
	} `json:"detail"`
}

// Error interface
type Error interface {
	Error() string
}

func (e *Errors) Error() string {
	return fmt.Sprintf("error: %v", e.Detail)
}

// RequestCFG general type for making requests
//type RequestCFG struct {
//	ScimID  string
//	Search  SearchRequest
//	Users   UsersRequest
//	Groups  GroupsRequest
//	Invites InvitesRequest
//	Events  EventsRequest
//}
