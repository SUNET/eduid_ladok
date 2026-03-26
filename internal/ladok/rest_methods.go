package ladok

import (
	"context"
	"eduid_ladok/pkg/helpers"
	"errors"
	"fmt"
)

func (s *Service) getSchoolID(ctx context.Context) error {
	r, resp, err := s.Rest.Ladok.Kataloginformation.GetGrunddataLarosatesinformation(ctx)
	if err != nil {
		return fmt.Errorf("%w %s", err, helpers.FormatResponse(resp))
	}

	if len(r.Larosatesinformation) == 0 {
		return errors.New("Larosatesinformation is empty")
	}
	s.SchoolID = r.Larosatesinformation[0].LarosateID
	return nil
}
