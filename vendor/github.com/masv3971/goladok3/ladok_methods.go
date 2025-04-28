package goladok3

import (
	"context"

	"github.com/masv3971/goladok3/ladoktypes"
)

func (c *Client) environment(ctx context.Context) (string, error) {
	switch c.certificate.Subject.OrganizationalUnit[1] {
	case ladoktypes.EnvIntTestAPI:
		return ladoktypes.EnvIntTestAPI, nil
	case ladoktypes.EnvProdAPI:
		return ladoktypes.EnvProdAPI, nil
	case ladoktypes.EnvTestAPI:
		return ladoktypes.EnvTestAPI, nil
	default:
		return "", ladoktypes.ErrNoEnvFound
	}
}

// StudentDegree is a student degree.
type StudentDegree struct {
	Name string `json:"name"`
}

// MyStudentDegrees array of student degrees.
type MyStudentDegrees []StudentDegree

// MarshalPDF marshal MyStudentDegrees to PDF.
func (degrees *MyStudentDegrees) MarshalPDF() {}

// GetMyStudentDegrees get student data.
func (c *Client) GetMyStudentDegrees(ctx context.Context) (MyStudentDegrees, error) {
	myStudentDegrees := []StudentDegree{}

	return myStudentDegrees, nil
}

// IsStudentReq is a request to check if a user is a student.
type IsStudentReq struct {
	UID          string `validate:"required_without_all=Personnummer ExterntUID"`
	ExterntUID   string `validate:"required_without_all=Personnummer UID"`
	Personnummer string `validate:"required_without_all=UID ExterntUID"`
}

// IsStudent check if requested user is a student.
func (c *Client) IsStudent(ctx context.Context, req *IsStudentReq) (bool, error) {
	getStudentReq := &GetStudentReq{
		UID:          req.UID,
		ExterntUID:   req.ExterntUID,
		Personnummer: req.Personnummer,
	}
	_, _, err := c.Studentinformation.GetStudent(ctx, getStudentReq)
	if err != nil {
		return false, err
	}
	return false, nil
}
