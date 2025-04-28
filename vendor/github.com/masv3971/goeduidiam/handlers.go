package goeduidiam

import (
	"context"
	"fmt"
	"net/http"
)

// UsersService holds the object for Users
type UsersService struct {
	client *Client
	path   string
}

// GetUsersRequest holds the request for user
type GetUsersRequest struct {
	ScimID string `validate:"required" json:"scim_id"`
}

// Get gets users from eduidIAM
func (s *UsersService) Get(ctx context.Context, req *GetUsersRequest) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// PutUsersRequest holds request for user
type PutUsersRequest struct {
	ScimID string       `validate:"required" json:"scim_id"`
	Data   UsersRequest `validate:"required" json:"data"`
}

// Put puts on user
func (s *UsersService) Put(ctx context.Context, req *PutUsersRequest) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "PUT", s.path, req.ScimID, req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// PostUsersRequest holds request for user
type PostUsersRequest struct {
	Data UsersRequest `validate:"required" json:"data"`
}

// Post posts user
func (s *UsersService) Post(ctx context.Context, req *PostUsersRequest) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// SearchUsersRequest type
type SearchUsersRequest struct {
	Data SearchRequest `validate:"required" json:"data"`
}

// Search searching for user
func (s *UsersService) Search(ctx context.Context, req *SearchUsersRequest) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	url := fmt.Sprintf("%s/.search", s.path)
	resp, err := s.client.call(ctx, "POST", url, "", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// StatusService holds the object for Status
type StatusService struct {
	client *Client
	path   string
}

// GetHealthy gets health status
func (s *StatusService) GetHealthy(ctx context.Context) (*HealthyReply, *http.Response, error) {
	reply := &HealthyReply{}
	resp, err := s.client.call(ctx, "GET", "status/healthy", "", nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// InvitesService holds the object for Invites
type InvitesService struct {
	client *Client
	path   string
}

// GetInvitesRequest type
type GetInvitesRequest struct {
	ScimID string `validate:"required" json:"scim_id"`
}

// Get gets one invite
func (s *InvitesService) Get(ctx context.Context, req *GetInvitesRequest) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// DeleteInvitesRequest type
type DeleteInvitesRequest struct {
	ScimID string `validate:"required" json:"scim_id"`
}

// Delete deletes one invite
func (s *InvitesService) Delete(ctx context.Context, req *DeleteInvitesRequest) (*EmptyStruct, *http.Response, error) {
	reply := &EmptyStruct{}
	resp, err := s.client.call(ctx, "DELETE", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// PostInvitesRequest type
type PostInvitesRequest struct {
	Data InvitesRequest `validate:"required" json:"data"`
}

// Post creates a new invite
func (s *InvitesService) Post(ctx context.Context, req *PostInvitesRequest) (*UsersReply, *http.Response, error) {
	reply := &UsersReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// SearchInvitesRequest type
type SearchInvitesRequest struct {
	Data SearchRequest `validate:"required" json:"data"`
}

// Search searches for an invite
func (s *InvitesService) Search(ctx context.Context, req *SearchInvitesRequest) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	url := fmt.Sprintf("%s/.search", s.path)
	resp, err := s.client.call(ctx, "POST", url, "", req, reply)
	if err != nil {
		return nil, nil, err
	}
	return reply, resp, nil
}

// GroupsService holds the service for groups
type GroupsService struct {
	client *Client
	path   string
}

// GetAll return all groups
func (s *GroupsService) GetAll(ctx context.Context) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, "", nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, err
}

// GetGroupsRequest type
type GetGroupsRequest struct {
	ScimID string `validate:"required" json:"scim_id"`
}

// Get gets one group
func (s *GroupsService) Get(ctx context.Context, req *GetGroupsRequest) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// PostGroupsRequest type
type PostGroupsRequest struct {
	Data GroupsRequest `validate:"required" json:"data"`
}

// Post posts a group
func (s *GroupsService) Post(ctx context.Context, req *PostGroupsRequest) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// PutGroupsRequest type
type PutGroupsRequest struct {
	ScimID string        `validate:"required" json:"scim_id"`
	Data   GroupsRequest `validate:"required" json:"data"`
}

// Put puts group
func (s *GroupsService) Put(ctx context.Context, req *PutGroupsRequest) (*GroupsReply, *http.Response, error) {
	reply := &GroupsReply{}
	resp, err := s.client.call(ctx, "PUT", s.path, req.ScimID, req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// DeleteGroupsRequest type
type DeleteGroupsRequest struct {
	ScimID string `validate:"required" json:"scim_id"`
}

// Delete deletes a group
func (s *GroupsService) Delete(ctx context.Context, req *DeleteGroupsRequest) (*EmptyStruct, *http.Response, error) {
	reply := &EmptyStruct{}
	resp, err := s.client.call(ctx, "DELETE", s.path, req.ScimID, nil, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// SearchGroupsRequest type
type SearchGroupsRequest struct {
	Data SearchRequest `validate:"required" json:"data"`
}

// Search group
func (s *GroupsService) Search(ctx context.Context, req *SearchGroupsRequest) (*SearchReply, *http.Response, error) {
	reply := &SearchReply{}
	url := fmt.Sprintf("%s/.search", s.path)
	resp, err := s.client.call(ctx, "POST", url, "", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// EventsService holds events object
type EventsService struct {
	client *Client
	path   string
}

// GetEventRequest type
type GetEventRequest struct {
	ScimID string        `validate:"required" json:"scim_id"`
	Data   EventsRequest `validate:"required" json:"data"`
}

// Get gets one event
func (s *EventsService) Get(ctx context.Context, req *GetEventRequest) (*EventsReply, *http.Response, error) {
	reply := &EventsReply{}
	resp, err := s.client.call(ctx, "GET", s.path, req.ScimID, req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}

// PostEventsRequest type
type PostEventsRequest struct {
	Data EventsRequest `validate:"required" json:"data"`
}

// Post create a new event
func (s *EventsService) Post(ctx context.Context, req *PostEventsRequest) (*EventsReply, *http.Response, error) {
	reply := &EventsReply{}
	resp, err := s.client.call(ctx, "POST", s.path, "", req, reply)
	if err != nil {
		return nil, resp, err
	}
	return reply, resp, nil
}
